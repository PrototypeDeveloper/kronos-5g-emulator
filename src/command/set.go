package command

import "kronos/common/diag"

// init registers set.
func init() {
	Register(Spec{
		Name:  "set",
		Args:  []Arg{{Name: "var = value", Kind: KindAssign}},
		Scope: ScopeByAssign,
		Where: InScenario | Interactive,
		Help: "assign a variable: set(a = 2), set(a = a + 5), set(s = s + \"add\"); the value is an expression " +
			"(+ - * / %, comparisons, && || !, ( )). A public variable is one value for all; on the CUI/GUI a private one is set on the subscribers",
		RunCall: runSet,
	})
}

// checkAssign checks var = value: the variable belongs to the subject and the value fits its type.
func checkAssign(c Call, e Expr, r Resolver, subject Scope, ds *diag.List) {
	vi, _ := r.Var(e.Name)
	if subject == ScopeNode && !vi.Public {
		ds.Errorf(e.Pos, "%s in node: %s is private (it belongs to each subscriber)", c.Name, e.Name)
	}
	if t := checkExpr(*e.Y, r, subject, ds); t != TypeNone && !assignable(t, vi.Type) {
		ds.Errorf(e.Y.Pos, "%s: %s is of type %s, so %s (%s) cannot be assigned to it", c.Name, e.Name, vi.Type, e.Y, t)
	}
}

// runSet evaluates the value and stores it, converted to the variable's type.
func runSet(env Env, c Call) error {
	a := c.Args[0]
	return env.Assign(a.Name, func() (Value, error) {
		v, err := EvalExpr(env, *a.Y)
		if err != nil {
			return Value{}, err
		}
		cur, err := env.Var(a.Name)
		if err != nil {
			return Value{}, err
		}
		return Convert(v, cur.Type)
	})
}
