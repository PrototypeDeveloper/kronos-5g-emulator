package scenario

import (
	"strings"
	"testing"

	"kronos/command"
	"kronos/models"
)

const ifScenario = `
var {
	public a = 1
	private s = "x"
	private t = 1s
}
node {
	entry {
		send(NGSetupRequest)
		if (a == 1) { set(a = a + 5) }
	}
}
start {
	entry {
		set(s = s + "add")
		if (a == 1) {
			snap("one")
		} else if (a == 2 || s == "mod") {
			snap("two")
			if (!(a > 3)) { timer(t) }
		} else {
			set(s = "mod")
			snap("other %s", s)
		}
		snap("end")
	}
	expire(t) { }
}
`

func TestIfScenario(t *testing.T) {
	f, ds := ParseScenario("t.txt", []byte(ifScenario))
	ds = append(ds, Validate(f, readSample(t), models.ValidateOptions{SnapEnabled: true})...)
	if ds.HasErrors() {
		t.Fatal(joinDiags(ds))
	}
	// The timer started inside an if counts: no warning about expire(t).
	if strings.Contains(joinDiags(ds), "never runs") {
		t.Fatal(joinDiags(ds))
	}
	start, _ := f.State("start")
	got := strings.Join(StmtLines(start.Entry.Body), "\n")
	want := `set(s = s + "add")
if (a == 1) {
  snap("one")
} else if (a == 2 || s == "mod") {
  snap("two")
  if (!(a > 3)) {
    timer(t)
  }
} else {
  set(s = "mod")
  snap("other %s", s)
}
snap("end")`
	if got != want {
		t.Fatalf("lines:\n%s", got)
	}
	if got := start.Entry.Body[1].String(); got != `if (a == 1) { snap("one") } else if (a == 2 || s == "mod") { snap("two"), if (!(a > 3)) { timer(t) } } else { set(s = "mod"), snap("other %s", s) }` {
		t.Fatalf("one line: %s", got)
	}
	// snap is taken out of the bodies of if too.
	if n := f.StripCalls("snap"); n != 4 {
		t.Fatalf("stripped %d", n)
	}
	n := 0
	f.Calls(func(_ *Block, c command.Call) { n++ })
	if n != 5 { // send, set; set, timer, set
		t.Fatalf("%d calls left", n)
	}
}

func TestNegativeInit(t *testing.T) {
	f, ds := ParseScenario("t.txt", []byte("var { public a = -1 private d = -1.5 private t = -5s private s = \"-\" }\nstart { }"))
	if ds.HasErrors() || len(f.Vars) != 4 {
		t.Fatal(joinDiags(ds))
	}
	if got := f.Vars[0].Init.Literal() + " " + f.Vars[1].Init.Literal() + " " + f.Vars[2].Init.Literal(); got != "-1 -1.5 -5s" {
		t.Fatalf("initial values %s", got)
	}
	if _, ds := ParseScenario("t.txt", []byte(`var { public a = -"x" }`)); !strings.Contains(joinDiags(ds), "must be a literal") {
		t.Fatal(joinDiags(ds))
	}
}

func TestIfErrors(t *testing.T) {
	const vars = "var { public a = 1 private s = \"x\" private i = 2 }\n"
	set := readSample(t)
	for src, want := range map[string]string{
		vars + `start { if (a == 1) { } }`:                          "if can be written only inside receive(…) { … }, entry { … } or expire(…) { … }",
		vars + `start { else { } }`:                                 "else can be written only inside",
		vars + `start { entry { else { } } }`:                       "else without if",
		vars + `start { entry { if a == 1 { } } }`:                  "if must be followed by ( and a condition",
		vars + `start { entry { if (a = 1) { } } }`:                 "= assigns a variable (only in set(var = value)); compare with ==",
		vars + `start { entry { if (a == 1) snap("x") } }`:          "if (a == 1) must be followed by {",
		vars + `start { entry { if (a == 1) { } else snap("x") } }`: "else must be followed by { or if",
		vars + `start { entry { if (a == ) { } } }`:                 "a value is needed",
		vars + `start { entry { if (a) { } } }`:                     "the condition of if must be a comparison such as a == 1 (a is of type int)",
		vars + `start { entry { if (a + 1) { } } }`:                 "(a + 1 is of type int)",
		vars + `start { entry { if (s == 1) { } } }`:                "== cannot be applied to string and int",
		vars + `start { entry { if (x == 1) { } } }`:                "variable x is not declared",
		vars + `node { entry { if (s == "x") { } } } start { }`:     "private variable s cannot be used in node",
		vars + `node { entry { set(i = 1) } } start { }`:            "set in node: i is private",
		vars + `start { entry { set(a = "x") } }`:                   "a is of type int, so \"x\" (string)",
		vars + `start { entry { if (a == 1) { snap(1) } } }`:        "snap: format must be a string",
		vars + `if { }`:   "the reserved word if cannot name a state",
		vars + `else { }`: "the reserved word else cannot name a state",
	} {
		f, ds := ParseScenario("t.txt", []byte(src))
		if f != nil && !ds.HasErrors() {
			ds = append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...)
		}
		if !ds.HasErrors() || !strings.Contains(joinDiags(ds), want) {
			t.Errorf("%s\nwant %q, got:\n%s", src, want, joinDiags(ds))
		}
	}
	// A syntax error in one body does not hide the next handler.
	f, ds := ParseScenario("t.txt", []byte(vars+`start { entry { if (a = 1) { snap("x") } snap("y") } receive(Paging) { snap("p") } }`))
	if start, _ := f.State("start"); len(ds) != 1 || start == nil || len(start.Handlers) != 1 {
		t.Fatalf("recovery: %s", joinDiags(ds))
	}
}

// Everything after # (or //) up to the end of the line is a comment, in any language.
func TestCommentsAndControl(t *testing.T) {
	src := `var { public fuka = 1 }
start {
        entry {
                if(fuka == 1) {
                    control(2)  #JP1
                }
                send(InitialUEMeesageRegistrationRequest) // JP2
        }
}
`
	// Comments in Japanese (written as escapes: the source stays ASCII).
	src = strings.NewReplacer("JP1", "\u3053\u306e\u5834\u5408\u306f0.5s\u306bsend(InitialUE)\u304c\u52d5\u304f", "JP2", "\u767b\u9332").Replace(src)
	f, ds := ParseScenario("t.txt", []byte(src))
	ds = append(ds, Validate(f, readSample(t), models.ValidateOptions{SnapEnabled: true})...)
	if ds.HasErrors() {
		t.Fatal(joinDiags(ds))
	}
	start, _ := f.State("start")
	if got := strings.Join(StmtLines(start.Entry.Body), "|"); got != "if (fuka == 1) {|  control(2)|}|send(InitialUEMeesageRegistrationRequest)" {
		t.Fatalf("body %s", got)
	}
}

func TestSyncReceive(t *testing.T) {
	src := `node { sync_receive(SYNC) { snap("node") } }
start { sync_receive(HO) { snap("ho") } sync_receive(SYNC) { } }
any { sync_receive(SYNC) { sync_send(SYNC, 127.0.0.1:12345) } }
`
	f, ds := ParseScenario("t.txt", []byte(src))
	ds = append(ds, Validate(f, readSample(t), models.ValidateOptions{SnapEnabled: true, Sync: true})...)
	if ds.HasErrors() || strings.Contains(joinDiags(ds), "warning") {
		t.Fatal(joinDiags(ds))
	}
	start, _ := f.State("start")
	if x, ok := start.Sync("HO"); !ok || len(x.Body) != 1 {
		t.Fatalf("sync_receive(HO) %+v", x)
	}
	// Without gnodeb.http the scenario cannot talk to the others.
	f, ds = ParseScenario("t.txt", []byte(src))
	if text := joinDiags(append(ds, Validate(f, readSample(t), models.ValidateOptions{SnapEnabled: true})...)); !strings.Contains(text, "warning: sync_send / sync_receive need gnodeb.http.src") {
		t.Fatal(text)
	}
	for src, want := range map[string]string{
		`start { sync_receive(XN) { } }`:                          "sync_receive(XN): the kind is SYNC or HO",
		`node { sync_receive(HO) { } }`:                           "sync_receive(HO) in node",
		`start { sync_receive(SYNC) { } sync_receive(SYNC) { } }`: "sync_receive(SYNC) is written twice in start",
		`start { sync_receive SYNC { } }`:                         "sync_receive must be followed by (",
		`start { sync_receive("x") { } }`:                         "sync_receive() takes the kind of message",
		`sync_receive { }`:                                        "the reserved word sync_receive cannot name a state",
		`node { entry { sync_send(HO, 127.0.0.1:1) } }`:           "the node cannot send it",
	} {
		f, ds := ParseScenario("t.txt", []byte(src))
		if !ds.HasErrors() {
			ds = append(ds, Validate(f, readSample(t), models.ValidateOptions{Sync: true})...)
		}
		if !strings.Contains(joinDiags(ds), want) {
			t.Errorf("%s\nwant %q, got:\n%s", src, want, joinDiags(ds))
		}
	}
}
