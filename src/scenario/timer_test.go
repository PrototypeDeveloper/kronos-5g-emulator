package scenario

import (
	"strings"
	"testing"

	"kronos/models"
)

const timerScenario = `
var {
	private t = 5s
	private u = 1s
	public n = 2s
	private i = 1
}
node {
	entry { send(NGSetupRequest) timer(n) }
	expire(n) { snap("node timer") }
}
any {
	expire(u) { snap("u in any") }
}
start {
	entry { timer(t) timer(u) }
	receive(Paging) { stoptimer(t) timer(t) }
	expire(t) { snap("5 s passed") move(start) }
}
`

func TestTimerScenario(t *testing.T) {
	f, ds := ParseScenario("t.txt", []byte(timerScenario))
	ds = append(ds, Validate(f, readSample(t), models.ValidateOptions{SnapEnabled: true})...)
	if ds.HasErrors() || strings.Contains(joinDiags(ds), "warning") {
		t.Fatalf("diagnostics:\n%s", joinDiags(ds))
	}
	start, _ := f.State("start")
	if x, ok := start.Expire("t"); !ok || len(x.Body) != 2 || x.Body[0].Call.Name != "snap" {
		t.Fatalf("expire(t) %+v", x)
	}
	if _, ok := f.Any.Expire("u"); !ok {
		t.Fatal("no expire(u) in any")
	}
	if _, ok := f.Node.Expire("n"); !ok {
		t.Fatal("no expire(n) in node")
	}
	if _, ok := start.Expire("u"); ok {
		t.Fatal("expire(u) found in start")
	}
}

func TestTimerErrors(t *testing.T) {
	set := readSample(t)
	const vars = `var { private t = 5s public n = 1s private i = 1 } `
	for src, want := range map[string]string{
		vars + `start { expire(x) { } }`:                                  "variable x is not declared",
		vars + `start { entry { timer(i) } expire(i) { } }`:               "i is not a duration variable",
		vars + `start { entry { timer(t) } expire(n) { } }`:               "a subscriber's timers are private variables (n is public",
		vars + `node { entry { timer(n) } expire(t) { } } start { }`:      "the node's timers are public variables (t is private)",
		vars + `start { entry { timer(t) } expire(t) { } expire(t) { } }`: "expire(t) is written twice in start",
		vars + `start { entry { timer(n) } }`:                             "timer: a subscriber's timers are private variables (n is public",
		vars + `node { entry { timer(t) } } start { }`:                    "timer: the node's timers are public variables (t is private",
		vars + `start { entry { timer(5s) } }`:                            "timer takes the name of a variable",
		vars + `start { entry { timer(x) } }`:                             "variable x is not declared",
		vars + `start { entry { stoptimer(i) } }`:                         "stoptimer: i is not a duration variable",
		vars + `start { expire t { } }`:                                   "expire must be followed by (",
		vars + `start { expire(5s) { } }`:                                 "expire() takes the timer variable",
		vars + `expire { }`:                                               "the reserved word expire cannot name a state",
	} {
		f, ds := ParseScenario("t.txt", []byte(src))
		ds = append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...)
		if !ds.HasErrors() || !strings.Contains(joinDiags(ds), want) {
			t.Errorf("%s\nwant %q, got:\n%s", src, want, joinDiags(ds))
		}
	}
	// An expire that no timer starts is a warning; a timer without an expire a note.
	f, ds := ParseScenario("t.txt", []byte(vars+`start { expire(t) { } receive(Paging) { timer(t) } } any { receive(UEContextReleaseCommand) { timer(t) } }`))
	if ds = append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...); ds.HasErrors() {
		t.Fatal(joinDiags(ds))
	}
	f, ds = ParseScenario("t.txt", []byte(vars+`start { expire(t) { } entry { timer(t) } } idle { entry { timer(t) } } any { expire(u) {} } var { private u = 1s }`))
	text := joinDiags(append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...))
	if !strings.Contains(text, "warning: expire(u) never runs from the scenario") {
		t.Fatalf("no warning:\n%s", text)
	}
	// An expire with an error gets no warning on top of it.
	f, ds = ParseScenario("t.txt", []byte(vars+`start { expire(i) { } expire(zz) { } }`))
	if text := joinDiags(append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...)); strings.Contains(text, "never runs") {
		t.Fatalf("warning on a bad expire:\n%s", text)
	}
	f, ds = ParseScenario("t.txt", []byte(vars+`start { entry { timer(t) } }`))
	if text := joinDiags(append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...)); !strings.Contains(text, "timer(t) has no expire(t)") {
		t.Fatalf("no note:\n%s", text)
	}
}
