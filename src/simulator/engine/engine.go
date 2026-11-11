// Package engine is the role-independent core of the pseudo-call: it owns
// the SCTP association (linkup / linkdown), runs the node actor
// and one actor per subscriber, executes scenario handlers and CUI/GUI
// commands, and publishes every step on the event bus.
package engine

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"net/netip"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	libngap "libngap"
	"libngap/ngapMessage"

	"kronos/command"
	"kronos/common/diag"
	"kronos/models"
	"kronos/scenario"
	"kronos/simulator/role"
	"kronos/state"
	"kronos/transport/peer"
	"kronos/transport/sctp"
)

// New prepares an engine. scen must have passed scenario.Validate. When
// base.snap is false, snap calls are removed here so they cost nothing at
// run time.
func New(cfg *state.Config, sigs *scenario.SignalSet, scen *scenario.File, r role.Role) *Engine {
	e := &Engine{
		cfg:          cfg,
		scen:         scen,
		role:         r,
		bus:          NewBus(eventHistory),
		vars:         newVarStore(scen),
		controls:     newControls(),
		snap:         cfg.Base.Snap,
		nodeHandlers: map[string]*scenario.Handler{},
		routes:       newRoutes(),
		stopped:      make(chan struct{}),
		started:      time.Now(),
	}
	if src := cfg.SyncListen(); src != nil {
		if a, err := netip.ParseAddr(src.Address); err == nil {
			e.peer = peer.NewClient(netip.AddrPortFrom(a, uint16(src.Port)))
		}
	}
	e.sigs.Store(sigs)
	bg := context.Background()
	e.ctx.Store(&bg)
	e.uplane = newUplane(e)
	if hasNetAdmin() { // the UEs get TUN devices whenever gnodeb may make them
		e.openTUN = openHostTUN
	}
	if !e.snap {
		scen.StripCalls("snap")
	}
	if scen.Node != nil {
		for _, h := range scen.Node.Handlers {
			if info, ok := scenario.LookupNGAP(h.Trigger.NGAP); ok {
				e.nodeHandlers[info.Name] = h
			}
		}
	}
	for i := 0; i < cfg.UE.Subscription; i++ {
		imsi, err := cfg.SubscriberIMSI(i)
		if err != nil {
			break // config validation reports it
		}
		e.subs = append(e.subs, newSubscriber(e, i, imsi))
	}
	return e
}

// Events returns the event bus.
func (e *Engine) Events() *Bus { return e.bus }

// publish puts an event on the bus.
func (e *Engine) publish(ev models.Event) { e.bus.Publish(ev) }

// runCtx returns the context of the running engine (Background before Start).
func (e *Engine) runCtx() context.Context { return *e.ctx.Load() }

// Start runs the subscriber actors until ctx is cancelled; the association,
// if any, is then shut down. Start does not connect: that is LinkUp.
func (e *Engine) Start(ctx context.Context) {
	e.mu.Lock()
	if e.running {
		e.mu.Unlock()
		return
	}
	e.running = true
	e.ctx.Store(&ctx)
	e.mu.Unlock()
	var wg sync.WaitGroup
	for _, s := range e.subs {
		wg.Add(1)
		go func(s *subscriber) {
			defer wg.Done()
			s.loop(ctx)
		}(s)
	}
	go func() {
		<-ctx.Done()
		e.LinkDown()
		wg.Wait()
		for _, s := range e.subs {
			s.tunDown()
			s.timers.stopAll() // the actors have stopped
		}
		close(e.stopped)
	}()
}

// Wait blocks until the engine has stopped after its context was cancelled.
func (e *Engine) Wait() { <-e.stopped }

// Status returns a snapshot of the node.
func (e *Engine) Status() models.NodeStatus {
	if n := e.currentNode(); n != nil {
		return n.snapshot()
	}
	return models.NodeStatus{Role: e.role.Name()}
}

// currentNode returns the node of the current association, or nil while the link is down.
func (e *Engine) currentNode() *node {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.node
}

// LinkUp sets up the N2 association (command linkup). The node's entry block
// runs once it is up.
func (e *Engine) LinkUp() error {
	e.mu.Lock()
	switch {
	case !e.running:
		e.mu.Unlock()
		return errors.New("the engine is not running")
	case e.linkBusy || e.node != nil:
		e.mu.Unlock()
		return errors.New("SCTP is already connected (disconnect with linkdown)")
	}
	e.linkBusy = true
	e.mu.Unlock()
	defer func() {
		e.mu.Lock()
		e.linkBusy = false
		e.mu.Unlock()
	}()
	if e.role.Listen() {
		return fmt.Errorf("the %s role (which accepts connections) is not implemented", e.role.Name())
	}
	local, remote, err := e.endpoints()
	if err != nil {
		return err
	}
	e.publish(models.Event{Kind: models.KindLink, Source: models.SourceNode, Text: fmt.Sprintf("SCTP connecting: %s → %s", fmtAddr(local), remote)})
	conn, err := sctp.Dial(local, remote, sctp.Options{Timeout: sctpDialTimeout})
	if err != nil {
		e.publish(models.Event{Kind: models.KindError, Source: models.SourceNode, Text: fmt.Sprintf("SCTP connection failed: %v", err)})
		return fmt.Errorf("SCTP connection failed: %w", err)
	}
	n := newNode(e, conn)
	e.mu.Lock()
	e.node = n
	e.mu.Unlock()
	n.start()
	return nil
}

// LinkDown shuts the association down (command linkdown) and waits until the
// node actor has finished; the subscribers then return to their initial state.
func (e *Engine) LinkDown() error {
	n := e.currentNode()
	if n == nil {
		return errors.New("SCTP is not connected (connect with linkup)")
	}
	n.conn.Close()
	<-n.done
	return nil
}

// linkLost runs on the node actor when its association has ended.
func (e *Engine) linkLost(n *node) {
	e.mu.Lock()
	if e.node == n {
		e.node = nil
	}
	e.mu.Unlock()
	e.routes.clear()
	e.uplane.close()
	e.controls.reset() // first, so that no subscriber still waits for its turn
	for _, s := range e.subs {
		s.post(resetMsg{})
	}
	if len(e.subs) > 0 {
		e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: fmt.Sprintf("%d subscriber(s) back to the initial state", len(e.subs))})
	}
}

// startSubscribers runs the start state's entry of every subscriber, in IMSI
// order, once NG Setup has succeeded.
func (e *Engine) startSubscribers() {
	for _, s := range e.subs {
		s.post(startMsg{})
	}
}

// endpoints returns the local (optional) and remote N2 SCTP addresses from the config.
func (e *Engine) endpoints() (local, remote netip.AddrPort, err error) {
	g := e.cfg.GNodeB
	if s := g.NGAP.Src; s != nil {
		a, err := netip.ParseAddr(s.Address)
		if err != nil {
			return local, remote, fmt.Errorf("gnodeb.ngap.src: %v", err)
		}
		local = netip.AddrPortFrom(a, uint16(s.Port))
	}
	if g.NGAP.Dst == nil {
		return local, remote, errors.New("gnodeb.ngap.dst is missing")
	}
	a, err := netip.ParseAddr(g.NGAP.Dst.Address)
	if err != nil {
		return local, remote, fmt.Errorf("gnodeb.ngap.dst: %v", err)
	}
	return local, netip.AddrPortFrom(a, uint16(g.NGAP.Dst.Port)), nil
}

// fmtAddr formats a local address; "(auto)" when the kernel chooses it.
func fmtAddr(a netip.AddrPort) string {
	if !a.IsValid() {
		return "(auto)"
	}
	return a.String()
}

// Exec runs one CUI/GUI command line. Text output (help) is
// returned; everything else is reported on the event bus. "@file" runs the
// commands of file as a batch.
func (e *Engine) Exec(line string) (string, error) { return e.exec(line, "", 0, false) }

// exec runs a line; dir resolves relative batch files and batch makes
// commands for every subscriber wait until all of them are done.
func (e *Engine) exec(line, dir string, depth int, batch bool) (string, error) {
	if t := strings.TrimSpace(line); strings.HasPrefix(t, "@") {
		// "@file", optionally followed by a comment; the name has no spaces.
		name, tail, _ := strings.Cut(strings.TrimSpace(t[1:]), " ")
		if tail = strings.TrimSpace(tail); tail != "" && !strings.HasPrefix(tail, "#") && !strings.HasPrefix(tail, "//") {
			return "", fmt.Errorf("@ takes one file name only (found %q)", tail)
		}
		return e.runBatch(strings.TrimSpace(strings.Split(name, "\t")[0]), dir, depth)
	}
	ln, err := command.ParseLine(line)
	if errors.Is(err, command.ErrEmptyLine) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	_, subject, ds := command.CheckLine(ln, e)
	for _, d := range ds {
		if d.Sev == diag.Error {
			return "", errors.New(d.Msg)
		}
	}
	if ln.Call.Name == "snap" && !e.snap {
		return "", errors.New("snap is disabled (base.snap of config is not true)")
	}
	switch subject {
	case command.ScopeConsole:
		var out strings.Builder
		err := command.Exec(consoleEnv{e, &out}, ln.Call)
		return strings.TrimSuffix(out.String(), "\n"), err
	case command.ScopeNode:
		n := e.currentNode()
		if n == nil {
			return "", errors.New("SCTP is not connected (connect with linkup)")
		}
		return "", n.exec(ln.Call)
	default:
		return "", e.execSubscribers(ln, batch)
	}
}

// runBatch executes the lines of a file in order, each after the previous one
// has finished, and stops at the first error. Blank lines and comments are
// skipped; "@file" lines nest, relative to the including file.
func (e *Engine) runBatch(name, dir string, depth int) (string, error) {
	if name == "" {
		return "", errors.New("write a file name after @ (e.g. @cmds.txt)")
	}
	if depth >= maxBatchDepth {
		return "", fmt.Errorf("@ nested too deep (%d levels at most)", maxBatchDepth)
	}
	path := name
	if !filepath.IsAbs(path) && dir != "" {
		path = filepath.Join(dir, path)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read the batch file: %v", err)
	}
	e.publish(models.Event{Kind: models.KindCommand, Source: models.SourceNode, Text: "batch started: " + path})
	n := 0
	for i, raw := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(raw)
		if _, err := command.ParseLine(line); errors.Is(err, command.ErrEmptyLine) {
			continue // blank or comment
		}
		where := fmt.Sprintf("%s:%d", path, i+1)
		e.publish(models.Event{Kind: models.KindCommand, Source: models.SourceNode, Text: where + ": " + line})
		out, err := e.exec(line, filepath.Dir(path), depth+1, true)
		for _, o := range strings.Split(out, "\n") {
			if o != "" {
				e.publish(models.Event{Kind: models.KindCommand, Source: models.SourceNode, Text: o})
			}
		}
		if err != nil {
			e.publish(models.Event{Kind: models.KindError, Source: models.SourceNode, Text: fmt.Sprintf("%s: %v", where, err)})
			return "", fmt.Errorf("batch stopped at %s: %v", where, err)
		}
		n++
	}
	return fmt.Sprintf("batch %s done (%d lines)", path, n), nil
}

// execSubscribers runs a subscriber command on the target, or on every
// subscriber when no target is given. A single target is
// waited for; for all subscribers a summary is published when they are done,
// and in a batch (wait) the call returns only then, failing if anyone failed.
func (e *Engine) execSubscribers(ln command.Line, wait bool) error {
	e.mu.Lock()
	running := e.running
	e.mu.Unlock()
	if !running {
		return errors.New("the engine is not running")
	}
	// Pick the target subscribers.
	targets := e.subs
	if ln.Target != "" {
		i, ok := e.cfg.SubscriberIndex(ln.Target)
		if !ok {
			return fmt.Errorf("there is no subscriber with IMSI %s", ln.Target)
		}
		targets = e.subs[i : i+1]
	}
	if len(targets) == 0 {
		return errors.New("no subscribers (ue.subscription)")
	}
	// Post the command to every target actor.
	results := make(chan execResult, len(targets))
	for _, s := range targets {
		s.post(execMsg{call: ln.Call, reply: results})
	}
	// A single target: wait for its answer.
	if ln.Target != "" {
		t := time.NewTimer(subscriberReplyTimeout) // stopped at once, not left until it fires
		defer t.Stop()
		select {
		case r := <-results:
			return r.err
		case <-t.C:
			return fmt.Errorf("the subscriber does not answer (%d s)", subscriberReplyTimeout/time.Second)
		}
	}
	// collect waits for all the results, reports the failures and publishes a summary.
	collect := func() int {
		ok := 0
		for range targets {
			r := <-results
			if r.err == nil {
				ok++
				continue
			}
			e.publish(models.Event{Kind: models.KindError, Source: r.imsi, Text: fmt.Sprintf("%s: %v", ln.Call, r.err)})
		}
		e.publish(models.Event{Kind: models.KindCommand, Source: models.SourceNode,
			Text: fmt.Sprintf("%s: %d succeeded, %d failed", ln.Call, ok, len(targets)-ok)})
		return ok
	}
	// Outside a batch, collect in the background and return at once.
	if !wait {
		go collect()
		return nil
	}
	if ok := collect(); ok < len(targets) {
		return fmt.Errorf("%d subscriber(s) failed", len(targets)-ok)
	}
	return nil
}

// buildNode turns a non-UE template into the bytes to send: decode a fresh
// copy, apply the role's auto-fix, encode.
func (e *Engine) buildNode(sig *scenario.Signal) ([]byte, error) {
	if sig.Mode == scenario.ModeRaw {
		return sig.Bytes, nil
	}
	msg, err := libngap.Decode(sig.Bytes)
	if msg == nil {
		return nil, err
	}
	if err := e.role.PatchNode(msg, models.NodeBuild{Off: models.MapOff(e.nodeMapOff.Load()), Associations: e.routes.associations()}); err != nil {
		return nil, err
	}
	return libngap.Encode(msg)
}

// deliverUE routes a received UE-associated PDU to its subscriber by
// RAN-UE-NGAP-ID, falling back to AMF-UE-NGAP-ID.
func (e *Engine) deliverUE(msg any, m rxMsg, name string) {
	ran, hasRAN, amf, hasAMF := scenario.UEIDs(msg)
	var s *subscriber
	if hasRAN {
		s = e.routes.byRAN(ran)
	}
	if s == nil && hasAMF {
		s = e.routes.byAMF(amf)
	}
	if _, ho := msg.(*ngapMessage.HandoverRequest); s == nil && ho {
		// A new association at this target: the subscriber that took the
		// handover context over (sync_receive(HO)).
		s = e.routes.takeHandover(scenario.HandoverULTEIDs(msg))
	}
	if s == nil {
		e.publish(models.Event{Kind: models.KindWarn, Source: models.SourceNode, NGAP: name, Dir: "rx", Len: len(m.b), Hex: hexOf(m.b),
			Text: fmt.Sprintf("no subscriber for %s (RAN-UE-NGAP-ID %s, AMF-UE-NGAP-ID %s)", name, idText(ran, hasRAN), idText(amf, hasAMF))})
		return
	}
	s.post(ueRx{msg: msg, raw: m.b, info: m.info})
}

// idText formats an optional UE NGAP ID for messages.
func idText(id int64, ok bool) string {
	if !ok {
		return "none"
	}
	return fmt.Sprint(id)
}

// Signal implements command.Resolver for the command line.
func (e *Engine) Signal(name string) (command.Scope, bool) {
	sig, ok := e.signals().Lookup(name)
	if !ok {
		return 0, false
	}
	return sig.Scope, true
}

// State implements command.Resolver.
func (e *Engine) State(name string) bool {
	_, ok := e.scen.State(name)
	return ok
}

// Var implements command.Resolver.
func (e *Engine) Var(name string) (command.VarInfo, bool) {
	for _, v := range e.scen.Vars {
		if v.Name == name {
			return command.VarInfo{Type: v.Init.Type, Public: v.Public}, true
		}
	}
	return command.VarInfo{}, false
}

// Subscriber implements command.Resolver.
func (e *Engine) Subscriber(imsi string) bool {
	_, ok := e.cfg.SubscriberIndex(imsi)
	return ok
}

// Names implements command.Resolver.
func (e *Engine) Names(k command.Kind) []string {
	var out []string
	switch k {
	case command.KindSignal:
		return e.signals().Names()
	case command.KindState:
		for _, b := range e.scen.States {
			out = append(out, b.Name)
		}
	case command.KindValue:
		for _, v := range e.scen.Vars {
			out = append(out, v.Name)
		}
		sort.Strings(out)
	}
	return out
}

// Scope, Print, LinkUp, LinkDown and Reload implement command.Env for the console.
func (c consoleEnv) Scope() command.Scope { return command.ScopeConsole }
func (c consoleEnv) Print(text string)    { c.out.WriteString(text + "\n") }
func (c consoleEnv) LinkUp() error        { return c.e.LinkUp() }
func (c consoleEnv) LinkDown() error      { return c.e.LinkDown() }
func (c consoleEnv) Reload() error        { return c.e.Reload() }

// Sleep implements command.Env (wait); cancelling the engine interrupts it.
func (c consoleEnv) Sleep(d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-c.e.runCtx().Done():
		return errors.New("interrupted")
	}
}

// The other command.Env methods belong to subscribers or the node and fail on the console.
func (consoleEnv) Ping(models.PingArgs) error { return errors.New("ping is a subscriber command") }
func (consoleEnv) Iperf(models.IperfArgs) error {
	return errors.New("iperf is a subscriber command")
}
func (consoleEnv) Stop(string) error { return errors.New("stop is a subscriber command") }
func (consoleEnv) Send(string) error { return errors.New("the console cannot send") }
func (consoleEnv) Move(string) error { return errors.New("the console has no state") }

// Var implements command.Env; the console sees the public variables.
func (c consoleEnv) Var(name string) (command.Value, error) {
	if v, ok := c.e.vars.get(name); ok {
		return v, nil
	}
	return command.Value{}, fmt.Errorf("there is no public variable %s", name)
}

// Assign implements command.Env: set of a public variable on the CUI/GUI, which prints the new value.
func (c consoleEnv) Assign(name string, eval func() (command.Value, error)) error {
	if err := c.e.vars.assign(name, eval); err != nil {
		return err
	}
	v, _ := c.e.vars.get(name)
	c.Print(name + " = " + v.Literal())
	return nil
}

// Builtin implements command.Env; the console has $time only.
func (consoleEnv) Builtin(name string) (command.Value, error) {
	if name == "time" {
		return command.StrVal(time.Now().Format(timeLayout)), nil
	}
	return command.Value{}, fmt.Errorf("$%s cannot be used on the console", name)
}

// newVarStore holds the initial values of the public variables of f.
func newVarStore(f *scenario.File) *varStore {
	s := &varStore{m: map[string]command.Value{}}
	for _, v := range f.Vars {
		if v.Public {
			s.m[v.Name] = v.Init
		}
	}
	return s
}

// get returns the current value of a public variable.
func (s *varStore) get(name string) (command.Value, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.m[name]
	return v, ok
}

// assign stores the value eval returns in a public variable. eval runs
// under assignMu, so the sets of all subscribers and the node happen one
// at a time and set(a = a + 1) never loses an addition.
func (s *varStore) assign(name string, eval func() (command.Value, error)) error {
	s.assignMu.Lock()
	defer s.assignMu.Unlock()
	v, err := eval()
	if err != nil {
		return err
	}
	return s.set(name, v)
}

// set assigns a public variable; its type cannot change.
func (s *varStore) set(name string, v command.Value) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	old, ok := s.m[name]
	switch {
	case !ok:
		return fmt.Errorf("there is no public variable %s", name)
	case old.Type != v.Type:
		return fmt.Errorf("variable %s is of type %s (a %s cannot be assigned)", name, old.Type, v.Type)
	}
	s.m[name] = v
	return nil
}

// newRoutes returns empty routes.
func newRoutes() *routes {
	return &routes{ran: map[int64]*subscriber{}, amf: map[int64]*subscriber{}, paging: map[string]*subscriber{}}
}

// bindPaging maps the 5G-S-TMSI key of s; a new 5G-GUTI replaces the old one.
func (r *routes) bindPaging(key string, s *subscriber) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for k, x := range r.paging {
		if x == s {
			delete(r.paging, k)
		}
	}
	r.paging[key] = s
}

// byPaging returns the subscriber with a 5G-S-TMSI key, or nil.
func (r *routes) byPaging(key string) *subscriber {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.paging[key]
}

// bindRAN and bindAMF route a RAN-UE-NGAP-ID or an AMF-UE-NGAP-ID to s.
func (r *routes) bindRAN(id int64, s *subscriber) { r.mu.Lock(); r.ran[id] = s; r.mu.Unlock() }

// bindAMF maps the AMF-UE-NGAP-ID of s; a new one (UE context
// modification) replaces the old.
func (r *routes) bindAMF(id int64, s *subscriber) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for k, x := range r.amf {
		if x == s {
			delete(r.amf, k)
		}
	}
	r.amf[id] = s
}

// associations lists the UE associations (RAN-UE-NGAP-ID and, when known,
// AMF-UE-NGAP-ID) by RAN-UE-NGAP-ID.
func (r *routes) associations() []models.UEAssociation {
	r.mu.Lock()
	defer r.mu.Unlock()
	amfOf := map[*subscriber]int64{}
	for id, s := range r.amf {
		amfOf[s] = id
	}
	var out []models.UEAssociation
	for id, s := range r.ran {
		a := models.UEAssociation{RANUENGAPID: id}
		a.AMFUENGAPID, a.HasAMF = amfOf[s]
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].RANUENGAPID < out[j].RANUENGAPID })
	return out
}

// awaitHandover lets s wait for a HandoverRequest, matched by ulTEIDs.
func (r *routes) awaitHandover(s *subscriber, ulTEIDs []uint32) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.dropWait(s)
	r.awaitHO = append(r.awaitHO, hoWait{s: s, ulTEIDs: ulTEIDs})
}

// takeHandover returns the waiting subscriber a HandoverRequest is for: the
// one with one of its uplink TEIDs, else the one that has waited longest.
func (r *routes) takeHandover(ulTEIDs []uint32) *subscriber {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.awaitHO) == 0 {
		return nil
	}
	pick := r.awaitHO[0].s
	for _, w := range r.awaitHO {
		if sharesTEID(w.ulTEIDs, ulTEIDs) {
			pick = w.s
			break
		}
	}
	r.dropWait(pick)
	return pick
}

// dropWait removes s from the subscribers waiting for a HandoverRequest (locked).
func (r *routes) dropWait(s *subscriber) {
	out := r.awaitHO[:0]
	for _, w := range r.awaitHO {
		if w.s != s {
			out = append(out, w)
		}
	}
	r.awaitHO = out
}

// sharesTEID reports whether a and b have a TEID in common.
func sharesTEID(a, b []uint32) bool {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return true
			}
		}
	}
	return false
}

// byRAN and byAMF return the subscriber of a UE NGAP ID, or nil.
func (r *routes) byRAN(id int64) *subscriber { r.mu.Lock(); defer r.mu.Unlock(); return r.ran[id] }
func (r *routes) byAMF(id int64) *subscriber { r.mu.Lock(); defer r.mu.Unlock(); return r.amf[id] }

// unbind removes the UE association routes of s (UE context release). The
// Paging route stays: the UE is still registered.
func (r *routes) unbind(s *subscriber) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, x := range r.ran {
		if x == s {
			delete(r.ran, id)
		}
	}
	for id, x := range r.amf {
		if x == s {
			delete(r.amf, id)
		}
	}
}

// clear drops every route (the association went down).
func (r *routes) clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ran, r.amf, r.paging = map[int64]*subscriber{}, map[int64]*subscriber{}, map[string]*subscriber{}
	r.awaitHO = nil
}

// hexOf renders octets as hex for events.
func hexOf(b []byte) string { return hex.EncodeToString(b) }
