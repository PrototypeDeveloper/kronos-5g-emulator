package engine

import (
	"errors"
	"fmt"

	"libngap/ngapType"

	"kronos/command"
	"kronos/models"
	"kronos/transport/peer"
)

// DeliverSync hands a message of another pseudo-call to the subscriber with
// its IMSI, or to the node when it has none. The handler runs later.
func (e *Engine) DeliverSync(m models.SyncMessage) error {
	if m.IMSI == "" {
		n := e.currentNode()
		if n == nil {
			return fmt.Errorf("%w: SCTP is not connected, so the node is not running", peer.ErrUnavailable)
		}
		select {
		case n.inbox <- syncMsg{m}:
			return nil
		case <-n.done:
			return fmt.Errorf("%w: the SCTP association went down", peer.ErrUnavailable)
		}
	}
	i, ok := e.cfg.SubscriberIndex(m.IMSI)
	if !ok || i >= len(e.subs) {
		return fmt.Errorf("%w: there is no subscriber with IMSI %s", peer.ErrNotFound, m.IMSI)
	}
	select {
	case e.subs[i].inbox <- syncMsg{m}:
		return nil
	case <-e.runCtx().Done():
		return fmt.Errorf("%w: shutting down", peer.ErrUnavailable)
	}
}

// send posts m from this pseudo-call to dst.
func (e *Engine) syncSend(dst string, m models.SyncMessage) error {
	if e.peer == nil {
		return errors.New("gnodeb.http.src is not set in the config (the address to talk to other pseudo-calls from)")
	}
	if err := command.SyncDst(command.StrVal(dst)); err != nil {
		return err
	}
	return e.peer.Send(e.runCtx(), dst, m)
}

// SyncSend implements command.Env: SYNC, or HO with the subscriber's mapping context.
func (s *subscriber) SyncSend(kind, dst string) error {
	m := models.SyncMessage{Kind: kind, IMSI: s.imsi}
	if kind == models.SyncKindHO {
		c := s.ue.HandoverContext()
		m.Handover = &c
	}
	if err := s.e.syncSend(dst, m); err != nil {
		return err
	}
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: fmt.Sprintf("sync_send(%s) to %s: accepted", kind, dst)})
	return nil
}

// SyncSend implements command.Env: the node sends SYNC only.
func (n *node) SyncSend(kind, dst string) error {
	if kind != models.SyncKindSync {
		return fmt.Errorf("%s is a subscriber's (the node sends SYNC only)", kind)
	}
	if err := n.e.syncSend(dst, models.SyncMessage{Kind: kind}); err != nil {
		return err
	}
	n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: fmt.Sprintf("sync_send(%s) to %s: accepted", kind, dst)})
	return nil
}

// SyncSend implements command.Env; the console has no subject to send for.
func (consoleEnv) SyncSend(string, string) error {
	return errors.New("sync_send runs on a subscriber or the node")
}

// syncReceived runs on the subscriber: a HO first hands the mapping context
// over, then sync_receive(kind) of the current state (then any's) runs.
func (s *subscriber) syncReceived(m models.SyncMessage) {
	trigger := "sync_receive(" + m.Kind + ")"
	if m.Kind == models.SyncKindHO {
		in, err := s.ue.AdoptHandover(*m.Handover)
		if err != nil {
			s.e.publish(models.Event{Kind: models.KindError, Source: s.imsi, Text: fmt.Sprintf("%s from %s: the handover context: %v", trigger, m.From, err)})
			return
		}
		s.e.routes.unbind(s) // a new association at this pseudo-call
		s.dropTunnels()
		if in.PagingKey != "" {
			s.e.routes.bindPaging(in.PagingKey, s)
		}
		for _, t := range in.Notes {
			s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: t})
		}
		// The AMF's HandoverRequest comes to this subscriber, known by its uplink TEIDs.
		var teids []uint32
		for _, x := range m.Handover.Sessions {
			teids = append(teids, x.ULTEID)
		}
		s.e.routes.awaitHandover(s, teids)
	}
	s.lastMsg = trigger
	st, _ := s.e.scen.State(s.state)
	x, ok := st.Sync(m.Kind)
	block := s.state
	if !ok {
		x, ok = s.e.scen.Any.Sync(m.Kind)
		block = "any"
	}
	if !ok {
		s.e.publish(models.Event{Kind: models.KindWarn, Source: s.imsi, Handler: "unhandled",
			Text: fmt.Sprintf("%s from %s: no %s in %s or any", trigger, m.From, trigger, s.state)})
		return
	}
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Handler: block, State: s.state, Text: trigger + " from " + m.From})
	s.run(block+"/"+trigger, x.Body)
	s.afterHandler()
}

// syncReceived runs on the node: sync_receive(kind) of the node block.
func (n *node) syncReceived(m models.SyncMessage) {
	trigger := "sync_receive(" + m.Kind + ")"
	x, ok := n.e.scen.Node.Sync(m.Kind)
	if !ok {
		n.e.publish(models.Event{Kind: models.KindWarn, Source: models.SourceNode, Handler: "unhandled",
			Text: fmt.Sprintf("%s from %s: no %s in node", trigger, m.From, trigger)})
		return
	}
	n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Handler: "node", Text: trigger + " from " + m.From})
	n.run("node/"+trigger, x.Body)
}

// resetAssociations ends the UE associations an NG Reset names: all for
// the whole interface, else those of its list (by RAN- or AMF-UE-NGAP-ID).
func (e *Engine) resetAssociations(t ngapType.ResetType, why string) {
	var subs []*subscriber
	if t.Present == ngapType.ResetTypePresentPartOfNGInterface {
		for _, it := range t.PartOfNGInterface.List {
			var s *subscriber
			if it.HasRANUENGAPID {
				s = e.routes.byRAN(it.RANUENGAPID.Value)
			}
			if s == nil && it.HasAMFUENGAPID {
				s = e.routes.byAMF(it.AMFUENGAPID.Value)
			}
			if s != nil {
				subs = append(subs, s)
			}
		}
	} else {
		for _, a := range e.routes.associations() {
			if s := e.routes.byRAN(a.RANUENGAPID); s != nil {
				subs = append(subs, s)
			}
		}
	}
	for _, s := range subs {
		s.post(releaseMsg{why: why})
	}
}
