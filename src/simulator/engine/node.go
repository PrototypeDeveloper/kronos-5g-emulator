package engine

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	libngap "libngap"
	"libngap/ngapMessage"

	"kronos/command"
	"kronos/models"
	"kronos/scenario"
	"kronos/transport/sctp"
)

// newNode makes the actor of a connected association.
func newNode(e *Engine, conn *sctp.Conn) *node {
	out, in := conn.Streams()
	n := &node{e: e, conn: conn, inbox: make(chan any, nodeInbox), done: make(chan struct{}), outStreams: out}
	n.status = models.NodeStatus{
		Role:      e.role.Name(),
		Connected: true,
		Local:     conn.LocalAddr().String(), Remote: conn.RemoteAddr().String(),
		OutStreams: out, InStreams: in,
	}
	return n
}

// start runs the actor and the goroutine that receives from the association.
func (n *node) start() {
	go n.loop()
	go n.read()
}

// read passes received PDUs to the inbox until the association ends.
func (n *node) read() {
	for {
		b, info, err := n.conn.Recv()
		if err != nil {
			n.inbox <- lostMsg{err}
			return
		}
		n.inbox <- rxMsg{b, info}
	}
}

// loop is the node actor: it runs the node's entry block, then the inbox until the link is lost.
func (n *node) loop() {
	defer close(n.done)
	st := n.snapshot()
	n.e.publish(models.Event{Kind: models.KindLink, Source: models.SourceNode,
		Text: fmt.Sprintf("SCTP connected: %s → %s (streams out %d / in %d)", st.Local, st.Remote, st.OutStreams, st.InStreams)})
	if b := n.e.scen.Node; b != nil && b.Entry != nil {
		n.run("node/entry", b.Entry.Body)
	}
	for m := range n.inbox {
		switch m := m.(type) {
		case rxMsg:
			n.receive(m)
		case execMsg:
			m.reply <- execResult{imsi: models.SourceNode, err: command.Exec(n, m.call)}
		case timerMsg:
			n.timerExpired(m)
		case syncMsg:
			n.syncReceived(m.m)
		case lostMsg:
			n.timers.stopAll()
			// The association ended: report it and let the engine clean up.
			n.update(func(st *models.NodeStatus) { st.Connected = false })
			text := fmt.Sprintf("SCTP disconnected: %v", m.err)
			switch {
			case errors.Is(m.err, os.ErrClosed):
				text = "SCTP disconnected"
			case errors.Is(m.err, io.EOF):
				text = "the peer disconnected SCTP"
			}
			n.e.publish(models.Event{Kind: models.KindLink, Source: models.SourceNode, Text: text})
			n.e.linkLost(n)
			return
		}
	}
}

// exec runs a command on the node and waits for it to finish.
func (n *node) exec(c command.Call) error {
	reply := make(chan execResult, 1)
	select {
	case n.inbox <- execMsg{call: c, reply: reply}:
	case <-n.done:
		return errors.New("SCTP is disconnected (connect with linkup)")
	}
	select {
	case r := <-reply:
		return r.err
	case <-n.done:
		return errors.New("SCTP is disconnected (connect with linkup)")
	}
}

// receive handles a received PDU: Paging and UE-associated PDUs go to subscribers, the rest to the node block.
func (n *node) receive(m rxMsg) {
	ev := models.Event{Kind: models.KindMessage, Source: models.SourceNode, Dir: "rx", Stream: m.info.Stream, Len: len(m.b), Hex: hexOf(m.b)}
	msg, err := libngap.Decode(m.b)
	// An undecodable PDU is only reported; a partial decode goes on with a notice.
	if msg == nil {
		ev.Kind, ev.Text = models.KindError, fmt.Sprintf("cannot decode NGAP: %v", err)
		n.e.publish(ev)
		return
	}
	if err != nil {
		ev.Text = fmt.Sprintf("(decoding notice: %v)", err)
	}
	info, _ := scenario.NGAPOf(msg)
	ev.NGAP = info.Name
	if info.Name == models.PagingName {
		// Paging names the UE by its 5G-S-TMSI: the subscriber handles it.
		key, ok := scenario.PagingKeyOf(msg)
		if s := n.e.routes.byPaging(key); ok && s != nil {
			s.post(ueRx{msg: msg, raw: m.b, info: m.info})
			return
		}
		if n.e.nodeHandlers[info.Name] == nil {
			ev.Kind, ev.Handler = models.KindWarn, "unhandled"
			if ok {
				ev.Text = "no subscriber for Paging (5G-S-TMSI " + key + ")"
			} else {
				ev.Text = "the UE identity of Paging is not a 5G-S-TMSI (I-RNTI is not supported)"
			}
			n.e.publish(ev)
			return
		}
	}
	// UE-associated PDUs go to the subscriber actors.
	if scenario.ScopeOf(msg) == command.ScopeUE {
		n.e.deliverUE(msg, m, info.Name)
		return
	}
	// A non-UE PDU: publish it, update the node status through the role and run its handler.
	n.lastMsg = info.Name
	h := n.e.nodeHandlers[info.Name]
	ev.Handler = "unhandled"
	if h != nil {
		ev.Handler = "node"
	}
	n.e.publish(ev)
	var note string
	var setupOK bool
	n.update(func(st *models.NodeStatus) {
		note = n.e.role.NodeReceived(msg, st)
		setupOK = st.Setup == models.SetupOK
	})
	if note != "" {
		n.e.publish(models.Event{Kind: models.KindLink, Source: models.SourceNode, NGAP: info.Name, Text: note})
	}
	if r, ok := msg.(*ngapMessage.NGReset); ok {
		n.e.resetAssociations(r.ResetType, "NG Reset from the AMF")
	}
	if h != nil {
		n.run("node/receive("+h.Trigger.String()+")", h.Body)
	}
	// Start the subscribers once NG Setup has succeeded.
	if setupOK && !n.subsStarted {
		n.subsStarted = true
		n.e.startSubscribers()
	}
}

// run executes scenario calls in order. A failing call is reported and the
// rest still run.
func (n *node) run(where string, body []scenario.Stmt) {
	scenario.Run(n, body, func(what string, err error) {
		n.e.publish(models.Event{Kind: models.KindError, Source: models.SourceNode, Handler: where, Text: fmt.Sprintf("%s: %v", what, err)})
	})
}

// update changes the status under its lock.
func (n *node) update(f func(st *models.NodeStatus)) {
	n.mu.Lock()
	defer n.mu.Unlock()
	f(&n.status)
}

// snapshot returns a copy of the status (any goroutine).
func (n *node) snapshot() models.NodeStatus {
	n.mu.Lock()
	defer n.mu.Unlock()
	st := n.status
	st.PeerInfo = append([]string(nil), st.PeerInfo...)
	return st
}

// Scope implements command.Env.
func (n *node) Scope() command.Scope { return command.ScopeNode }

// Send implements command.Env: build a node signal and send it on stream 0.
func (n *node) Send(name string) error {
	sig, ok := n.e.signals().Lookup(name)
	switch {
	case !ok:
		return fmt.Errorf("signal %s is not in signal.txt", name)
	case sig.Scope&command.ScopeNode == 0:
		return fmt.Errorf("%s is a UE signal, so the node cannot send it", name)
	}
	b, err := n.e.buildNode(sig)
	if err != nil {
		return fmt.Errorf("cannot build %s: %v", name, err)
	}
	if err := n.conn.Send(b, 0, models.PPIDNGAP); err != nil {
		return fmt.Errorf("cannot send %s: %v", name, err)
	}
	// An NG Reset this gNB sends releases the UE associations it names.
	if sig.NGAP != nil && sig.NGAP.Name == "NGReset" {
		if msg, _ := libngap.Decode(b); msg != nil {
			n.e.resetAssociations(msg.(*ngapMessage.NGReset).ResetType, "NG Reset sent")
		}
	}
	ev := models.Event{Kind: models.KindMessage, Source: models.SourceNode, Dir: "tx", Signal: name, Len: len(b), Hex: hexOf(b)}
	if sig.NGAP != nil {
		ev.NGAP = sig.NGAP.Name
	}
	n.e.publish(ev)
	return nil
}

// Move implements command.Env; the node has no states.
func (n *node) Move(string) error { return errors.New("the node has no state") }

// Print implements command.Env (snap output).
func (n *node) Print(text string) {
	n.e.publish(models.Event{Kind: models.KindSnap, Source: models.SourceNode, Text: text})
}

// Var implements command.Env; the node sees public variables only.
func (n *node) Var(name string) (command.Value, error) {
	if v, ok := n.e.vars.get(name); ok {
		return v, nil
	}
	return command.Value{}, fmt.Errorf("there is no public variable %s", name)
}

// Assign implements command.Env; the node has public variables only.
func (n *node) Assign(name string, eval func() (command.Value, error)) error {
	return n.e.vars.assign(name, eval)
}

// Builtin implements command.Env.
func (n *node) Builtin(name string) (command.Value, error) {
	switch name {
	case "msg":
		return command.StrVal(n.lastMsg), nil
	case "time":
		return command.StrVal(time.Now().Format(timeLayout)), nil
	}
	return command.Value{}, fmt.Errorf("$%s cannot be used on the node", name)
}

// LinkUp, LinkDown and Reload implement command.Env; they are console commands.
func (n *node) LinkUp() error   { return errors.New("linkup is a console command") }
func (n *node) LinkDown() error { return errors.New("linkdown is a console command") }
func (n *node) Reload() error   { return errors.New("reload is a console command") }

// Ping, Iperf and Stop implement command.Env; the traffic runs on a subscriber's PDU session.
func (n *node) Ping(models.PingArgs) error   { return errors.New("ping is a subscriber command") }
func (n *node) Iperf(models.IperfArgs) error { return errors.New("iperf is a subscriber command") }
func (n *node) Stop(string) error            { return errors.New("stop is a subscriber command") }

// Sleep implements command.Env; wait belongs to the command line and batches.
func (n *node) Sleep(time.Duration) error {
	return errors.New("wait is a command of the CUI and batches")
}
