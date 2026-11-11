// Package command is the command system shared by scenario.txt actions and
// the CUI/GUI command line. Each command lives in its own file
// and registers a Spec from init(); the parser, the validator, the engine,
// the CUI and the GUI all work from the registry, so adding a command does
// not require changes anywhere else.
package command

import (
	"fmt"
	"sort"
	"strings"
)

// Kind is the kind of a parameter. It drives validation (does the name
// exist?) and the GUI's choice of input widget.
type Kind uint8

// Scope is the set of subjects a command can run on. In scenario.txt it
// decides which blocks may use the command; on the command line it decides
// whether `target(...)` applies.
type Scope uint8

// String lists the subjects, e.g. "node/subscriber".
func (s Scope) String() string {
	var parts []string
	if s&ScopeNode != 0 {
		parts = append(parts, "node")
	}
	if s&ScopeUE != 0 {
		parts = append(parts, "subscriber")
	}
	if s&ScopeConsole != 0 {
		parts = append(parts, "console")
	}
	if s&ScopeBySignal != 0 {
		parts = append(parts, "decided by the signal")
	}
	if s&ScopeByVar != 0 {
		parts = append(parts, "decided by the variable (public: node, private: subscriber)")
	}
	if s&ScopeByTarget != 0 {
		return "the block in a scenario; on the CUI/GUI the subscriber of target(), or without it the node and every subscriber"
	}
	if s&ScopeByAssign != 0 {
		parts = append(parts, "the block in a scenario; on the CUI/GUI decided by the variable (public: once, private: subscriber)")
	}
	return strings.Join(parts, "/")
}

// Where says where a command may be written.
type Where uint8

// Usage renders the call signature, e.g. "snap(format, value...)".
func (s Spec) Usage() string {
	parts := make([]string, len(s.Args))
	for i, a := range s.Args {
		p := a.Name
		if a.Variadic {
			p += "..."
		}
		if a.Default != "" {
			p += "=" + a.Default
		}
		if a.Optional && !a.Variadic {
			p = "[" + p + "]"
		}
		parts[i] = p
	}
	return s.Name + "(" + strings.Join(parts, ", ") + ")"
}

// argCount returns the minimum and maximum number of arguments (max -1 = unbounded).
func (s Spec) argCount() (min, max int) {
	for _, a := range s.Args {
		if !a.Optional {
			min++
		}
		if a.Variadic {
			return min, -1
		}
		max++
	}
	return min, max
}

// argAt returns the parameter that the i-th argument binds to.
func (s Spec) argAt(i int) Arg {
	if i < len(s.Args) {
		return s.Args[i]
	}
	return s.Args[len(s.Args)-1]
}

// registry holds the commands registered by the init functions, by name.
var registry = map[string]Spec{}

// Register adds a command. It panics on a malformed or duplicate Spec, which
// is a programming error caught by the first test run.
func Register(s Spec) {
	switch {
	case s.Name == "" || s.Run == nil && s.RunCall == nil || s.Scope == 0 || s.Where == 0:
		panic(fmt.Sprintf("command: incomplete spec %q", s.Name))
	case s.Name == TargetKeyword:
		panic("command: \"target\" is reserved for selecting subscribers")
	}
	if _, dup := registry[s.Name]; dup {
		panic(fmt.Sprintf("command: %q registered twice", s.Name))
	}
	for i, a := range s.Args {
		last := i == len(s.Args)-1
		if a.Variadic && !last {
			panic(fmt.Sprintf("command: %s: only the last parameter may be variadic", s.Name))
		}
		if i > 0 && s.Args[i-1].Optional && !a.Optional {
			panic(fmt.Sprintf("command: %s: required parameter after an optional one", s.Name))
		}
	}
	registry[s.Name] = s
}

// Lookup returns the command registered under name.
func Lookup(name string) (Spec, bool) {
	s, ok := registry[name]
	return s, ok
}

// All returns every command, sorted by name (for help and the GUI).
func All() []Spec {
	out := make([]Spec, 0, len(registry))
	for _, s := range registry {
		out = append(out, s)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// Names returns the registered command names, sorted.
func Names() []string {
	all := All()
	out := make([]string, len(all))
	for i, s := range all {
		out[i] = s.Name
	}
	return out
}
