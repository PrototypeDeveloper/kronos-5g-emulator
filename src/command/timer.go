package command

import "kronos/common/diag"

// init registers timer and stoptimer.
func init() {
	Register(Spec{
		Name:  "timer",
		Args:  []Arg{{Name: "var", Kind: KindVar}},
		Scope: ScopeByVar,
		Where: InScenario | Interactive,
		Help: "start the timer of a duration variable for its current value (again from the start if it runs): " +
			"when it runs out, expire(var) of the state the subscriber is then in runs (then any's); a public variable is a timer of the node (node's expire)",
		Check: checkTimerVar,
		Run:   func(env Env, args []Value) error { return env.Timer(args[0].Str) },
	})
	Register(Spec{
		Name:  "stoptimer",
		Args:  []Arg{{Name: "var", Kind: KindVar}},
		Scope: ScopeByVar,
		Where: InScenario | Interactive,
		Help:  "stop the timer of a variable (nothing happens if it does not run)",
		Check: checkTimerVar,
		Run:   func(env Env, args []Value) error { return env.StopTimer(args[0].Str) },
	})
}

// checkTimerVar: a timer is a duration variable (whose owner, the node or
// the subscriber, ScopeByVar has already matched with the block).
func checkTimerVar(c Call, r Resolver, _ Scope) diag.List {
	var ds diag.List
	name := c.Args[0].Name
	if vi, _ := r.Var(name); vi.Type != TypeDuration {
		ds.Errorf(c.Args[0].Pos, "%s: %s is not a duration variable (declare it like private %s = 5s)", c.Name, name, name)
	}
	return ds
}
