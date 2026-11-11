package command

// init registers snap.
func init() {
	Register(Spec{
		Name:  "snap",
		Args:  []Arg{{Name: "format", Kind: KindFormat}, {Name: "value", Kind: KindValue, Optional: true, Variadic: true}},
		Scope: ScopeNode | ScopeUE,
		Where: InScenario | Interactive,
		Help:  "print Printf style on the console (only when base.snap is true)",
		Check: checkSnap,
		Run: func(env Env, args []Value) error {
			env.Print(Format(args[0].Str, args[1:]))
			return nil
		},
	})
}
