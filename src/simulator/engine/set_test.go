package engine_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"kronos/models"
)

const setScenario = `
var {
	public count = 0
	public mode = 1
	private n = 10
	private s = "x"
	private f = 1.5
}
node {
	entry {
		send(NGSetupRequest)
		set(count = count + 100)
	}
}
start {
	entry {
		set(count = count + 1)
		set(n = n + $ue)
		set(s = s + "-" + $ue)
		set(f = f * 2)
		if (mode == 1) {
			snap("ue %d mode one n=%d s=%s f=%v", $ue, n, s, f)
		} else if (mode == 2) {
			snap("ue %d mode two", $ue)
		} else {
			if ($ue == 1 && mode > 2) {
				snap("ue %d mode other", $ue)
			}
		}
		if (mode == 9) {
		} else if ($ue == 1 && 1 / (mode - mode) == 0) {
			snap("never")
		}
		if ($ue == 1) { snap("after %d", mode) }
	}
}
`

// waitAll waits for snap events with every one of texts, in any order.
func waitAll(t *testing.T, hs *harness, texts ...string) {
	t.Helper()
	want := map[string]bool{}
	for _, s := range texts {
		want[s] = true
	}
	deadline := time.After(5 * time.Second)
	for len(want) > 0 {
		select {
		case e := <-hs.events:
			if e.Kind == models.KindSnap {
				if e.Text == "never" {
					t.Fatal("the body of a condition that failed ran")
				}
				delete(want, e.Text)
			}
		case <-deadline:
			t.Fatalf("no snap %v", want)
		}
	}
}

// publicVar returns the current value of a public variable as written.
func publicVar(t *testing.T, hs *harness, name string) string {
	t.Helper()
	for _, v := range hs.eng.Vars() {
		if v.Name == name {
			return v.Value
		}
	}
	t.Fatalf("no variable %s", name)
	return ""
}

func TestScenarioSetAndIf(t *testing.T) {
	const subs = 20
	hs := start(t, true, subs, setScenario)
	hs.linkup(t)
	var texts []string
	for i := 1; i <= subs; i++ {
		texts = append(texts, fmt.Sprintf("ue %d mode one n=%d s=x-%d f=3", i, 10+i, i))
	}
	waitAll(t, hs, append(texts, "after 1")...)
	// Every subscriber's set(count = count + 1) counts: none is lost.
	if got := publicVar(t, hs, "count"); got != fmt.Sprint(100+subs) {
		t.Fatalf("count = %s", got)
	}

	exec := func(line, want string) {
		t.Helper()
		out, err := hs.eng.Exec(line)
		if err != nil || out != want {
			t.Fatalf("%s: %q, %v (want %q)", line, out, err, want)
		}
	}
	// A public variable from the CUI/GUI: set once, and the scenario branches on it.
	exec("set(mode = mode + 1)", "mode = 2")
	exec("target(208930000000003).move(start)", "")
	waitAll(t, hs, "ue 3 mode two")
	exec("set(mode = 3)", "mode = 3")
	exec("target(208930000000001).move(start)", "")
	waitAll(t, hs, "ue 1 mode other", "after 3")
	// A private one: on the subscriber chosen by target.
	exec("target(208930000000001).set(n = n * 2)", "")
	exec(`target(208930000000001).snap("%d %s", n, s)`, "")
	waitAll(t, hs, "24 x-1-1")
	if got := publicVar(t, hs, "count"); got != fmt.Sprint(100+subs+2) {
		t.Fatalf("count = %s", got)
	}

	// Errors at run time leave the variable as it was.
	if _, err := hs.eng.Exec("set(count = count / 0)"); err == nil || !strings.Contains(err.Error(), "division by zero") {
		t.Fatalf("division by zero: %v", err)
	}
	for line, want := range map[string]string{
		"target(208930000000001).set(count = 1)": "count is public (one value for all), so set cannot take target",
		"set(mode = n)":                          "private variable n cannot be used here",
		`set(s = 1)`:                             "s is of type string",
	} {
		if _, err := hs.eng.Exec(line); err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("%s: want %q, got %v", line, want, err)
		}
	}
	if got := publicVar(t, hs, "count"); got != fmt.Sprint(100+subs+2) {
		t.Fatalf("count = %s", got)
	}
}

// TestScenarioIfError checks that a condition that cannot be evaluated is an error event and the body goes on.
func TestScenarioIfError(t *testing.T) {
	hs := start(t, true, 1, setScenario)
	hs.linkup(t)
	ev := hs.wait(t, "if error", func(e models.Event) bool { return e.Kind == models.KindError })
	if ev.Text != "if ($ue == 1 && 1 / (mode - mode) == 0): division by zero" || ev.Handler != "start/entry" {
		t.Fatalf("error event %+v", ev)
	}
	waitAll(t, hs, "after 1")
}
