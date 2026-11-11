package scenario

import (
	"os"

	"kronos/command"
	"kronos/common/diag"
)

// BlockKind distinguishes the blocks of scenario.txt.
type BlockKind uint8

// State returns the state block named name.
func (f *File) State(name string) (*Block, bool) {
	b, ok := f.states[name]
	return b, ok
}

// Blocks returns node, any and the state blocks, skipping absent ones.
func (f *File) Blocks() []*Block {
	var out []*Block
	if f.Node != nil {
		out = append(out, f.Node)
	}
	if f.Any != nil {
		out = append(out, f.Any)
	}
	return append(out, f.States...)
}

// Scope is the subject the block's commands run on.
func (b *Block) Scope() command.Scope {
	if b.Kind == BlockNode {
		return command.ScopeNode
	}
	return command.ScopeUE
}

// String returns the trigger as written ("*" for the wildcard).
func (t Trigger) String() string {
	switch {
	case t.Wildcard:
		return "*"
	case t.NAS != "":
		return t.NGAP + ":" + t.NAS
	default:
		return t.NGAP
	}
}

// StripCalls removes every call to the named command and returns how many
// were removed. Handlers stay in place even when they become empty, so the
// message still counts as handled. Used for snap when base.snap is false.
func (f *File) StripCalls(name string) int {
	n := 0
	strip := func(calls []command.Call) []command.Call {
		out := calls[:0]
		for _, c := range calls {
			if c.Name == name {
				n++
				continue
			}
			out = append(out, c)
		}
		return out
	}
	for _, b := range f.Blocks() {
		if b.Entry != nil {
			b.Entry.Calls = strip(b.Entry.Calls)
		}
		for _, h := range b.Handlers {
			h.Calls = strip(h.Calls)
		}
	}
	return n
}

// Calls visits every call in the file.
func (f *File) Calls(fn func(b *Block, c command.Call)) {
	for _, b := range f.Blocks() {
		if b.Entry != nil {
			for _, c := range b.Entry.Calls {
				fn(b, c)
			}
		}
		for _, h := range b.Handlers {
			for _, c := range h.Calls {
				fn(b, c)
			}
		}
	}
}

// reserved words cannot name a state.
var reserved = map[string]bool{
	"var": true, "public": true, "private": true, "node": true, "any": true,
	"receive": true, "entry": true, "timeout": true, command.TargetKeyword: true,
}

// ReadScenario reads and parses a scenario file.
func ReadScenario(path string) (*File, diag.List) {
	b, err := os.ReadFile(path)
	if err != nil {
		var ds diag.List
		ds.Errorf(diag.Pos{File: path}, "cannot read the scenario file: %v", err)
		return nil, ds
	}
	return ParseScenario(path, b)
}

// ParseScenario parses scenario.txt. It reports syntax errors
// only; Validate checks names, scopes and commands.
func ParseScenario(file string, src []byte) (*File, diag.List) {
	p := &parser{lx: command.NewLexer(file, string(src)), f: &File{Path: file, states: map[string]*Block{}}}
	for p.lx.Peek().Kind != command.TokEOF {
		p.block()
	}
	return p.f, p.ds
}

// errorf reports a syntax error at t; an illegal token reports its own message.
func (p *parser) errorf(t command.Token, format string, args ...any) {
	if t.Kind == command.TokIllegal {
		p.ds.Errorf(t.Pos, "%s", t.Text)
		return
	}
	p.ds.Errorf(t.Pos, format, args...)
}

// block parses one top-level block: var, node, any or a state.
func (p *parser) block() {
	t := p.lx.Next()
	if t.Kind != command.TokIdent {
		p.errorf(t, "a block is needed: var, node, any or a state name (%s)", t.Found())
		p.recoverTopLevel()
		return
	}
	if open := p.lx.Peek(); open.Kind != command.TokLBrace {
		p.errorf(open, "%s must be followed by {", t.Text)
		p.recoverTopLevel()
		return
	}
	p.lx.Next()
	switch t.Text {
	case "var":
		p.varBody()
	case "node", "any":
		kind, slot := BlockNode, &p.f.Node
		if t.Text == "any" {
			kind, slot = BlockAny, &p.f.Any
		}
		b := &Block{Kind: kind, Name: t.Text, Pos: t.Pos}
		if *slot != nil {
			p.ds.Errorf(t.Pos, "a second %s block (the first is on line %d)", t.Text, (*slot).Pos.Line)
		} else {
			*slot = b
		}
		p.items(b)
	default:
		b := &Block{Kind: BlockState, Name: t.Text, Pos: t.Pos}
		if reserved[t.Text] {
			p.ds.Errorf(t.Pos, "the reserved word %s cannot name a state", t.Text)
		} else if prev, dup := p.f.states[t.Text]; dup {
			p.ds.Errorf(t.Pos, "a second state %s (the first is on line %d)", t.Text, prev.Pos.Line)
		} else {
			p.f.states[t.Text] = b
			p.f.States = append(p.f.States, b)
		}
		p.items(b)
	}
}

// recoverTopLevel skips to the start of the next block after a syntax error.
func (p *parser) recoverTopLevel() {
	depth := 0
	for {
		t := p.lx.Peek()
		switch t.Kind {
		case command.TokEOF:
			return
		case command.TokLBrace:
			depth++
		case command.TokRBrace:
			p.lx.Next()
			if depth <= 1 {
				return
			}
			depth--
			continue
		case command.TokIdent:
			if depth == 0 && p.lx.PeekN(1).Kind == command.TokLBrace {
				return
			}
		}
		p.lx.Next()
	}
}

// varBody parses the var block: lines of "public|private name = literal".
func (p *parser) varBody() {
	for {
		// A declaration starts with public or private; } ends the block.
		t := p.lx.Next()
		switch {
		case t.Kind == command.TokRBrace:
			return
		case t.Kind == command.TokEOF:
			p.errorf(t, "the var block is not closed (} is missing)")
			return
		case t.Kind == command.TokIdent && (t.Text == "public" || t.Text == "private"):
		default:
			p.errorf(t, "a variable declaration starts with public or private (%s)", t.Found())
			continue
		}
		v := &Var{Pos: t.Pos, Public: t.Text == "public"}
		name := p.lx.Next()
		if name.Kind != command.TokIdent {
			p.errorf(name, "%s must be followed by a variable name (%s)", t.Text, name.Found())
			if name.Kind == command.TokRBrace {
				return
			}
			continue
		}
		v.Name, v.Pos = name.Text, name.Pos
		// The name is followed by = and a literal initial value.
		if eq := p.lx.Next(); eq.Kind != command.TokAssign {
			p.errorf(eq, "%s must be followed by = and an initial value", v.Name)
			if eq.Kind == command.TokRBrace {
				return
			}
			continue
		}
		lit := p.lx.Next()
		switch lit.Kind {
		case command.TokInt, command.TokFloat, command.TokString, command.TokDuration:
			if lit.Val.Type == command.TypeNone {
				p.errorf(lit, "the initial value %[2]s of %[1]s is too large", v.Name, lit.Text)
				continue
			}
			v.Init = lit.Val
		default:
			p.errorf(lit, "the initial value of %s must be a literal: 1, 1.1, \"text\", 5s … (%s)", v.Name, lit.Found())
			if lit.Kind == command.TokRBrace {
				return
			}
			continue
		}
		p.f.Vars = append(p.f.Vars, v)
	}
}

// items parses the receive and entry items of a block up to its closing }.
func (p *parser) items(b *Block) {
	for {
		t := p.lx.Peek()
		switch {
		case t.Kind == command.TokRBrace:
			p.lx.Next()
			return
		case t.Kind == command.TokEOF:
			p.ds.Errorf(b.Pos, "the %s block is not closed (} is missing)", b.Name)
			return
		case t.Kind == command.TokIdent && t.Text == "receive":
			p.handler(b)
		case t.Kind == command.TokIdent && t.Text == "entry":
			p.entry(b)
		case t.Kind == command.TokIdent && t.Text == "timeout":
			p.lx.Next()
			p.ds.Errorf(t.Pos, "timeout is not supported")
			p.recoverItem()
		case t.Kind == command.TokIdent && p.lx.PeekN(1).Kind == command.TokLParen:
			p.errorf(t, "the command %s cannot be written directly in a block: write it in receive(…) { … } or entry { … }", t.Text)
			if _, err := command.ParseCall(p.lx); err != nil {
				p.recoverItem()
			}
		default:
			p.errorf(t, "receive(…), entry or timeout(…) is needed (%s)", t.Found())
			p.lx.Next()
			p.recoverItem()
		}
	}
}

// recoverItem skips the rest of a broken item: up to and including its
// { ... } body, or up to the next item keyword or the end of the block.
func (p *parser) recoverItem() {
	for {
		t := p.lx.Peek()
		switch {
		case t.Kind == command.TokEOF || t.Kind == command.TokRBrace:
			return
		case t.Kind == command.TokLBrace:
			p.skipGroup()
			return
		case t.Kind == command.TokIdent && (t.Text == "receive" || t.Text == "entry" || t.Text == "timeout"):
			return
		}
		p.lx.Next()
	}
}

// skipGroup skips a balanced { ... } starting at the current '{'.
func (p *parser) skipGroup() {
	depth := 0
	for {
		switch p.lx.Next().Kind {
		case command.TokEOF:
			return
		case command.TokLBrace:
			depth++
		case command.TokRBrace:
			depth--
			if depth <= 0 {
				return
			}
		}
	}
}

// handler parses receive(NGAP[:NAS]) { calls } or receive(*) { calls } into b.
func (p *parser) handler(b *Block) {
	kw := p.lx.Next()
	h := &Handler{Pos: kw.Pos}
	if t := p.lx.Next(); t.Kind != command.TokLParen {
		p.errorf(t, "receive must be followed by (")
		p.recoverItem()
		return
	}
	switch t := p.lx.Next(); t.Kind {
	case command.TokStar:
		h.Trigger = Trigger{Pos: t.Pos, Wildcard: true}
	case command.TokIdent:
		h.Trigger = Trigger{Pos: t.Pos, NGAP: t.Text}
		if p.lx.Peek().Kind == command.TokColon {
			p.lx.Next()
			n := p.lx.Next()
			if n.Kind != command.TokIdent {
				p.errorf(n, ": must be followed by a NAS message name (%s)", n.Found())
				p.recoverItem()
				return
			}
			h.Trigger.NAS = n.Text
		}
	default:
		p.errorf(t, "receive() takes a message name: NGAP or NGAP:NAS (%s)", t.Found())
		p.recoverItem()
		return
	}
	if t := p.lx.Next(); t.Kind != command.TokRParen {
		p.errorf(t, "receive(%s needs a closing ) (%s)", h.Trigger, t.Found())
		p.recoverItem()
		return
	}
	if t := p.lx.Next(); t.Kind != command.TokLBrace {
		p.errorf(t, "receive(%s) must be followed by {", h.Trigger)
		p.recoverItem()
		return
	}
	h.Calls = p.calls()
	b.Handlers = append(b.Handlers, h)
}

// entry parses entry { calls }; any blocks take none and others at most one.
func (p *parser) entry(b *Block) {
	kw := p.lx.Next()
	if t := p.lx.Next(); t.Kind != command.TokLBrace {
		p.errorf(t, "entry must be followed by {")
		p.recoverItem()
		return
	}
	e := &Entry{Pos: kw.Pos, Calls: p.calls()}
	switch {
	case b.Kind == BlockAny:
		p.ds.Errorf(kw.Pos, "any cannot have an entry (any is not a state)")
	case b.Entry != nil:
		p.ds.Errorf(kw.Pos, "only one entry per block (the first is on line %d)", b.Entry.Pos.Line)
	default:
		b.Entry = e
	}
}

// calls parses commands up to the closing '}' of a handler or entry body.
func (p *parser) calls() []command.Call {
	var out []command.Call
	for {
		switch t := p.lx.Peek(); t.Kind {
		case command.TokRBrace:
			p.lx.Next()
			return out
		case command.TokEOF:
			p.errorf(t, "} is missing")
			return out
		}
		c, err := command.ParseCall(p.lx)
		if err != nil {
			p.ds.AddErr(err)
			p.skipToClose()
			return out
		}
		if d := p.lx.Peek(); d.Kind == command.TokDot {
			p.lx.Next()
			if c.Name == command.TargetKeyword {
				p.ds.Errorf(c.Pos, "target() cannot be used in a scenario (commands run on the subscriber whose handler runs)")
			} else {
				p.ds.Errorf(d.Pos, "commands cannot be joined with .")
			}
			if p.lx.Peek().Kind == command.TokIdent {
				if _, err := command.ParseCall(p.lx); err != nil {
					p.skipToClose()
					return out
				}
			}
			continue
		}
		out = append(out, c)
	}
}

// skipToClose skips to and past the '}' that closes the current body.
func (p *parser) skipToClose() {
	depth := 0
	for {
		switch p.lx.Next().Kind {
		case command.TokEOF:
			return
		case command.TokLBrace:
			depth++
		case command.TokRBrace:
			if depth == 0 {
				return
			}
			depth--
		}
	}
}
