package scenario

import (
	"errors"
	"strings"

	"kronos/command"
)

// bodies returns the bodies of the block: its entry, receive handlers and expires.
func (b *Block) bodies() []*[]Stmt {
	var out []*[]Stmt
	if b.Entry != nil {
		out = append(out, &b.Entry.Body)
	}
	for _, h := range b.Handlers {
		out = append(out, &h.Body)
	}
	for _, x := range b.Expires {
		out = append(out, &x.Body)
	}
	for _, x := range b.Syncs {
		out = append(out, &x.Body)
	}
	return out
}

// eachCall visits every call in body, also those inside if.
func eachCall(body []Stmt, fn func(c command.Call)) {
	for _, st := range body {
		if st.If == nil {
			fn(st.Call)
			continue
		}
		eachCall(st.If.Then, fn)
		eachCall(st.If.Else, fn)
	}
}

// stripCalls removes the calls to name from body, also inside if, adding their number to n.
func stripCalls(body []Stmt, name string, n *int) []Stmt {
	out := body[:0]
	for _, st := range body {
		switch {
		case st.If != nil:
			st.If.Then = stripCalls(st.If.Then, name, n)
			st.If.Else = stripCalls(st.If.Else, name, n)
		case st.Call.Name == name:
			*n++
			continue
		}
		out = append(out, st)
	}
	return out
}

// String returns the statement on one line, as gnodeb check -v shows it.
func (st Stmt) String() string {
	if st.If == nil {
		return st.Call.String()
	}
	x := st.If
	s := "if (" + x.Cond.String() + ") {" + inline(x.Then) + "}"
	switch {
	case x.ElseIf:
		s += " else " + x.Else[0].String()
	case len(x.Else) > 0:
		s += " else {" + inline(x.Else) + "}"
	}
	return s
}

// inline returns the statements of a body between spaces, joined with ", ".
func inline(body []Stmt) string {
	parts := make([]string, len(body))
	for i, st := range body {
		parts[i] = st.String()
	}
	return " " + strings.Join(parts, ", ") + " "
}

// StmtLines returns a body as lines, the bodies of if indented.
func StmtLines(body []Stmt) []string {
	var out []string
	appendLines(&out, body, "")
	return out
}

// appendLines adds the lines of body at the indent.
func appendLines(out *[]string, body []Stmt, indent string) {
	for _, st := range body {
		if st.If == nil {
			*out = append(*out, indent+st.Call.String())
			continue
		}
		head := indent + "if ("
		for x := st.If; ; x = x.Else[0].If {
			*out = append(*out, head+x.Cond.String()+") {")
			appendLines(out, x.Then, indent+bodyIndent)
			if !x.ElseIf {
				if len(x.Else) > 0 {
					*out = append(*out, indent+"} else {")
					appendLines(out, x.Else, indent+bodyIndent)
				}
				break
			}
			head = indent + "} else if ("
		}
		*out = append(*out, indent+"}")
	}
}

// Run executes body on env: the calls in order, and for each if the body
// of the first condition that holds. fail reports a call or a condition
// that failed (a condition that fails runs neither body); the rest of the
// body runs anyway, except after a control() wait cut short by the link
// going down, which ends the whole body quietly.
func Run(env command.Env, body []Stmt, fail func(what string, err error)) {
	run(env, body, fail)
}

// run is Run; it returns false when the body was cut short.
func run(env command.Env, body []Stmt, fail func(what string, err error)) bool {
	for _, st := range body {
		if st.If == nil {
			err := command.Exec(env, st.Call)
			switch {
			case errors.Is(err, command.ErrInterrupted):
				return false
			case err != nil:
				fail(st.Call.String(), err)
			}
			continue
		}
		ok, err := command.EvalCond(env, st.If.Cond)
		switch {
		case err != nil:
			fail("if ("+st.If.Cond.String()+")", err)
		case ok:
			if !run(env, st.If.Then, fail) {
				return false
			}
		default:
			if !run(env, st.If.Else, fail) {
				return false
			}
		}
	}
	return true
}
