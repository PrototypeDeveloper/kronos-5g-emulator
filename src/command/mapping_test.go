package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestMappingCommands(t *testing.T) {
	// On the command line: the node and every subscriber, or the subscriber of target().
	for line, want := range map[string]Scope{
		"mapon":                                  ScopeConsole,
		"mapoff(plmn)":                           ScopeConsole,
		"mapget()":                               ScopeConsole,
		"target(208930000000001).mapoff(ngapid)": ScopeUE,
		"target(208930000000001).mapget":         ScopeUE,
	} {
		scope, ds := checkLine(t, line)
		if ds.HasErrors() || scope != want {
			t.Errorf("%s: scope %v, diags %v", line, scope, ds)
		}
	}
	for line, want := range map[string]string{
		"mapoff(nope)":                        "item must be one of plmn, tac",
		"mapoff(plmn, tac)":                   "wrong number of arguments",
		"target(208930000000099).mapon(plmn)": "no subscriber with IMSI",
	} {
		if _, ds := checkLine(t, line); !ds.HasErrors() || !strings.Contains(fmt.Sprint(ds), want) {
			t.Errorf("%s: want %q, got %v", line, want, ds)
		}
	}
	// In a scenario: the node in node, the subscriber in the others.
	for _, sc := range []Scope{ScopeNode, ScopeUE} {
		c, _ := ParseCall(NewLexer("", "mapoff(plmn)"))
		if _, subject, ds := Check(c, Context{Where: InScenario, Scope: sc}, fakeResolver{}); ds.HasErrors() || subject != sc {
			t.Errorf("scenario %v: %v %v", sc, subject, ds)
		}
	}

	env := &fakeEnv{scope: ScopeUE}
	for _, line := range []string{"mapget", "mapoff(plmn)", "mapoff(ngapid)", "mapget(plmn)", "mapget(tac)", "mapget", "mapon(plmn)", "mapon", "mapget", "mapoff", "mapget(res)"} {
		ln, err := ParseLine(line)
		if err != nil {
			t.Fatal(err)
		}
		if err := Exec(env, ln.Call); err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	want := "info mapping: every item on; info mapping off: plmn; info mapping off: ngapid; info plmn: off; info tac: on; " +
		"info mapping: plmn off, ngapid off; the others on; info mapping on: plmn; info mapping on: every item; " +
		"info mapping: every item on; info mapping off: every item; info res: off"
	if got := strings.Join(env.log, "; "); got != want {
		t.Fatalf("log:\n%s\nwant\n%s", got, want)
	}
}
