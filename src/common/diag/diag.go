// Package diag holds validation results (errors, warnings and notes) together
// with the source position they refer to. The config, signal and scenario
// loaders all report through it, so `kronos check` can print them uniformly.
package diag

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// Severity orders diagnostics from the most to the least serious.
type Severity uint8

// String returns the name used in messages: error, warning or info.
func (s Severity) String() string {
	switch s {
	case Error:
		return "error"
	case Warning:
		return "warning"
	default:
		return "info"
	}
}

// String formats the position as file:line:col, leaving out unknown parts.
func (p Pos) String() string {
	var b strings.Builder
	b.WriteString(p.File)
	if p.Line > 0 {
		if b.Len() > 0 {
			b.WriteByte(':')
		}
		b.WriteString(strconv.Itoa(p.Line))
		if p.Col > 0 {
			b.WriteByte(':')
			b.WriteString(strconv.Itoa(p.Col))
		}
	}
	return b.String()
}

// String formats the diagnostic like a compiler message: "file:line:col: error: message".
func (d Diag) String() string {
	if p := d.Pos.String(); p != "" {
		return p + ": " + d.Sev.String() + ": " + d.Msg
	}
	return d.Sev.String() + ": " + d.Msg
}

// Error returns the formatted diagnostic.
func (d Diag) Error() string { return d.String() }

// Errorf builds an error-severity Diag, for parsers that return a single error.
func Errorf(pos Pos, format string, args ...any) Diag {
	return Diag{Sev: Error, Pos: pos, Msg: fmt.Sprintf(format, args...)}
}

// List collects diagnostics.
type List []Diag

// Add appends a diagnostic with a printf-style message.
func (l *List) Add(sev Severity, pos Pos, format string, args ...any) {
	*l = append(*l, Diag{Sev: sev, Pos: pos, Msg: fmt.Sprintf(format, args...)})
}

// Errorf, Warnf and Infof append a diagnostic of that severity.
func (l *List) Errorf(pos Pos, format string, args ...any) { l.Add(Error, pos, format, args...) }
func (l *List) Warnf(pos Pos, format string, args ...any)  { l.Add(Warning, pos, format, args...) }
func (l *List) Infof(pos Pos, format string, args ...any)  { l.Add(Info, pos, format, args...) }

// AddErr appends err, keeping its position when it is a Diag.
func (l *List) AddErr(err error) {
	if err == nil {
		return
	}
	if d, ok := err.(Diag); ok {
		*l = append(*l, d)
		return
	}
	*l = append(*l, Diag{Sev: Error, Msg: err.Error()})
}

// Count returns the number of diagnostics of the given severity.
func (l List) Count(sev Severity) int {
	n := 0
	for _, d := range l {
		if d.Sev == sev {
			n++
		}
	}
	return n
}

// HasErrors reports whether the list holds an error-severity diagnostic.
func (l List) HasErrors() bool { return l.Count(Error) > 0 }

// Sort orders the list by file (in order of first appearance), then by line
// and column; the original order breaks ties.
func (l List) Sort() {
	order := map[string]int{}
	for _, d := range l {
		if _, ok := order[d.Pos.File]; !ok {
			order[d.Pos.File] = len(order)
		}
	}
	sort.SliceStable(l, func(i, j int) bool {
		a, b := l[i].Pos, l[j].Pos
		if a.File != b.File {
			return order[a.File] < order[b.File]
		}
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		return a.Col < b.Col
	})
}

// Print writes one diagnostic per line.
func (l List) Print(w io.Writer) {
	for _, d := range l {
		fmt.Fprintln(w, d.String())
	}
}

// Summary returns a one-line count of errors, warnings and notes.
func (l List) Summary() string {
	return fmt.Sprintf("%d error(s), %d warning(s), %d note(s)", l.Count(Error), l.Count(Warning), l.Count(Info))
}
