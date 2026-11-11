package command

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"kronos/common/diag"
)

// Types of expressions. Variables keep the type of their initial value; a
// comparison gives a bool, which only if and && || ! take.

// isNum reports whether t is int or float.
func isNum(t Type) bool { return t == TypeInt || t == TypeFloat }

// unaryType returns the type of op x.
func unaryType(op string, x Type) (Type, error) {
	switch {
	case op == "-" && (isNum(x) || x == TypeDuration):
		return x, nil
	case op == "!" && x == TypeBool:
		return TypeBool, nil
	case op == "!":
		return TypeNone, fmt.Errorf("! takes a comparison such as !(a == 1) (found %s)", x)
	}
	return TypeNone, fmt.Errorf("- cannot be applied to %s", x)
}

// binaryType returns the type of x op y. Mixing int and float gives a float;
// + with a string joins the text of both sides.
func binaryType(op string, x, y Type) (Type, error) {
	switch op {
	case "&&", "||":
		if x == TypeBool && y == TypeBool {
			return TypeBool, nil
		}
		return TypeNone, fmt.Errorf("%s joins comparisons such as a == 1 (found %s and %s)", op, x, y)
	case "==", "!=":
		if x == y || isNum(x) && isNum(y) {
			return TypeBool, nil
		}
	case "<", "<=", ">", ">=":
		if x == y && x != TypeBool || isNum(x) && isNum(y) {
			return TypeBool, nil
		}
	case "+", "-", "*", "/", "%":
		if t := arithType(op, x, y); t != TypeNone {
			return t, nil
		}
	}
	return TypeNone, fmt.Errorf("%s cannot be applied to %s and %s", op, x, y)
}

// arithType returns the type of an arithmetic x op y, or TypeNone.
func arithType(op string, x, y Type) Type {
	switch {
	case op == "+" && (x == TypeString || y == TypeString) && x != TypeBool && y != TypeBool:
		return TypeString
	case x == TypeInt && y == TypeInt:
		return TypeInt
	case op == "%":
		return TypeNone
	case isNum(x) && isNum(y):
		return TypeFloat
	case x == TypeDuration && y == TypeDuration && (op == "+" || op == "-"):
		return TypeDuration
	case x == TypeDuration && y == TypeDuration && op == "/":
		return TypeFloat // the ratio
	case x == TypeDuration && isNum(y) && (op == "*" || op == "/"):
		return TypeDuration
	case isNum(x) && y == TypeDuration && op == "*":
		return TypeDuration
	}
	return TypeNone
}

// assignable reports whether a value of type from can be stored in a variable of type to.
func assignable(from, to Type) bool { return from == to || from == TypeInt && to == TypeFloat }

// Convert returns v as a value of the variable type to (an int becomes a float).
func Convert(v Value, to Type) (Value, error) {
	switch {
	case v.Type == to:
		return v, nil
	case v.Type == TypeInt && to == TypeFloat:
		return FloatVal(float64(v.Int)), nil
	}
	return Value{}, fmt.Errorf("%s cannot be stored in a variable of type %s", v.Type, to)
}

// checkExpr checks the names in e for the subject and returns the static
// type of e (TypeNone after an error, which it adds to ds).
func checkExpr(e Expr, r Resolver, subject Scope, ds *diag.List) Type {
	switch e.Kind {
	case ExprLiteral:
		return e.Val.Type
	case ExprIdent:
		vi, ok := r.Var(e.Name)
		switch {
		case !ok:
			ds.Errorf(e.Pos, "variable %s is not declared in var%s", e.Name, hint(e.Name, r.Names(KindValue)))
			return TypeNone
		case !vi.Public && subject == ScopeNode:
			ds.Errorf(e.Pos, "private variable %s cannot be used in node (it belongs to each subscriber)", e.Name)
		case !vi.Public && subject == ScopeConsole:
			ds.Errorf(e.Pos, "private variable %s cannot be used here (it belongs to each subscriber)", e.Name)
		}
		return vi.Type
	case ExprBuiltin:
		bi, ok := LookupBuiltin(e.Name)
		switch {
		case !ok:
			ds.Errorf(e.Pos, "unknown builtin $%s%s", e.Name, hint(e.Name, builtinNames()))
			return TypeNone
		case bi.Scope&subject != 0:
		case subject == ScopeNode:
			ds.Errorf(e.Pos, "$%s belongs to each subscriber, so node cannot use it", e.Name)
		default:
			ds.Errorf(e.Pos, "$%s cannot be used here (it belongs to a subscriber or the node)", e.Name)
		}
		return bi.Type
	case ExprParen:
		return checkExpr(*e.X, r, subject, ds)
	case ExprUnary:
		x := checkExpr(*e.X, r, subject, ds)
		if x == TypeNone {
			return TypeNone
		}
		t, err := unaryType(e.Op, x)
		if err != nil {
			ds.Errorf(e.Pos, "%v", err)
		}
		return t
	case ExprBinary:
		x := checkExpr(*e.X, r, subject, ds)
		y := checkExpr(*e.Y, r, subject, ds)
		if x == TypeNone || y == TypeNone {
			return TypeNone
		}
		t, err := binaryType(e.Op, x, y)
		if err != nil {
			ds.Errorf(e.Pos, "%v", err)
		}
		return t
	case ExprAssign:
		ds.Errorf(e.Pos, "= assigns a variable only in set(%s = …); compare with ==", e.Name)
	}
	return TypeNone
}

// CheckCond checks the condition of an if: it must be a comparison (a bool).
func CheckCond(e Expr, r Resolver, subject Scope) diag.List {
	var ds diag.List
	if t := checkExpr(e, r, subject, &ds); t != TypeNone && t != TypeBool {
		ds.Errorf(e.Pos, "the condition of if must be a comparison such as a == 1 (%s is of type %s)", e, t)
	}
	return ds
}

// exprType returns the static type of an argument, or TypeNone if unknown.
func exprType(e Expr, r Resolver) Type {
	var ds diag.List
	return checkExpr(e, r, ScopeNode|ScopeUE|ScopeConsole, &ds)
}

// EvalExpr evaluates e on env. && and || evaluate their right side only when needed.
func EvalExpr(env Env, e Expr) (Value, error) {
	switch e.Kind {
	case ExprLiteral:
		return e.Val, nil
	case ExprIdent:
		return env.Var(e.Name)
	case ExprBuiltin:
		return env.Builtin(e.Name)
	case ExprParen:
		return EvalExpr(env, *e.X)
	case ExprUnary:
		x, err := EvalExpr(env, *e.X)
		if err != nil {
			return Value{}, err
		}
		return applyUnary(e.Op, x)
	case ExprBinary:
		x, err := EvalExpr(env, *e.X)
		if err != nil {
			return Value{}, err
		}
		if x.Type == TypeBool && (e.Op == "&&" && !x.Bool || e.Op == "||" && x.Bool) {
			return x, nil
		}
		y, err := EvalExpr(env, *e.Y)
		if err != nil {
			return Value{}, err
		}
		return applyBinary(e.Op, x, y)
	}
	return Value{}, fmt.Errorf("%s cannot be evaluated", e)
}

// EvalCond evaluates the condition of an if.
func EvalCond(env Env, e Expr) (bool, error) {
	v, err := EvalExpr(env, e)
	if err != nil {
		return false, err
	}
	if v.Type != TypeBool {
		return false, fmt.Errorf("the condition %s is of type %s, not a comparison", e, v.Type)
	}
	return v.Bool, nil
}

// applyUnary computes op x.
func applyUnary(op string, x Value) (Value, error) {
	if _, err := unaryType(op, x.Type); err != nil {
		return Value{}, err
	}
	switch {
	case op == "!":
		return BoolVal(!x.Bool), nil
	case x.Type == TypeInt:
		return IntVal(-x.Int), nil
	case x.Type == TypeFloat:
		return FloatVal(-x.Float), nil
	}
	return DurVal(-x.Dur), nil
}

// errDivZero is the error of a division (or %) by zero.
var errDivZero = errors.New("division by zero")

// asFloat returns an int or a float as a float.
func asFloat(v Value) float64 {
	if v.Type == TypeInt {
		return float64(v.Int)
	}
	return v.Float
}

// applyBinary computes x op y (not && and ||, which EvalExpr short-circuits,
// except when both sides are given).
func applyBinary(op string, x, y Value) (Value, error) {
	t, err := binaryType(op, x.Type, y.Type)
	if err != nil {
		return Value{}, err
	}
	switch op {
	case "&&":
		return BoolVal(x.Bool && y.Bool), nil
	case "||":
		return BoolVal(x.Bool || y.Bool), nil
	case "==", "!=":
		_, eq := compare(x, y)
		return BoolVal(eq == (op == "==")), nil
	case "<", "<=", ">", ">=":
		c, _ := compare(x, y)
		return BoolVal(op == "<" && c < 0 || op == "<=" && c <= 0 || op == ">" && c > 0 || op == ">=" && c >= 0), nil
	}
	return arith(op, t, x, y)
}

// compare orders x and y (-1, 0, 1) and tells whether they are equal.
func compare(x, y Value) (int, bool) {
	switch {
	case isNum(x.Type) && isNum(y.Type) && (x.Type == TypeFloat || y.Type == TypeFloat):
		a, b := asFloat(x), asFloat(y)
		return cmpOrdered(a, b), a == b
	case x.Type == TypeInt:
		return cmpOrdered(x.Int, y.Int), x.Int == y.Int
	case x.Type == TypeString:
		return strings.Compare(x.Str, y.Str), x.Str == y.Str
	case x.Type == TypeDuration:
		return cmpOrdered(x.Dur, y.Dur), x.Dur == y.Dur
	}
	return 0, x.Bool == y.Bool
}

// cmpOrdered returns -1, 0 or 1 as a is less than, equal to or greater than b.
func cmpOrdered[T int64 | float64 | time.Duration](a, b T) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

// arith computes the arithmetic x op y whose result has type t.
func arith(op string, t Type, x, y Value) (Value, error) {
	switch t {
	case TypeString:
		return StrVal(x.String() + y.String()), nil
	case TypeInt:
		a, b := x.Int, y.Int
		switch op {
		case "+":
			return IntVal(a + b), nil
		case "-":
			return IntVal(a - b), nil
		case "*":
			return IntVal(a * b), nil
		}
		if b == 0 {
			return Value{}, errDivZero
		}
		if op == "/" {
			return IntVal(a / b), nil
		}
		return IntVal(a % b), nil
	case TypeFloat:
		if x.Type == TypeDuration { // duration / duration
			if y.Dur == 0 {
				return Value{}, errDivZero
			}
			return FloatVal(float64(x.Dur) / float64(y.Dur)), nil
		}
		a, b := asFloat(x), asFloat(y)
		switch op {
		case "+":
			return FloatVal(a + b), nil
		case "-":
			return FloatVal(a - b), nil
		case "*":
			return FloatVal(a * b), nil
		}
		if b == 0 {
			return Value{}, errDivZero
		}
		return FloatVal(a / b), nil
	}
	// A duration: ± another, or times / by a number.
	switch {
	case op == "+":
		return DurVal(x.Dur + y.Dur), nil
	case op == "-":
		return DurVal(x.Dur - y.Dur), nil
	case op == "*" && x.Type == TypeDuration:
		return DurVal(time.Duration(float64(x.Dur) * asFloat(y))), nil
	case op == "*":
		return DurVal(time.Duration(asFloat(x) * float64(y.Dur))), nil
	}
	if asFloat(y) == 0 {
		return Value{}, errDivZero
	}
	return DurVal(time.Duration(float64(x.Dur) / asFloat(y))), nil
}
