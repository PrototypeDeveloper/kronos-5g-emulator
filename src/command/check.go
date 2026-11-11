package command

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"kronos/common/diag"
)

// Resolver answers name lookups during validation. The scenario validator
// implements it for scenario.txt, and the engine for the command line.
type Resolver interface {
	// Signal reports whether the signal exists and whether it is sent by the
	// node (ScopeNode) or per subscriber (ScopeUE). Both bits mean unknown.
	Signal(name string) (Scope, bool)
	State(name string) bool
	Var(name string) (VarInfo, bool)
	// Subscriber reports whether imsi belongs to a configured subscriber.
	Subscriber(imsi string) bool
	// Names lists the defined names of a kind, for "did you mean" hints.
	// KindValue lists the variables.
	Names(k Kind) []string
}

// Check validates a call against the registry. It returns the spec and the
// subject the call runs on: ScopeUE, ScopeNode or ScopeConsole (0 on error).
func Check(c Call, ctx Context, r Resolver) (Spec, Scope, diag.List) {
	var ds diag.List
	// The command must exist, be allowed here and get a valid number of arguments.
	s, ok := Lookup(c.Name)
	if !ok {
		switch {
		case c.Name == TargetKeyword && ctx.Where == InScenario:
			ds.Errorf(c.Pos, "target() cannot be used in a scenario (commands run on the subscriber whose handler runs)")
		case c.Name == TargetKeyword:
			ds.Errorf(c.Pos, "target goes before a command (e.g. target(208930000000001).send(signal))")
		default:
			ds.Errorf(c.Pos, "unknown command %s%s", c.Name, hint(c.Name, Names()))
		}
		return Spec{}, 0, ds
	}
	if s.Where&ctx.Where == 0 {
		if ctx.Where == InScenario {
			ds.Errorf(c.Pos, "%s cannot be used in a scenario (CUI/GUI only)", c.Name)
		} else {
			ds.Errorf(c.Pos, "%s cannot be used on the CUI/GUI (scenario only)", c.Name)
		}
		return s, 0, ds
	}
	if min, max := s.argCount(); len(c.Args) < min || (max >= 0 && len(c.Args) > max) {
		ds.Errorf(c.Pos, "wrong number of arguments for %s (usage: %s)", c.Name, s.Usage())
		return s, 0, ds
	}

	// Names and the format first: for send, the signal decides the scope.
	scope := s.Scope
	for i, e := range c.Args {
		a := s.argAt(i)
		switch a.Kind {
		case KindSignal:
			if e.Kind != ExprIdent {
				ds.Errorf(e.Pos, "%s takes a signal name (found %s)", c.Name, e)
				continue
			}
			sc, ok := r.Signal(e.Name)
			if !ok {
				ds.Errorf(e.Pos, "signal %s is not in signal.txt%s", e.Name, hint(e.Name, r.Names(KindSignal)))
				continue
			}
			if scope&ScopeBySignal != 0 {
				scope = sc
			}
		case KindState:
			if e.Kind != ExprIdent {
				ds.Errorf(e.Pos, "%s takes a state name (found %s)", c.Name, e)
			} else if !r.State(e.Name) {
				ds.Errorf(e.Pos, "state %s is not defined in scenario.txt%s", e.Name, hint(e.Name, r.Names(KindState)))
			}
		case KindCommand:
			if e.Kind != ExprIdent {
				ds.Errorf(e.Pos, "%s takes a command name (found %s)", c.Name, e)
			} else if _, ok := Lookup(e.Name); !ok {
				ds.Errorf(e.Pos, "unknown command %s%s", e.Name, hint(e.Name, Names()))
			}
		case KindFormat:
			if e.Kind != ExprLiteral || e.Val.Type != TypeString {
				ds.Errorf(e.Pos, "%s: %s must be a string (\"…\")", c.Name, a.Name)
			}
		case KindChoice:
			if e.Kind != ExprIdent || !slices.Contains(a.Choices, e.Name) {
				ds.Errorf(e.Pos, "%s: %s must be one of %s (found %s)", c.Name, a.Name, strings.Join(a.Choices, ", "), e)
			}
		case KindRate:
			if v := EvalLiteral(e); a.ZeroRate && v.Type == TypeInt && v.Int == 0 {
				break
			}
			if _, err := RateOf(EvalLiteral(e)); err != nil {
				ds.Errorf(e.Pos, "%s: %s: %v", c.Name, a.Name, err)
			}
		}
	}
	if ds.HasErrors() {
		return s, 0, ds
	}

	// The subject: on the command line it follows from the scope; in a
	// scenario the block must be one the command may run in.
	subject := ctx.Scope
	if ctx.Where == Interactive {
		switch {
		case scope&ScopeUE != 0:
			subject = ScopeUE
		case scope&ScopeNode != 0:
			subject = ScopeNode
		default:
			subject = ScopeConsole
		}
	} else if scope&ctx.Scope == 0 {
		ds.Errorf(c.Pos, "%s", scopeMismatch(c, s, ctx.Scope))
		return s, 0, ds
	}

	// Variables and builtins must exist and belong to the subject.
	for i, e := range c.Args {
		if s.argAt(i).Kind != KindValue {
			continue
		}
		switch e.Kind {
		case ExprIdent:
			vi, ok := r.Var(e.Name)
			if !ok {
				ds.Errorf(e.Pos, "variable %s is not declared in var%s", e.Name, hint(e.Name, r.Names(KindValue)))
			} else if subject == ScopeNode && !vi.Public {
				ds.Errorf(e.Pos, "private variable %s cannot be used in node (it belongs to each subscriber)", e.Name)
			}
		case ExprBuiltin:
			bi, ok := LookupBuiltin(e.Name)
			if !ok {
				ds.Errorf(e.Pos, "unknown builtin $%s%s", e.Name, hint(e.Name, builtinNames()))
			} else if bi.Scope&subject == 0 {
				ds.Errorf(e.Pos, "$%s belongs to each subscriber, so node cannot use it", e.Name)
			}
		}
	}
	// Command-specific checks run only on an otherwise valid call.
	if s.Check != nil && !ds.HasErrors() {
		ds = append(ds, s.Check(c, r, subject)...)
	}
	if ds.HasErrors() {
		return s, 0, ds
	}
	return s, subject, ds
}

// scopeMismatch explains why a command cannot be used in a block of scope block.
func scopeMismatch(c Call, s Spec, block Scope) string {
	if s.Scope&ScopeBySignal != 0 && len(c.Args) > 0 {
		if block == ScopeNode {
			return fmt.Sprintf("%s is a UE signal, so the node block cannot send it (write it in a state block or any)", c.Args[0].Name)
		}
		return fmt.Sprintf("%s is a node signal, so subscriber blocks (states, any) cannot send it (write it in the node block)", c.Args[0].Name)
	}
	if block == ScopeNode {
		return fmt.Sprintf("%s cannot be used in the node block (it acts on a subscriber)", c.Name)
	}
	return fmt.Sprintf("%s cannot be used in subscriber blocks (states, any): write it in the node block", c.Name)
}

// CheckLine validates a CUI/GUI command line, including its target.
func CheckLine(l Line, r Resolver) (Spec, Scope, diag.List) {
	s, subject, ds := Check(l.Call, Context{Where: Interactive}, r)
	if l.Target == "" || ds.HasErrors() {
		return s, subject, ds
	}
	switch {
	case subject != ScopeUE && s.Scope&ScopeBySignal != 0:
		ds.Errorf(l.TargetPos, "%s is a node signal, so it cannot take target", l.Call.Args[0].Name)
	case subject != ScopeUE:
		ds.Errorf(l.TargetPos, "%s does not act on subscribers, so it cannot take target", l.Call.Name)
	case len(l.Target) != imsiDigits || !allDigits(l.Target):
		ds.Errorf(l.TargetPos, "target() takes a 15-digit IMSI (%s)", l.Target)
	case !r.Subscriber(l.Target):
		ds.Errorf(l.TargetPos, "there is no subscriber with IMSI %s", l.Target)
	}
	if ds.HasErrors() {
		return s, 0, ds
	}
	return s, subject, ds
}

// builtinNames returns the builtin names, sorted, for hints.
func builtinNames() []string {
	out := make([]string, 0, len(builtins))
	for n := range builtins {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}

// hint returns a " (did you mean X?)" suffix for an error message, or "".
func hint(name string, candidates []string) string {
	if s := Suggest(name, candidates); s != "" {
		return " (did you mean " + s + "?)"
	}
	return ""
}

// Suggest returns the candidate closest to name (ignoring case), or "" when
// nothing is close enough to be a plausible typo.
func Suggest(name string, candidates []string) string {
	ln := []rune(strings.ToLower(name))
	best, bestD := "", -1
	for _, c := range candidates {
		d := levenshtein(ln, []rune(strings.ToLower(c)))
		if bestD < 0 || d < bestD {
			best, bestD = c, d
		}
	}
	limit := len(ln) / suggestRunesPerEdit
	if limit < suggestMinEdits {
		limit = suggestMinEdits
	}
	if bestD >= 0 && bestD <= limit {
		return best
	}
	return ""
}

// levenshtein returns the edit distance between a and b.
func levenshtein(a, b []rune) int {
	prev := make([]int, len(b)+1)
	cur := make([]int, len(b)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(a); i++ {
		cur[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev, cur = cur, prev
	}
	return prev[len(b)]
}
