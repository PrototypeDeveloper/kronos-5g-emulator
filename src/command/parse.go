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
		if a, err := netip.ParseAddrPort(e.Val.Str); e.Val.Type == TypeString && err == nil && a.Addr().Is4() {
			return e.Val.Str
		}
		return e.Val.Literal()
	case ExprParen:
		return "(" + e.X.String() + ")"
	case ExprUnary:
		return e.Op + e.X.String()
	case ExprBinary:
		return e.X.String() + " " + e.Op + " " + e.Y.String()
	case ExprAssign:
		return e.Name + " = " + e.Y.String()
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
		e, err := parseArg(lx)
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

// parseArg parses one argument: an expression, or `name = expression` (for set).
func parseArg(lx *Lexer) (Expr, error) {
	if t := lx.Peek(); t.Kind == TokIdent && lx.PeekN(1).Kind == TokAssign {
		lx.Next()
		lx.Next()
		y, err := ParseExpr(lx)
		if err != nil {
			return Expr{}, err
		}
		return Expr{Kind: ExprAssign, Name: t.Text, Y: &y, Pos: t.Pos}, nil
	}
	e, err := ParseExpr(lx)
	if err != nil {
		return Expr{}, err
	}
	if t := lx.Peek(); t.Kind == TokAssign {
		return Expr{}, errAt(t, "only a variable can be assigned (%s is not one); compare with ==", e)
	}
	return e, nil
}

// ParseExpr parses an expression: literals, variables and builtins joined
// by operators as in C (|| && == != < <= > >= + - * / % and unary - !), with ( ).
func ParseExpr(lx *Lexer) (Expr, error) { return parseBinary(lx, precOr) }

// binaryPrec returns the precedence of a binary operator token (precNone if it is not one).
func binaryPrec(k TokenKind) int {
	switch k {
	case TokOr:
		return precOr
	case TokAnd:
		return precAnd
	case TokEq, TokNe:
		return precEq
	case TokLt, TokLe, TokGt, TokGe:
		return precCmp
	case TokPlus, TokMinus:
		return precAdd
	case TokStar, TokSlash, TokPercent:
		return precMul
	}
	return precNone
}

// parseBinary parses operands joined by binary operators of precedence min
// or higher, left to right.
func parseBinary(lx *Lexer, min int) (Expr, error) {
	x, err := parseUnary(lx)
	if err != nil {
		return Expr{}, err
	}
	for {
		op := lx.Peek()
		prec := binaryPrec(op.Kind)
		if prec == precNone || prec < min {
			return x, nil
		}
		lx.Next()
		y, err := parseBinary(lx, prec+1)
		if err != nil {
			return Expr{}, err
		}
		l, r := x, y
		x = Expr{Kind: ExprBinary, Op: op.Text, X: &l, Y: &r, Pos: op.Pos}
	}
}

// parseUnary parses - or ! before an operand. A minus before a number or a
// duration makes a negative literal (-1, -5s).
func parseUnary(lx *Lexer) (Expr, error) {
	t := lx.Peek()
	if t.Kind != TokMinus && t.Kind != TokNot {
		return parsePrimary(lx)
	}
	lx.Next()
	x, err := parseUnary(lx)
	if err != nil {
		return Expr{}, err
	}
	if t.Kind == TokMinus && x.Kind == ExprLiteral {
		switch x.Val.Type {
		case TypeInt:
			return Expr{Kind: ExprLiteral, Val: IntVal(-x.Val.Int), Pos: t.Pos}, nil
		case TypeFloat:
			return Expr{Kind: ExprLiteral, Val: FloatVal(-x.Val.Float), Pos: t.Pos}, nil
		case TypeDuration:
			return Expr{Kind: ExprLiteral, Val: DurVal(-x.Val.Dur), Pos: t.Pos}, nil
		}
	}
	return Expr{Kind: ExprUnary, Op: t.Text, X: &x, Pos: t.Pos}, nil
}

// parsePrimary parses a name, a builtin, a literal or a parenthesized expression.
func parsePrimary(lx *Lexer) (Expr, error) {
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
	case TokLParen:
		x, err := ParseExpr(lx)
		if err != nil {
			return Expr{}, err
		}
		if r := lx.Next(); r.Kind != TokRParen {
			return Expr{}, errAt(r, "( needs a closing ) (%s)", r.Found())
		}
		return Expr{Kind: ExprParen, X: &x, Pos: t.Pos}, nil
	}
	return Expr{}, errAt(t, "a value is needed (%s)", t.Found())
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
