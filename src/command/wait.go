package command

import (
	"fmt"

	"kronos/common/diag"
)

// init registers wait.
func init() {
	Register(Spec{
		Name:  "wait",
		Args:  []Arg{{Name: "time", Kind: KindValue}},
		Scope: ScopeConsole,
		Where: Interactive,
		Help:  "wait for a while (in a batch, to wait for answers; e.g. wait(3s))",
		Check: checkWait,
		Run: func(env Env, args []Value) error {
			if args[0].Type != TypeDuration || args[0].Dur <= 0 || args[0].Dur > waitMaxTime {
				return fmt.Errorf("wait takes a duration up to 1h (e.g. wait(3s); found %s)", args[0].Literal())
			}
			return env.Sleep(args[0].Dur)
		},
	})
}

// checkWait rejects a literal time that is not a positive duration up to 1h.
func checkWait(c Call, _ Resolver, _ Scope) diag.List {
	var ds diag.List
	if e := c.Args[0]; e.Kind == ExprLiteral && (e.Val.Type != TypeDuration || e.Val.Dur <= 0 || e.Val.Dur > waitMaxTime) {
		ds.Errorf(e.Pos, "wait takes a duration up to 1h (e.g. wait(3s); found %s)", e)
	}
	return ds
}
