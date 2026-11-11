package command

import (
	"fmt"
	"strings"
)

// init registers help.
func init() {
	Register(Spec{
		Name:  "help",
		Args:  []Arg{{Name: "command", Kind: KindCommand, Optional: true}},
		Scope: ScopeConsole,
		Where: Interactive,
		Help:  "show the commands, or how to use one",
		Run:   runHelp,
	})
}

// runHelp prints how to use one command, or the list of all commands.
func runHelp(env Env, args []Value) error {
	if len(args) == 1 {
		s, ok := Lookup(args[0].Str)
		if !ok {
			return fmt.Errorf("unknown command %s", args[0].Str)
		}
		env.Print(HelpText(s))
		return nil
	}
	var b strings.Builder
	b.WriteString("Commands (target(IMSI).command(…) picks one subscriber; without it, all subscribers)\n")
	for _, s := range All() {
		fmt.Fprintf(&b, "  %-26s %s\n", s.Usage(), s.Help)
	}
	env.Print(strings.TrimSuffix(b.String(), "\n"))
	return nil
}

// HelpText describes one command for help(name) and the GUI.
func HelpText(s Spec) string {
	var where []string
	if s.Where&InScenario != 0 {
		where = append(where, "scenario")
	}
	if s.Where&Interactive != 0 {
		where = append(where, "CUI/GUI")
	}
	target := "no"
	if s.Scope&(ScopeUE|ScopeBySignal) != 0 {
		target = "yes (without it: all subscribers)"
		if s.Scope&ScopeBySignal != 0 {
			target = "yes for UE signals (without it: all subscribers)"
		}
	}
	return fmt.Sprintf("%s\n  %s\n  where: %s\n  runs on: %s\n  target(IMSI): %s",
		s.Usage(), s.Help, strings.Join(where, ", "), s.Scope, target)
}
