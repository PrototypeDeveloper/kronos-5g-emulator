package command

import (
	"strconv"
	"time"
)

// Type is the type of a scenario value. It is fixed by the literal that
// declares a variable (`public a = 1` is an int).
type Type uint8

// String returns the type name used in messages.
func (t Type) String() string {
	switch t {
	case TypeInt:
		return "int"
	case TypeFloat:
		return "float"
	case TypeString:
		return "string"
	case TypeDuration:
		return "duration"
	case TypeBool:
		return "bool"
	case TypeAny:
		return "IE value"
	default:
		return "none"
	}
}

// IntVal, FloatVal, StrVal and DurVal make a Value of each type.
func IntVal(v int64) Value         { return Value{Type: TypeInt, Int: v} }
func FloatVal(v float64) Value     { return Value{Type: TypeFloat, Float: v} }
func StrVal(v string) Value        { return Value{Type: TypeString, Str: v} }
func DurVal(v time.Duration) Value { return Value{Type: TypeDuration, Dur: v} }

// BoolVal makes the Value of a comparison.
func BoolVal(v bool) Value { return Value{Type: TypeBool, Bool: v} }

// String returns the value as snap prints it with %v.
func (v Value) String() string {
	switch v.Type {
	case TypeInt:
		return strconv.FormatInt(v.Int, 10)
	case TypeFloat:
		return strconv.FormatFloat(v.Float, 'g', -1, 64)
	case TypeString:
		return v.Str
	case TypeDuration:
		return v.Dur.String()
	case TypeBool:
		return strconv.FormatBool(v.Bool)
	default:
		return ""
	}
}

// Literal returns the value as it would be written in a scenario (strings quoted).
func (v Value) Literal() string {
	if v.Type == TypeString {
		return strconv.Quote(v.Str)
	}
	return v.String()
}

// Interface returns the Go value handed to fmt by snap.
func (v Value) Interface() any {
	switch v.Type {
	case TypeInt:
		return v.Int
	case TypeFloat:
		return v.Float
	case TypeString:
		return v.Str
	case TypeDuration:
		return v.Dur
	case TypeBool:
		return v.Bool
	default:
		return nil
	}
}
