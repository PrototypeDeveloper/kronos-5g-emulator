package command

// init registers linkup, linkdown and reload.
func init() {
	Register(Spec{
		Name:  "linkup",
		Scope: ScopeConsole,
		Where: Interactive,
		Help:  "connect SCTP (then the entry of node runs)",
		Run:   func(env Env, _ []Value) error { return env.LinkUp() },
	})
	Register(Spec{
		Name:  "linkdown",
		Scope: ScopeConsole,
		Where: Interactive,
		Help:  "disconnect SCTP (the subscribers go back to their initial state)",
		Run:   func(env Env, _ []Value) error { return env.LinkDown() },
	})
	Register(Spec{
		Name:  "reload",
		Scope: ScopeConsole,
		Where: Interactive,
		Help:  "read signal.txt again and send its signals from now on (with an error in it, nothing changes)",
		Run:   func(env Env, _ []Value) error { return env.Reload() },
	})
}
