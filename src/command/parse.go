package command

import (
	"errors"
	"net/netip"
	"strings"

	"kronos/common/diag"
)

// ExprKind classifies a command argument as written.
type ExprKind uint8

// String returns the argument as it is written.
func (e Expr) String() string {
	switch e.Kind {
	case ExprBuiltin:
		return "$" + e.Name
	case ExprLiteral:
		if a, err := netip.ParseAddr(e.Val.Str); e.Val.Type == TypeString && err == nil && a.Is4() {
			return e.Val.Str // written without quotes (see the lexer)
		}
		return e.Val.Literal()
	default:
		return e.Name
	}
}

// String returns the call as name(args).
func (c Call) String() string {
	args := make([]string, len(c.Args))
	for i, a := range c.Args {
		args[i] = a.String()
	}
	return c.Name + "(" + strings.Join(args, ", ") + ")"
}

// String returns the line as written, with its target(IMSI). prefix if any.
func (l Line) String() string {
	if l.Target == "" {
		return l.Call.String()
	}
	return "target(" + l.Target + ")." + l.Call.String()
}

// ErrEmptyLine is returned by ParseLine for a blank or comment-only line.
var ErrEmptyLine = errors.New("empty command line")

// errAt returns an error at t; an illegal token reports its own message instead.
func errAt(t Token, format string, args ...any) error {
	if t.Kind == TokIllegal {
		return diag.Errorf(t.Pos, "%s", t.Text)
	}
	return diag.Errorf(t.Pos, format, args...)
}

// ParseCall parses `IDENT "(" [ expr { "," expr } ] ")"` at the lexer's position.
// The scenario parser calls it for every action inside a handler.
func ParseCall(lx *Lexer) (Call, error) {
	t := lx.Next()
	if t.Kind != TokIdent {
		return Call{}, errAt(t, "a command is needed, e.g. send(signal) (%s)", t.Found())
	}
	c := Call{Name: t.Text, Pos: t.Pos}
	if p := lx.Next(); p.Kind != TokLParen {
		return c, errAt(p, "%s must be followed by ( (a command is written name(arguments, …))", c.Name)
	}
	if lx.Peek().Kind == TokRParen {
		lx.Next()
		return c, nil
	}
	for {
		e, err := parseExpr(lx)
		if err != nil {
			return c, err
		}
		c.Args = append(c.Args, e)
		switch t := lx.Next(); t.Kind {
		case TokComma:
		case TokRParen:
			return c, nil
		default:
			return c, errAt(t, "an argument must be followed by , or ) (%s)", t.Found())
		}
	}
}

// parseExpr parses one argument: a name, a builtin or a literal.
func parseExpr(lx *Lexer) (Expr, error) {
	t := lx.Next()
	switch t.Kind {
	case TokIdent:
		return Expr{Kind: ExprIdent, Name: t.Text, Pos: t.Pos}, nil
	case TokBuiltin:
		return Expr{Kind: ExprBuiltin, Name: t.Text, Pos: t.Pos}, nil
	case TokInt, TokFloat, TokDuration, TokString:
		if t.Val.Type == TypeNone {
			return Expr{}, errAt(t, "integer %s is too large", t.Text)
		}
		return Expr{Kind: ExprLiteral, Val: t.Val, Pos: t.Pos}, nil
	}
	return Expr{}, errAt(t, "an argument is needed (%s)", t.Found())
}

// ParseLine parses one CUI/GUI command line. It returns ErrEmptyLine for a
// blank line. Validation against the registry is done by CheckLine.
func ParseLine(src string) (Line, error) {
	lx := NewLexer("", src)
	if lx.Peek().Kind == TokEOF {
		return Line{}, ErrEmptyLine
	}
	var ln Line
	// The optional target(IMSI). prefix.
	if t := lx.Peek(); t.Kind == TokIdent && t.Text == TargetKeyword {
		lx.Next()
		if p := lx.Next(); p.Kind != TokLParen {
			return ln, errAt(p, "target must be followed by ( (e.g. target(208930000000001).send(signal))")
		}
		a := lx.Next()
		switch a.Kind {
		case TokInt, TokString:
			ln.Target, ln.TargetPos = a.Text, a.Pos
		default:
			return ln, errAt(a, "target() takes an IMSI (%s)", a.Found())
		}
		switch r := lx.Next(); r.Kind {
		case TokRParen:
		case TokComma:
			return ln, errAt(r, "target() takes only one IMSI")
		default:
			return ln, errAt(r, "target(IMSI) needs a closing ) (%s)", r.Found())
		}
		if d := lx.Next(); d.Kind != TokDot {
			return ln, errAt(d, "target(…) must be followed by . and a command (e.g. target(%s).send(signal))", ln.Target)
		}
	}
	// A command without arguments may omit the parentheses on the command line: linkup
	if t := lx.Peek(); t.Kind == TokIdent && lx.PeekN(1).Kind == TokEOF {
		lx.Next()
		if t.Text == TargetKeyword {
			return ln, diag.Errorf(t.Pos, "target goes before a command (e.g. target(208930000000001).send(signal))")
		}
		ln.Call = Call{Name: t.Text, Pos: t.Pos}
		return ln, nil
	}
	// The command itself; nothing may follow it.
	c, err := ParseCall(lx)
	if err != nil {
		return ln, err
	}
	if c.Name == TargetKeyword {
		return ln, diag.Errorf(c.Pos, "target goes before a command (e.g. target(208930000000001).send(signal))")
	}
	ln.Call = c
	if t := lx.Next(); t.Kind != TokEOF {
		if t.Kind == TokDot {
			return ln, errAt(t, "only one command per line (only target(…) can be joined with .)")
		}
		return ln, errAt(t, "unexpected text after the command (%s)", t.Describe())
	}
	return ln, nil
}
