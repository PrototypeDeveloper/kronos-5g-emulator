package scenario

import (
	"sort"
	"strings"

	"kronos/command"
	"kronos/common/diag"
	"kronos/models"
)

// Validate checks a parsed scenario against the signals and the command
// registry. sigs may be nil when signal.txt could not be read;
// signal names are then not checked.
func Validate(f *File, sigs *SignalSet, opt models.ValidateOptions) diag.List {
	v := &validator{
		f:     f,
		sigs:  sigs,
		vars:  map[string]*Var{},
		sent:  map[string]bool{},
		moved: map[string]bool{},
		noted: map[string]bool{},
	}
	v.checkVars()
	if len(f.States) > 0 {
		if _, ok := f.State("start"); !ok {
			v.ds.Errorf(f.States[0].Pos, "there is no initial state start (subscribers begin in start)")
		}
	}
	// Check each block's entry and handlers; seen catches repeated triggers.
	for _, b := range f.Blocks() {
		seen := map[string]*Handler{}
		if b.Entry != nil {
			v.checkCalls(b, b.Entry.Calls)
		}
		for _, h := range b.Handlers {
			v.checkTrigger(b, h, seen)
			v.checkCalls(b, h.Calls)
		}
	}
	v.reportUsage()
	// Note snap calls that do not run because base.snap is off.
	n, first := 0, diag.Pos{}
	f.Calls(func(_ *Block, c command.Call) {
		if c.Name == "snap" {
			if n == 0 {
				first = c.Pos
			}
			n++
		}
	})
	if n > 0 && !opt.SnapEnabled {
		v.ds.Infof(first, "base.snap is not true, so snap (%d place(s)) does not run", n)
	}
	return v.ds
}

// checkVars registers the declared variables, reporting duplicates.
func (v *validator) checkVars() {
	for _, x := range v.f.Vars {
		if prev, dup := v.vars[x.Name]; dup {
			v.ds.Errorf(x.Pos, "variable %s is declared twice (the first is on line %d)", x.Name, prev.Pos.Line)
			continue
		}
		v.vars[x.Name] = x
	}
}

// checkCalls checks each call against the command registry in the block's scope.
func (v *validator) checkCalls(b *Block, calls []command.Call) {
	ctx := command.Context{Where: command.InScenario, Scope: b.Scope()}
	for _, c := range calls {
		_, _, ds := command.Check(c, ctx, v)
		v.ds = append(v.ds, ds...)
	}
}

// checkTrigger resolves a receive() trigger and checks that it fits the block and is not repeated.
func (v *validator) checkTrigger(b *Block, h *Handler, seen map[string]*Handler) {
	tr := &h.Trigger
	if tr.Wildcard {
		v.ds.Errorf(tr.Pos, "receive(*) is not supported")
		return
	}
	info, ok := LookupNGAP(tr.NGAP)
	if !ok {
		v.ds.Errorf(tr.Pos, "unknown NGAP message %s%s", tr.NGAP, suggest(tr.NGAP, NGAPNames()))
		return
	}
	tr.NGAPInfo = info
	v.noteSpelling(tr.Pos, tr.NGAP, info.Name)
	key := info.Name
	// A NAS name is allowed only after an NGAP message that carries NAS.
	if tr.NAS != "" {
		nasPos := tr.Pos
		nasPos.Col += len([]rune(tr.NGAP)) + 1
		if !info.CarriesNAS {
			v.ds.Errorf(nasPos, "%s carries no NAS, so :%s cannot follow it", info.Name, tr.NAS)
			return
		}
		ni, ok := LookupNAS(tr.NAS)
		if !ok {
			v.ds.Errorf(nasPos, "unknown NAS message %s%s", tr.NAS, suggest(tr.NAS, NASNames()))
			return
		}
		tr.NASInfo = ni
		v.noteSpelling(nasPos, tr.NAS, ni.Name)
		key += ":" + ni.Name
	}
	// UE-associated messages belong in subscriber blocks, node ones in node (except Paging).
	switch {
	case b.Kind == BlockNode && info.Assoc == AssocUE:
		v.ds.Errorf(tr.Pos, "%s is UE-associated, so node cannot receive it (write it in a state block or any)", info.Name)
	case b.Kind != BlockNode && info.Assoc == AssocNode && info.Name != models.PagingName:
		v.ds.Errorf(tr.Pos, "%s belongs to the node, so a subscriber block (%s) cannot receive it (write it in node)", info.Name, b.Name)
	}
	if prev, dup := seen[key]; dup {
		v.ds.Errorf(tr.Pos, "the same trigger %s is also on line %[3]d of the %[2]s block", key, b.Name, prev.Pos.Line)
		return
	}
	seen[key] = h
}

// noteSpelling notes, once per spelling, a name that matched only by ignoring case.
func (v *validator) noteSpelling(pos diag.Pos, written, canonical string) {
	if written == canonical || v.noted[written] {
		return
	}
	v.noted[written] = true
	v.ds.Infof(pos, "%s is matched as %s (case does not matter)", written, canonical)
}

// reportUsage notes the signals never sent and the states nothing moves to.
func (v *validator) reportUsage() {
	if v.sigs != nil {
		if names := unusedSignals(v.sigs, v.sent); len(names) > 0 {
			v.ds.Infof(diag.Pos{File: v.f.Path}, "signals the scenario does not send (sent from the CUI/GUI to start something): %s", strings.Join(names, ", "))
		}
	}
	var idle []string
	for _, b := range v.f.States {
		if b.Name != "start" && !v.moved[b.Name] {
			idle = append(idle, b.Name)
		}
	}
	if len(idle) > 0 {
		v.ds.Infof(diag.Pos{File: v.f.Path}, "states that nothing moves to: %s", strings.Join(idle, ", "))
	}
}

// suggest returns a " (did you mean X?)" hint, or "" when no name is close.
func suggest(name string, candidates []string) string {
	if s := command.Suggest(name, candidates); s != "" {
		return " (did you mean " + s + "?)"
	}
	return ""
}

// Signal implements command.Resolver.
func (v *validator) Signal(name string) (command.Scope, bool) {
	if v.sigs == nil {
		return command.ScopeNode | command.ScopeUE, true
	}
	sig, ok := v.sigs.Lookup(name)
	if !ok {
		return 0, false
	}
	v.sent[name] = true
	return sig.Scope, true
}

// State implements command.Resolver.
func (v *validator) State(name string) bool {
	_, ok := v.f.State(name)
	if ok {
		v.moved[name] = true
	}
	return ok
}

// Var implements command.Resolver.
func (v *validator) Var(name string) (command.VarInfo, bool) {
	x, ok := v.vars[name]
	if !ok {
		return command.VarInfo{}, false
	}
	return command.VarInfo{Type: x.Init.Type, Public: x.Public}, true
}

// Subscriber implements command.Resolver; scenarios never name a subscriber.
func (v *validator) Subscriber(string) bool { return false }

// Names implements command.Resolver.
func (v *validator) Names(k command.Kind) []string {
	var out []string
	switch k {
	case command.KindSignal:
		out = v.sigs.Names()
	case command.KindState:
		for _, b := range v.f.States {
			out = append(out, b.Name)
		}
	case command.KindValue:
		for n := range v.vars {
			out = append(out, n)
		}
		sort.Strings(out)
	}
	return out
}
