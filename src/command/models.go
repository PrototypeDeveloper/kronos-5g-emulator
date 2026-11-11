// Data types of package command.

package command

import (
	"time"

	"kronos/common/diag"
)

// Command definitions.

// Spec defines a command.
type Spec struct {
	Name  string
	Args  []Arg
	Scope Scope
	Where Where
	Help  string
	// Check optionally adds command-specific static checks (e.g. the snap format).
	Check func(c Call, r Resolver, scope Scope) diag.List
	// Run executes the command for one subject. args are already evaluated.
	Run func(env Env, args []Value) error
	// RunCall, if set, runs instead of Run with the arguments as written (set).
	RunCall func(env Env, c Call) error
}

// Arg describes one parameter of a command.
type Arg struct {
	Name     string
	Kind     Kind
	Optional bool
	Variadic bool     // the last parameter may repeat (snap values)
	Choices  []string // KindChoice: the words allowed
	ZeroRate bool     // KindRate: 0 is allowed (no limit)
	// Default is what an omitted optional argument means, as written on the
	// command line. The GUI writes it when a later argument is given.
	Default string
}

// Values, variables and builtins.

// Value is an int, float, string or duration (or a bool, the result of a comparison).
type Value struct {
	Type  Type
	Int   int64
	Float float64
	Str   string
	Dur   time.Duration
	Bool  bool
}

// VarInfo describes a scenario variable.
type VarInfo struct {
	Type   Type
	Public bool
}

// BuiltinInfo describes a builtin value such as $imsi.
type BuiltinInfo struct {
	Type  Type
	Scope Scope // subjects that have the value
	Help  string
}

// The lexer.

// Token is one lexical token.
type Token struct {
	Kind TokenKind
	Text string
	Val  Value // value of a literal token
	Pos  diag.Pos
}

// Lexer splits source text into tokens. Whitespace and newlines are not
// significant; `#` and `//` start a comment that runs to the end of the line.
type Lexer struct {
	src       string
	file      string
	off       int
	line, col int
	buf       []Token // lookahead
}

// Parsed calls and command lines.

// Expr is one argument of a call, or a part of one.
type Expr struct {
	Kind ExprKind
	Name string     // ExprIdent and ExprBuiltin; the variable of ExprAssign
	Val  Value      // ExprLiteral
	Op   string     // ExprUnary and ExprBinary: "-", "!", "+", "==", "&&" …
	X, Y *Expr      // the operands (X of ExprParen and ExprUnary; Y, the value, of ExprAssign)
	Path []FieldSeg // ExprField: the IE, from the outside in
	Pos  diag.Pos
}

// FieldSeg is one step of the path of an IE: a name, and an index into a list.
type FieldSeg struct {
	Name     string
	Index    int
	HasIndex bool
}

// Call is `name(arg, ...)`, the form shared by scenario actions and the command line.
type Call struct {
	Name string
	Args []Expr
	Pos  diag.Pos
}

// Line is one CUI/GUI command line: `[target(IMSI).]name(args)`.
// An empty Target means every subscriber.
type Line struct {
	Target    string
	TargetPos diag.Pos
	Call      Call
}

// Validation.

// Context says where a call is checked.
type Context struct {
	Where Where // InScenario or Interactive
	Scope Scope // InScenario only: ScopeNode for the node block, ScopeUE for any and state blocks
}

// FormatVerb is one conversion (%d, %s ...) in a snap format.
type FormatVerb struct {
	Verb rune
	Off  int // byte offset of the '%'
}

// Arguments handed to the environment.
