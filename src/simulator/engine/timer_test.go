package engine_test

import (
	"strings"
	"testing"
	"time"

	"kronos/models"
)

const timerScenario = `
var {
	private t = 80ms
	private u = 80ms
	public n = 80ms
}
node {
	entry { send(NGSetupRequest) timer(n) }
	expire(n) { snap("node timer expired") }
}
any {
	expire(u) { snap("u expired in any, state %s", $state) }
}
start {
	expire(t) { snap("t expired in start") }
}
other {
	expire(t) { snap("t expired in other") }
}
`

// none checks that no event with text comes within d.
func none(t *testing.T, hs *harness, d time.Duration, text string) {
	t.Helper()
	deadline := time.After(d)
	for {
		select {
		case e := <-hs.events:
			if strings.Contains(e.Text, text) {
				t.Fatalf("unexpected event %q", e.Text)
			}
		case <-deadline:
			return
		}
	}
}

func TestScenarioTimers(t *testing.T) {
	hs := start(t, true, 1, timerScenario)
	hs.linkup(t)
	snap := func(text string) {
		t.Helper()
		hs.wait(t, text, func(e models.Event) bool { return e.Kind == models.KindSnap && e.Text == text })
	}
	exec := func(line string) {
		t.Helper()
		if _, err := hs.eng.Exec("target(208930000000001)." + line); err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	// The node's timer (a public variable), started by node's entry.
	snap("node timer expired")

	// expire runs in the state the subscriber is in when the timer runs out.
	exec("timer(t)")
	hs.wait(t, "started", func(e models.Event) bool { return e.Text == "timer t started (80ms)" })
	exec("move(other)")
	snap("t expired in other")
	// A state without expire(u): any's runs.
	exec("timer(u)")
	snap("u expired in any, state other")

	// timer() again restarts it: one expiry, later.
	exec("move(start)")
	exec("timer(t)")
	time.Sleep(50 * time.Millisecond)
	exec("timer(t)")
	snap("t expired in start")
	none(t, hs, 150*time.Millisecond, "t expired")

	// stoptimer, and a linkdown, stop it.
	exec("timer(t)")
	exec("stoptimer(t)")
	hs.wait(t, "stopped", func(e models.Event) bool { return e.Text == "timer t stopped" })
	none(t, hs, 150*time.Millisecond, "t expired")
	exec("timer(t)")
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	none(t, hs, 150*time.Millisecond, "t expired")

	// A public variable is the node's timer, also from the command line; it takes no target.
	hs.linkup(t)
	snap("node timer expired")
	if _, err := hs.eng.Exec("timer(n)"); err != nil {
		t.Fatalf("timer(n): %v", err)
	}
	snap("node timer expired")
	if _, err := hs.eng.Exec("target(208930000000001).timer(n)"); err == nil || !strings.Contains(err.Error(), "cannot take target") {
		t.Fatalf("timer(n) with a target: %v", err)
	}
}
