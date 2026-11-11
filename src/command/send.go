package command

// init registers send.
func init() {
	Register(Spec{
		Name:  "send",
		Args:  []Arg{{Name: "signal", Kind: KindSignal}},
		Scope: ScopeBySignal,
		Where: InScenario | Interactive,
		Help:  "send a signal of signal.txt with its dynamic fields filled in: UE signals from each subscriber, node signals from the node",
		Run: func(env Env, args []Value) error {
			return env.Send(args[0].Str)
		},
	})
}
