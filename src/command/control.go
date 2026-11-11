package command

import (
	"errors"
	"fmt"
	"time"

	"kronos/common/diag"
)

// ErrInterrupted ends a control() wait cut short by the link going down (or
// the exit); the rest of the handler is then skipped.
var ErrInterrupted = errors.New("the wait was cut short (the link went down)")

// init registers control.
func init() {
	Register(Spec{
		Name:  "control",
		Args:  []Arg{{Name: "per second", Kind: KindValue}},
		Scope: ScopeUE,
		Where: InScenario,
		Help: "pace the scenario: at most n passes per second through this control(), counted over all subscribers " +
			"and evenly spaced (control(5): one every 0.2 s); a subscriber waits for its turn, and so do its messages. control(0): no limit",
		Check:   checkControl,
		RunCall: runControl,
	})
}

// controlRate returns the passes per second of v, checking the range.
func controlRate(v Value) (float64, error) {
	var r float64
	switch v.Type {
	case TypeInt:
		r = float64(v.Int)
	case TypeFloat:
		r = v.Float
	default:
		return 0, fmt.Errorf("the passes per second must be a number, e.g. control(5) (found the %s %s)", v.Type, v.Literal())
	}
	if r < 0 || r > controlMaxRate {
		return 0, fmt.Errorf("the passes per second must be from 0 (no limit) to %d (found %s)", controlMaxRate, v.Literal())
	}
	return r, nil
}

// checkControl checks a literal rate, and the type of an expression.
func checkControl(c Call, r Resolver, _ Scope) diag.List {
	var ds diag.List
	e := c.Args[0]
	if e.Kind == ExprLiteral {
		if _, err := controlRate(e.Val); err != nil {
			ds.Errorf(e.Pos, "control: %v", err)
		}
	} else if t := exprType(e, r); t != TypeNone && !isNum(t) {
		ds.Errorf(e.Pos, "control takes the passes per second, e.g. control(5) (%s is of type %s)", e, t)
	}
	return ds
}

// runControl waits for the turn at this control(); its position in the
// scenario tells it from the others.
func runControl(env Env, c Call) error {
	v, err := EvalExpr(env, c.Args[0])
	if err != nil {
		return err
	}
	rate, err := controlRate(v)
	if err != nil || rate == 0 {
		return err
	}
	return env.Control(c.Pos.String(), time.Duration(float64(time.Second)/rate))
}
