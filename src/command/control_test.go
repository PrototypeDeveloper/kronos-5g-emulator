package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestControl(t *testing.T) {
	ue := Context{Where: InScenario, Scope: ScopeUE}
	for line, want := range map[string]string{
		"control(5)":       "",
		"control(0.5)":     "",
		"control(0)":       "",
		"control(a)":       "",
		"control(a * 2)":   "",
		"control(-1)":      "from 0 (no limit) to 1000000",
		"control(2000000)": "from 0 (no limit) to 1000000",
		`control("x")`:     "the passes per second must be a number",
		"control(s)":       "control takes the passes per second, e.g. control(5) (s is of type string)",
		"control(1s)":      "found the duration 1s",
		"control()":        "wrong number of arguments",
	} {
		c, err := ParseCall(NewLexer("", line))
		if err != nil {
			t.Fatal(err)
		}
		_, _, ds := Check(c, ue, fakeResolver{})
		if want == "" && ds.HasErrors() || want != "" && !strings.Contains(fmt.Sprint(ds), want) {
			t.Errorf("%s: want %q, got %v", line, want, ds)
		}
	}
	c, _ := ParseCall(NewLexer("", "control(5)"))
	if _, _, ds := Check(c, Context{Where: InScenario, Scope: ScopeNode}, fakeResolver{}); !strings.Contains(fmt.Sprint(ds), "cannot be used in the node block") {
		t.Errorf("node: %v", ds)
	}
	if _, ds := checkLine(t, "control(5)"); !strings.Contains(fmt.Sprint(ds), "cannot be used on the CUI/GUI (scenario only)") {
		t.Errorf("CUI: %v", ds)
	}
	// Each control() is told apart by its position; 0 waits for nothing.
	env := &fakeEnv{scope: ScopeUE}
	for _, src := range []string{"control(5)", "control(a + 1)", "control(0)"} {
		c, _ := ParseCall(NewLexer("s.txt", src))
		if err := Exec(env, c); err != nil {
			t.Fatal(err)
		}
	}
	if got := strings.Join(env.log, "; "); got != "control s.txt:1:1 200ms; control s.txt:1:1 125ms" {
		t.Fatalf("log: %s", got)
	}
}
