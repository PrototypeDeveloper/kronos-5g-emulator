package command

import (
	"fmt"
	"net/netip"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"kronos/common/diag"
)

// TokenKind classifies tokens. The same lexer serves scenario.txt and the
// CUI/GUI command line, so the two cannot drift apart.
type TokenKind uint8

var tokenNames = [...]string{
	TokEOF:      "end of input",
	TokIdent:    "name",
	TokBuiltin:  "builtin",
	TokInt:      "integer",
	TokFloat:    "float",
	TokDuration: "duration",
	TokString:   "string",
	TokLParen:   "(",
	TokRParen:   ")",
	TokLBrace:   "{",
	TokRBrace:   "}",
	TokComma:    ",",
	TokDot:      ".",
	TokColon:    ":",
	TokAssign:   "=",
	TokStar:     "*",
	TokIllegal:  "invalid token",
}

// String names the token kind for error messages.
func (k TokenKind) String() string { return tokenNames[k] }

// Describe names the token for error messages.
func (t Token) Describe() string {
	switch t.Kind {
	case TokIdent:
		return t.Text
	case TokBuiltin:
		return "$" + t.Text
	case TokString:
		return strconv.Quote(t.Text)
	case TokInt, TokFloat, TokDuration:
		return t.Text
	default:
		return t.Kind.String()
	}
}

// Found describes the token as found where something else was expected,
// e.g. "found send", or "the input ends here" at the end of the input.
func (t Token) Found() string {
	if t.Kind == TokEOF {
		return "the input ends here"
	}
	return "found " + t.Describe()
}

// NewLexer returns a lexer for src; file names the source in positions.
func NewLexer(file, src string) *Lexer {
	return &Lexer{src: src, file: file, line: 1, col: 1}
}

// Next consumes and returns the next token.
func (l *Lexer) Next() Token {
	if len(l.buf) > 0 {
		t := l.buf[0]
		l.buf = l.buf[1:]
		return t
	}
	return l.scan()
}

// Peek returns the next token without consuming it.
func (l *Lexer) Peek() Token { return l.PeekN(0) }

// PeekN returns the token n positions ahead (0 is the next token).
func (l *Lexer) PeekN(n int) Token {
	for len(l.buf) <= n {
		l.buf = append(l.buf, l.scan())
	}
	return l.buf[n]
}

// pos returns the current position.
func (l *Lexer) pos() diag.Pos { return diag.Pos{File: l.file, Line: l.line, Col: l.col} }

// peekRune returns the next rune without consuming it, or -1 at the end.
func (l *Lexer) peekRune() rune {
	if l.off >= len(l.src) {
		return -1
	}
	r, _ := utf8.DecodeRuneInString(l.src[l.off:])
	return r
}

// advance consumes the next rune and updates the line and column (-1 at the end).
func (l *Lexer) advance() rune {
	if l.off >= len(l.src) {
		return -1
	}
	r, n := utf8.DecodeRuneInString(l.src[l.off:])
	l.off += n
	if r == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return r
}

// skipSpaceAndComments skips whitespace, a byte order mark and comments.
func (l *Lexer) skipSpaceAndComments() {
	for {
		switch r := l.peekRune(); {
		case r == ' ' || r == '\t' || r == '\r' || r == '\n' || r == '\ufeff':
			l.advance()
		case r == '#' || (r == '/' && strings.HasPrefix(l.src[l.off:], "//")):
			for r := l.peekRune(); r != -1 && r != '\n'; r = l.peekRune() {
				l.advance()
			}
		default:
			return
		}
	}
}

// punct maps the one-character tokens to their kinds.
var punct = map[rune]TokenKind{
	'(': TokLParen, ')': TokRParen, '{': TokLBrace, '}': TokRBrace,
	',': TokComma, '.': TokDot, ':': TokColon, '=': TokAssign, '*': TokStar,
}

// isIdentStart reports whether r can start a name (an ASCII letter or _).
func isIdentStart(r rune) bool {
	return r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// isDigit reports whether r is an ASCII digit.
func isDigit(r rune) bool { return r >= '0' && r <= '9' }

// isIdentPart reports whether r can continue a name.
func isIdentPart(r rune) bool { return isIdentStart(r) || isDigit(r) }

// scan reads the next token from the source.
func (l *Lexer) scan() Token {
	l.skipSpaceAndComments()
	p := l.pos()
	r := l.peekRune()
	switch {
	case r == -1:
		return Token{Kind: TokEOF, Pos: p}
	case isIdentStart(r):
		start := l.off
		for isIdentPart(l.peekRune()) {
			l.advance()
		}
		return Token{Kind: TokIdent, Text: l.src[start:l.off], Pos: p}
	case r == '$':
		l.advance()
		start := l.off
		for isIdentPart(l.peekRune()) {
			l.advance()
		}
		if start == l.off {
			return Token{Kind: TokIllegal, Text: "$ must be followed by the name of a builtin", Pos: p}
		}
		return Token{Kind: TokBuiltin, Text: l.src[start:l.off], Pos: p}
	case isDigit(r) || (r == '-' && l.off+1 < len(l.src) && isDigit(rune(l.src[l.off+1]))):
		return l.scanNumber(p)
	case r == '"':
		return l.scanString(p)
	}
	l.advance()
	if k, ok := punct[r]; ok {
		return Token{Kind: k, Text: string(r), Pos: p}
	}
	return Token{Kind: TokIllegal, Text: fmt.Sprintf("invalid character %q", r), Pos: p}
}

// scanNumber reads a word that starts with a digit (or "-digit") and decides
// whether it is an int, a float, a duration or a name such as a signal label.
func (l *Lexer) scanNumber(p diag.Pos) Token {
	start := l.off
	if l.peekRune() == '-' {
		l.advance()
	}
	for r := l.peekRune(); isIdentPart(r) || r == '.' || r == 'µ'; r = l.peekRune() {
		l.advance()
	}
	word := l.src[start:l.off]
	digits := strings.TrimPrefix(word, "-")
	switch {
	case allDigits(digits):
		v, err := strconv.ParseInt(word, 10, 64)
		if err != nil {
			// Too large for an int but still a valid IMSI-like digit string.
			return Token{Kind: TokInt, Text: word, Pos: p}
		}
		return Token{Kind: TokInt, Text: word, Val: IntVal(v), Pos: p}
	case isFloat(digits):
		v, err := strconv.ParseFloat(word, 64)
		if err != nil {
			return Token{Kind: TokIllegal, Text: fmt.Sprintf("cannot read the number %s", word), Pos: p}
		}
		return Token{Kind: TokFloat, Text: word, Val: FloatVal(v), Pos: p}
	}
	if d, err := time.ParseDuration(word); err == nil {
		return Token{Kind: TokDuration, Text: word, Val: DurVal(d), Pos: p}
	}
	if a, err := netip.ParseAddr(word); err == nil && a.Is4() {
		// ping(10.10.10.20): an IPv4 address needs no quotes
		return Token{Kind: TokString, Text: word, Val: StrVal(word), Pos: p}
	}
	if word[0] != '-' && !strings.ContainsAny(word, ".µ") {
		// A name that starts with a digit, e.g. a signal label.
		return Token{Kind: TokIdent, Text: word, Pos: p}
	}
	return Token{Kind: TokIllegal, Text: fmt.Sprintf("cannot read %s as a number or a duration (duration units: ns, us, ms, s, m, h)", word), Pos: p}
}

// allDigits reports whether s is non-empty and all ASCII digits.
func allDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) {
			return false
		}
	}
	return true
}

// isFloat reports whether s is digits, a dot and digits (e.g. 1.5).
func isFloat(s string) bool {
	i := strings.IndexByte(s, '.')
	return i > 0 && allDigits(s[:i]) && allDigits(s[i+1:])
}

// scanString reads a double-quoted string on one line and resolves its escapes.
func (l *Lexer) scanString(p diag.Pos) Token {
	l.advance() // opening quote
	var b strings.Builder
	for {
		r := l.advance()
		switch r {
		case -1, '\n':
			return Token{Kind: TokIllegal, Text: "unterminated string (a closing \" is missing)", Pos: p}
		case '"':
			return Token{Kind: TokString, Text: b.String(), Val: StrVal(b.String()), Pos: p}
		case '\\':
			e := l.advance()
			switch e {
			case 'n':
				b.WriteByte('\n')
			case 't':
				b.WriteByte('\t')
			case 'r':
				b.WriteByte('\r')
			case '"':
				b.WriteByte('"')
			case '\\':
				b.WriteByte('\\')
			default:
				return Token{Kind: TokIllegal, Text: fmt.Sprintf("invalid escape \\%c (use \\n \\t \\r \\\" \\\\)", e), Pos: p}
			}
		default:
			b.WriteRune(r)
		}
	}
}
