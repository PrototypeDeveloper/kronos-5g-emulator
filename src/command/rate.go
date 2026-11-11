package command

import (
	"errors"

	"kronos/traffic"
)

// EvalLiteral is the value a KindRate argument has before it runs: the word
// 100M, or a literal.
func EvalLiteral(e Expr) Value {
	if e.Kind == ExprIdent {
		return StrVal(e.Name)
	}
	return e.Val
}

// RateOf reads a bit rate argument: 100M, 1G, 500k ("1.5G" in quotes for a
// fraction) or bit/s as an int.
func RateOf(v Value) (uint64, error) {
	switch v.Type {
	case TypeInt:
		if v.Int <= 0 {
			return 0, errors.New("the rate must be positive")
		}
		return uint64(v.Int), nil
	case TypeString:
		return traffic.ParseRate(v.Str)
	case TypeDuration:
		return 0, errors.New("write the rate like 100M (100m reads as 100 minutes: use a capital M)")
	}
	return 0, errors.New("write the rate like 100M, 1G or 500k")
}
