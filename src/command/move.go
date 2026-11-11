package command

// init registers move.
func init() {
	Register(Spec{
		Name:  "move",
		Args:  []Arg{{Name: "state", Kind: KindState}},
		Scope: ScopeUE,
		Where: InScenario | Interactive,
		Help:  "move the subscriber to a state and run its entry (forced from the CUI/GUI)",
		Run: func(env Env, args []Value) error {
			return env.Move(args[0].Str)
		},
	})
}
