package engine

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"kronos/command"
	"kronos/models"
	"kronos/scenario"
	"kronos/traffic"
)

// newSubscriber makes the actor of subscriber index, in its initial state.
func newSubscriber(e *Engine, index int, imsi string) *subscriber {
	s := &subscriber{e: e, index: index, imsi: imsi, inbox: make(chan any, subscriberInbox), ue: e.role.NewUE(index, imsi),
		stack: traffic.NewStack(), runs: trafficRuns{runs: map[int]trafficRun{}}}
	s.reset()
	s.refresh()
	return s
}

// refresh takes a new snapshot of the subscriber (on its actor) and returns it.
func (s *subscriber) refresh() models.UEStatus {
	st := models.UEStatus{IMSI: s.imsi, State: s.state, LastMsg: s.lastMsg}
	s.ue.Status(&st)
	for i := range st.Sessions {
		_, st.Sessions[i].Active = s.tunnels[st.Sessions[i].ID]
	}
	if len(s.vars) > 0 {
		st.Vars = make(map[string]string, len(s.vars))
		for k, v := range s.vars {
			st.Vars[k] = v.Literal()
		}
	}
	s.statMu.Lock()
	s.stat = st
	s.statMu.Unlock()
	return st
}

// status returns the last snapshot (any goroutine).
func (s *subscriber) status() models.UEStatus {
	s.statMu.Lock()
	defer s.statMu.Unlock()
	return s.stat
}

// plainHex renders the plaintext NAS of an event ("" where in clear).
func plainHex(plain [][]byte) []string {
	if len(plain) == 0 {
		return nil
	}
	out := make([]string, len(plain))
	found := false
	for i, p := range plain {
		if p != nil {
			out[i], found = hex.EncodeToString(p), true
		}
	}
	if !found {
		return nil
	}
	return out
}

// post queues a message for the actor.
func (s *subscriber) post(m any) { s.inbox <- m }

// loop is the subscriber actor: it handles the inbox until ctx is cancelled.
func (s *subscriber) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-s.inbox:
			s.handle(m)
		}
	}
}

// handle processes one inbox message on the actor.
func (s *subscriber) handle(m any) {
	// A session goes on a receive or on a send (a release, a deregistration).
	defer func() { s.dropGone(s.refresh()) }()
	switch m := m.(type) {
	case startMsg:
		// Entering the initial state runs its entry.
		s.moved = "start"
		s.afterHandler()
	case resetMsg:
		s.reset()
	case releaseMsg:
		// The UE association ends without a release procedure (NG Reset).
		s.e.routes.unbind(s)
		s.dropTunnels()
		s.ue.ReleaseNGAP()
		s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: "UE association released (" + m.why + "); the NAS state stays"})
	case timerMsg:
		s.timerExpired(m)
	case syncMsg:
		s.syncReceived(m.m)
	case ueRx:
		s.receive(m)
	case execMsg:
		err := command.Exec(s, m.call)
		s.afterHandler()
		s.refresh() // before the reply, so the caller sees the result
		m.reply <- execResult{imsi: s.imsi, err: err}
	}
}

// reset puts the subscriber back to its initial state (link down).
func (s *subscriber) reset() {
	s.stopRuns("")
	s.timers.stopAll()
	s.dropTunnels()
	s.tunDown()
	s.ue.Reset()
	s.state, s.lastMsg, s.moved = "start", "", ""
	s.vars = map[string]command.Value{}
	for _, v := range s.e.scen.Vars {
		if !v.Public {
			s.vars[v.Name] = v.Init
		}
	}
}

// receive handles a UE-associated PDU: the role's UE context, routes, N3 tunnels and the matching handler.
func (s *subscriber) receive(m ueRx) {
	in := s.ue.Receive(m.msg)
	s.rx = newReceived(m.msg, in.Plain)
	info, _ := scenario.NGAPOf(m.msg)
	// Remember the IDs that route later PDUs (AMF-UE-NGAP-ID, 5G-S-TMSI for Paging).
	if amf, ok := s.ue.AMFID(); ok {
		s.e.routes.bindAMF(amf, s)
	}
	if in.PagingKey != "" {
		s.e.routes.bindPaging(in.PagingKey, s)
	}
	// Publish the PDU with the handler that takes it, then the role's notes and errors.
	ev := models.Event{Kind: models.KindMessage, Source: s.imsi, Dir: "rx", NGAP: info.Name, State: s.state,
		Stream: m.info.Stream, Len: len(m.raw), Hex: hexOf(m.raw), Plain: plainHex(in.Plain)}
	s.lastMsg = info.Name
	if len(in.NAS) > 0 {
		ev.NAS = in.NAS[0]
		s.lastMsg += ":" + in.NAS[0]
	}
	h, block := s.findHandler(info.Name, in.NAS)
	ev.Handler = "unhandled"
	if h != nil {
		ev.Handler = block
	}
	s.e.publish(ev)
	for _, t := range in.Notes {
		s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: t})
	}
	for _, t := range in.Errors {
		s.e.publish(models.Event{Kind: models.KindError, Source: s.imsi, Text: t})
	}
	// Route the N3 tunnels that went down, then those that came up (a
	// modified session goes down and up again).
	for _, dl := range in.TunnelsDown {
		s.e.uplane.unbind(dl)
		for id, t := range s.tunnels {
			if t.DLTEID == dl {
				delete(s.tunnels, id)
			}
		}
	}
	for _, t := range in.TunnelsUp {
		if err := s.e.uplane.bind(t, s); err != nil {
			s.e.publish(models.Event{Kind: models.KindError, Source: s.imsi, Text: err.Error()})
			continue
		}
		s.tunnels[t.SessionID] = t
		s.tunUp(t)
	}
	// Run the handler, then the entry of the state it moved to.
	if h != nil {
		s.run(block+"/receive("+h.Trigger.String()+")", h.Body)
		s.afterHandler()
	}
	// UE context release: drop the NGAP routes and the tunnels.
	if in.Release {
		s.e.routes.unbind(s)
		s.dropTunnels()
		s.ue.ReleaseNGAP()
		s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: "UE context released (the NGAP IDs are dropped, the NAS state stays)"})
	}
}

// findHandler looks in the current state, then in any. Within
// a block, NGAP:NAS matches first (in the order of the NAS names), then NGAP alone.
func (s *subscriber) findHandler(ngap string, nas []string) (*scenario.Handler, string) {
	var blocks []*scenario.Block
	if b, ok := s.e.scen.State(s.state); ok {
		blocks = append(blocks, b)
	}
	if s.e.scen.Any != nil {
		blocks = append(blocks, s.e.scen.Any)
	}
	for _, b := range blocks {
		for _, name := range nas {
			if h := matchHandler(b, ngap, name); h != nil {
				return h, b.Name
			}
		}
		if h := matchHandler(b, ngap, ""); h != nil {
			return h, b.Name
		}
	}
	return nil, ""
}

// matchHandler returns the handler of b for an NGAP message and a NAS message ("" for NGAP alone).
func matchHandler(b *scenario.Block, ngap, nas string) *scenario.Handler {
	for _, h := range b.Handlers {
		t := h.Trigger
		if t.NGAPInfo == nil || t.NGAPInfo.Name != ngap {
			continue
		}
		if nas == "" && t.NAS == "" || nas != "" && t.NASInfo != nil && t.NASInfo.Name == nas {
			return h
		}
	}
	return nil
}

// run executes scenario calls in order; a failing call is reported and the
// rest still run.
func (s *subscriber) run(where string, body []scenario.Stmt) {
	scenario.Run(s, body, func(what string, err error) {
		s.e.publish(models.Event{Kind: models.KindError, Source: s.imsi, Handler: where, Text: fmt.Sprintf("%s: %v", what, err)})
	})
}

// afterHandler runs the entry of the state that move() went to last, and
// follows moves made by entries, at most maxEntryChain times.
func (s *subscriber) afterHandler() {
	for i := 0; s.moved != ""; i++ {
		if i == maxEntryChain {
			s.e.publish(models.Event{Kind: models.KindError, Source: s.imsi,
				Text: fmt.Sprintf("stopped after %d moves in a row from entry blocks (state %s)", maxEntryChain, s.state)})
			s.moved = ""
			return
		}
		st := s.moved
		s.moved = ""
		if b, ok := s.e.scen.State(st); ok && b.Entry != nil {
			s.run(st+"/entry", b.Entry.Body)
		}
	}
}

// Scope implements command.Env.
func (s *subscriber) Scope() command.Scope { return command.ScopeUE }

// Send implements command.Env: build the UE signal for this subscriber and
// send it on the subscriber's stream.
func (s *subscriber) Send(name string) error {
	sig, ok := s.e.signals().Lookup(name)
	switch {
	case !ok:
		return fmt.Errorf("signal %s is not in signal.txt", name)
	case sig.Scope&command.ScopeUE == 0:
		return fmt.Errorf("%s is a node signal, so a subscriber cannot send it", name)
	}
	n := s.e.currentNode()
	if n == nil {
		return errors.New("SCTP is not connected (connect with linkup)")
	}
	built, err := s.ue.Build(sig, models.MapOff(s.mapOff.Load()))
	if err != nil {
		return fmt.Errorf("cannot build %s: %v", name, err)
	}
	b := built.PDU
	ran, hasRAN := s.ue.RANID()
	if hasRAN {
		s.e.routes.bindRAN(ran, s) // before sending, so the answer cannot be missed
	}
	stream := streamFor(ran, n.outStreams)
	if err := n.conn.Send(b, stream, models.PPIDNGAP); err != nil {
		return fmt.Errorf("cannot send %s: %v", name, err)
	}
	ev := models.Event{Kind: models.KindMessage, Source: s.imsi, Dir: "tx", Signal: name, State: s.state, Stream: stream, Len: len(b), Hex: hexOf(b),
		Plain: plainHex(built.Plain)}
	if sig.NGAP != nil {
		ev.NGAP = sig.NGAP.Name
	}
	if len(sig.NAS) > 0 {
		ev.NAS = sig.NAS[0].Name
		if sig.NAS[0].Inner != "" {
			ev.NAS = sig.NAS[0].Inner
		}
	}
	s.e.publish(ev)
	return nil
}

// streamFor spreads UE-associated messages over the streams other than 0.
func streamFor(ranID int64, out uint16) uint16 {
	if out <= 1 {
		return 0
	}
	return uint16(1 + ranID%int64(out-1))
}

// Move implements command.Env.
func (s *subscriber) Move(state string) error {
	if _, ok := s.e.scen.State(state); !ok {
		return fmt.Errorf("there is no state %s", state)
	}
	s.e.publish(models.Event{Kind: models.KindState, Source: s.imsi, State: state, Text: fmt.Sprintf("state %s → %s", s.state, state)})
	s.state, s.moved = state, state
	return nil
}

// Print implements command.Env (snap output).
func (s *subscriber) Print(text string) {
	s.e.publish(models.Event{Kind: models.KindSnap, Source: s.imsi, Text: text})
}

// Var implements command.Env: private variables first, then public ones.
func (s *subscriber) Var(name string) (command.Value, error) {
	if v, ok := s.vars[name]; ok {
		return v, nil
	}
	if v, ok := s.e.vars.get(name); ok {
		return v, nil
	}
	return command.Value{}, fmt.Errorf("there is no variable %s", name)
}

// Assign implements command.Env: a private variable of the subscriber, or a public one.
func (s *subscriber) Assign(name string, eval func() (command.Value, error)) error {
	old, ok := s.vars[name]
	if !ok {
		return s.e.vars.assign(name, eval)
	}
	v, err := eval()
	switch {
	case err != nil:
		return err
	case old.Type != v.Type:
		return fmt.Errorf("variable %s is of type %s (a %s cannot be assigned)", name, old.Type, v.Type)
	}
	s.vars[name] = v
	return nil
}

// Builtin implements command.Env.
func (s *subscriber) Builtin(name string) (command.Value, error) {
	switch name {
	case "ue":
		return command.IntVal(int64(s.index + 1)), nil
	case "imsi":
		return command.StrVal(s.imsi), nil
	case "state":
		return command.StrVal(s.state), nil
	case "msg":
		return command.StrVal(s.lastMsg), nil
	case "time":
		return command.StrVal(time.Now().Format(timeLayout)), nil
	}
	return command.Value{}, fmt.Errorf("unknown builtin $%s", name)
}

// LinkUp, LinkDown and Reload implement command.Env; they are console commands.
func (s *subscriber) LinkUp() error { return errors.New("linkup is a console command") }
func (s *subscriber) LinkDown() error {
	return errors.New("linkdown is a console command")
}
func (s *subscriber) Reload() error { return errors.New("reload is a console command") }

// dropTunnels stops routing downlink packets to this subscriber.
func (s *subscriber) dropTunnels() {
	for _, t := range s.tunnels {
		s.e.uplane.unbind(t.DLTEID)
	}
	s.tunnels = map[uint8]models.Tunnel{}
}

// Sleep implements command.Env; wait belongs to the command line and batches.
func (s *subscriber) Sleep(time.Duration) error {
	return errors.New("wait is a command of the CUI and batches")
}
