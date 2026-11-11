package command

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"kronos/common/diag"
)

// ParseFormat returns the conversions of a Go fmt format that consume an
// argument. Explicit argument indexes ([n]) and '*' widths are rejected so
// the number of values can be checked statically.
func ParseFormat(f string) ([]FormatVerb, error) {
	var out []FormatVerb
	for i := 0; i < len(f); i++ {
		if f[i] != '%' {
			continue
		}
		start := i
		i++
		if i < len(f) && f[i] == '%' {
			continue
		}
		for i < len(f) && strings.IndexByte("+-# 0", f[i]) >= 0 {
			i++
		}
		for i < len(f) && isDigit(rune(f[i])) {
			i++
		}
		if i < len(f) && f[i] == '.' {
			i++
			for i < len(f) && isDigit(rune(f[i])) {
				i++
			}
		}
		if i >= len(f) {
			return nil, fmt.Errorf("%q has no conversion character (write %%%% to print %%)", f[start:])
		}
		switch c := f[i]; {
		case c == '*' || c == '[':
			return nil, fmt.Errorf("* and [n] are not supported")
		case strings.IndexByte(allowedVerbs, c) < 0:
			r, _ := utf8.DecodeRuneInString(f[i:])
			return nil, fmt.Errorf("%%%c is not supported (use %%d %%s %%v %%x %%f %%q …)", r)
		}
		out = append(out, FormatVerb{Verb: rune(f[i]), Off: start})
	}
	return out, nil
}

// verbAccepts reports whether fmt prints a value of type t sensibly with verb.
func verbAccepts(verb rune, t Type) bool {
	switch verb {
	case 'v', 'T', 'x', 'X':
		return true
	case 'd', 'o', 'O', 'b', 'c', 'U':
		return t == TypeInt || t == TypeDuration
	case 'e', 'E', 'f', 'F', 'g', 'G':
		return t == TypeFloat
	case 's', 'q':
		return t == TypeString || t == TypeDuration
	default: // 't' needs a bool, which scenarios do not have
		return false
	}
}

// Format renders a snap line.
func Format(format string, vals []Value) string {
	args := make([]any, len(vals))
	for i, v := range vals {
		args[i] = v.Interface()
	}
	return fmt.Sprintf(format, args...)
}

// exprType returns the static type of an argument, or TypeNone if unknown.
func exprType(e Expr, r Resolver) Type {
	switch e.Kind {
	case ExprLiteral:
		return e.Val.Type
	case ExprBuiltin:
		if b, ok := LookupBuiltin(e.Name); ok {
			return b.Type
		}
	case ExprIdent:
		if v, ok := r.Var(e.Name); ok {
			return v.Type
		}
	}
	return TypeNone
}

// checkSnap verifies the format and that it consumes exactly the given values.
// It runs even when base.snap is false, so enabling snap later cannot fail.
func checkSnap(c Call, r Resolver, _ Scope) diag.List {
	var ds diag.List
	f := c.Args[0]
	verbs, err := ParseFormat(f.Val.Str)
	if err != nil {
		ds.Errorf(f.Pos, "snap format: %v", err)
		return ds
	}
	vals := c.Args[1:]
	if len(verbs) != len(vals) {
		ds.Errorf(c.Pos, "snap: %d conversion(s) in the format but %d value(s)", len(verbs), len(vals))
		return ds
	}
	for i, v := range verbs {
		if t := exprType(vals[i], r); t != TypeNone && !verbAccepts(v.Verb, t) {
			ds.Warnf(vals[i].Pos, "%%%c gets the %s value %s (%%v takes any type)", v.Verb, t, vals[i])
		}
	}
	return ds
}
