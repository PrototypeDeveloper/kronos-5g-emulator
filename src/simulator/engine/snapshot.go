package engine

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"kronos/command"
	"kronos/models"
	"kronos/scenario"
)

// The read-only views of the web console. They return
// snapshots and never wait for an actor.

// Started is when the engine was made: a GUI tells a restarted process by it.
func (e *Engine) Started() time.Time { return e.started }

// Overview returns the dashboard: the node, the subscribers by state and the counters.
func (e *Engine) Overview() models.Overview {
	o := models.Overview{
		Role: e.role.Name(), Config: e.cfg.Path, Started: e.started, Snap: e.snap,
		Node: e.Status(), Subscribers: len(e.subs),
		States: map[string]int{}, RM: map[string]int{}, CM: map[string]int{},
		Counters: e.bus.Counters(),
	}
	for _, s := range e.subs {
		st := s.status()
		o.States[st.State]++
		o.RM[st.RM]++
		o.CM[st.CM]++
		for _, ss := range st.Sessions {
			if ss.Active {
				o.Sessions++
			}
		}
	}
	return o
}

// UEs returns the snapshots of the subscribers in IMSI order, limit of them
// from offset, optionally only those in a scenario state, and how many match.
func (e *Engine) UEs(offset, limit int, state string) ([]models.UEStatus, int) {
	var out []models.UEStatus
	total := 0
	for _, s := range e.subs {
		st := s.status()
		if state != "" && st.State != state {
			continue
		}
		if total >= offset && (limit <= 0 || len(out) < limit) {
			out = append(out, st)
		}
		total++
	}
	return out, total
}

// UE returns the snapshot of one subscriber.
func (e *Engine) UE(imsi string) (models.UEStatus, bool) {
	i, ok := e.cfg.SubscriberIndex(imsi)
	if !ok || i >= len(e.subs) {
		return models.UEStatus{}, false
	}
	return e.subs[i].status(), true
}

// Signals describes the templates of signal.txt in file order.
func (e *Engine) Signals() []models.SignalInfo {
	out := make([]models.SignalInfo, 0, len(e.signals().List))
	for _, s := range e.signals().List {
		si := models.SignalInfo{Name: s.Name, Scope: scopeName(s.Scope), Mode: s.Mode.String(), Label: s.Label(), Line: s.Pos.Line, Hex: hex.EncodeToString(s.Bytes)}
		if s.NGAP != nil {
			si.NGAP = s.NGAP.Name
		}
		out = append(out, si)
	}
	return out
}

// scopeName names a signal scope for the web console.
func scopeName(s command.Scope) string {
	switch s {
	case command.ScopeNode:
		return "node"
	case command.ScopeUE:
		return "UE"
	}
	return "?"
}

// Scenario describes the scenario: variables (public ones with their
// current value) and blocks in file order.
func (e *Engine) Scenario() models.ScenarioInfo {
	info := models.ScenarioInfo{Path: e.scen.Path, Vars: e.Vars()}
	for _, b := range e.scen.Blocks() {
		bi := models.BlockInfo{Name: b.Name, Line: b.Pos.Line}
		switch b.Kind {
		case scenario.BlockNode:
			bi.Kind = "node"
		case scenario.BlockAny:
			bi.Kind = "any"
		default:
			bi.Kind = "state"
		}
		if b.Entry != nil {
			bi.Entry = scenario.StmtLines(b.Entry.Body)
		}
		for _, h := range b.Handlers {
			bi.Handlers = append(bi.Handlers, models.HandlerInfo{Trigger: h.Trigger.String(), Line: h.Pos.Line, Calls: scenario.StmtLines(h.Body)})
		}
		for _, x := range b.Expires {
			bi.Handlers = append(bi.Handlers, models.HandlerInfo{Kind: "expire", Trigger: x.Var, Line: x.Pos.Line, Calls: scenario.StmtLines(x.Body)})
		}
		for _, x := range b.Syncs {
			bi.Handlers = append(bi.Handlers, models.HandlerInfo{Kind: "sync_receive", Trigger: x.Kind, Line: x.Pos.Line, Calls: scenario.StmtLines(x.Body)})
		}
		info.Blocks = append(info.Blocks, bi)
	}
	return info
}

// Vars returns the scenario variables by name: the current value of the
// public ones and the initial value of the private ones.
func (e *Engine) Vars() []models.VarValue {
	out := make([]models.VarValue, 0, len(e.scen.Vars))
	for _, v := range e.scen.Vars {
		val := v.Init
		if v.Public {
			if cur, ok := e.vars.get(v.Name); ok {
				val = cur
			}
		}
		out = append(out, models.VarValue{Name: v.Name, Type: v.Init.Type.String(), Public: v.Public, Value: val.Literal()})
	}
	sort.Slice(out, func(i, j int) bool { return strings.Compare(out[i].Name, out[j].Name) < 0 })
	return out
}

// DecodeEvent returns a kept event and the decoded tree of its PDU (nil when
// the event carries no PDU). The plaintext kept with the event shows the
// ciphered NAS.
func (e *Engine) DecodeEvent(seq uint64) (models.Event, *models.TreeNode, error) {
	ev, ok := e.bus.Event(seq)
	if !ok {
		return ev, nil, fmt.Errorf("event %d is no longer kept (the oldest ones are dropped)", seq)
	}
	if ev.Hex == "" {
		return ev, nil, nil
	}
	pdu, err := hex.DecodeString(ev.Hex)
	if err != nil {
		return ev, nil, err
	}
	plain := make([][]byte, len(ev.Plain))
	for i, p := range ev.Plain {
		if p != "" {
			plain[i], _ = hex.DecodeString(p)
		}
	}
	tree, err := scenario.Tree(pdu, plain, false)
	return ev, tree, err
}

// DecodeSignal returns a template and its decoded tree.
func (e *Engine) DecodeSignal(name string) (models.SignalInfo, *models.TreeNode, error) {
	for _, si := range e.Signals() {
		if si.Name != name {
			continue
		}
		s, _ := e.signals().Lookup(name)
		tree, err := scenario.Tree(s.Bytes, nil, true)
		return si, tree, err
	}
	return models.SignalInfo{}, nil, fmt.Errorf("signal %s is not in signal.txt", name)
}
