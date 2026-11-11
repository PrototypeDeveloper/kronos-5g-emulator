package command

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"kronos/common/diag"
	"kronos/models"
)

func TestLexer(t *testing.T) {
	lx := NewLexer("f", "send(X) # comment\n  snap(\"a%d\\n\", $ue, 1.5, 5s, -3, 208930000000001, 5GMMStatus) // c")
	var got []string
	for tok := lx.Next(); tok.Kind != TokEOF; tok = lx.Next() {
		got = append(got, tok.Kind.String()+":"+tok.Text)
	}
	want := []string{"name:send", "(:(", "name:X", "):)", "name:snap", "(:(", "string:a%d\n", ",:,", "builtin:ue", ",:,",
		"float:1.5", ",:,", "duration:5s", ",:,", "integer:-3", ",:,", "integer:208930000000001", ",:,", "name:5GMMStatus", "):)"}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Fatalf("tokens\n got %q\nwant %q", got, want)
	}
	tok := NewLexer("f", "\n  \"abc").Next()
	if tok.Kind != TokIllegal || tok.Pos.Line != 2 || tok.Pos.Col != 3 {
		t.Fatalf("unterminated string: %+v", tok)
	}
	if v := NewLexer("", "1m30s").Next(); v.Kind != TokDuration || v.Val.Dur != 90*time.Second {
		t.Fatalf("duration: %+v", v)
	}
}

func TestParseLine(t *testing.T) {
	for _, tc := range []struct {
		in, target, call string
	}{
		{"send(NGSetupRequest)", "", "send(NGSetupRequest)"},
		{"target(208930000000001).send(X)", "208930000000001", "send(X)"},
		{`target("208930000000001").move(start)  # comment`, "208930000000001", "move(start)"},
		{`snap("%s=%d", $imsi, a)`, "", `snap("%s=%d", $imsi, a)`},
		{"help()", "", "help()"},
		{"linkup", "", "linkup()"},
		{"ping(10.10.10.20, 3)", "", "ping(10.10.10.20, 3)"},
		{"  linkdown  # stop", "", "linkdown()"},
	} {
		ln, err := ParseLine(tc.in)
		if err != nil {
			t.Fatalf("%s: %v", tc.in, err)
		}
		if ln.Target != tc.target || ln.Call.String() != tc.call {
			t.Fatalf("%s: got target %q call %q", tc.in, ln.Target, ln.Call)
		}
	}
	for _, in := range []string{"", "   ", "# only a comment"} {
		if _, err := ParseLine(in); !errors.Is(err, ErrEmptyLine) {
			t.Fatalf("%q: want ErrEmptyLine, got %v", in, err)
		}
	}
	for in, want := range map[string]string{
		"target(1,2).send(X)":                "only one IMSI",
		"target(1)send(X)":                   "followed by . and a command",
		"target(1)":                          "followed by . and a command",
		"send(X).move(Y)":                    "only one command per line",
		"send(X) extra":                      "unexpected text",
		"send(":                              "an argument is needed",
		`snap("x)`:                           "unterminated string",
		"target(x).send(X)":                  "takes an IMSI",
		"target(1).target(2)":                "goes before a command",
		"send X":                             "must be followed by (",
		"snap(\"%d\", 99999999999999999999)": "too large",
	} {
		_, err := ParseLine(in)
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("%q: want error containing %q, got %v", in, want, err)
		}
	}
}

// fakeResolver knows two signals, two states, two variables and one subscriber.
type fakeResolver struct{}

func (fakeResolver) Signal(name string) (Scope, bool) {
	switch name {
	case "NGSetupRequest":
		return ScopeNode, true
	case "InitialUEMessageRegistrationRequest":
		return ScopeUE, true
	}
	return 0, false
}
func (fakeResolver) State(name string) bool { return name == "start" || name == "registered" }
func (fakeResolver) Var(name string) (VarInfo, bool) {
	switch name {
	case "a":
		return VarInfo{Type: TypeInt, Public: true}, true
	case "s":
		return VarInfo{Type: TypeString}, true
	}
	return VarInfo{}, false
}
func (fakeResolver) Subscriber(imsi string) bool { return imsi == "208930000000001" }
func (fakeResolver) Names(k Kind) []string {
	switch k {
	case KindSignal:
		return []string{"NGSetupRequest", "InitialUEMessageRegistrationRequest"}
	case KindState:
		return []string{"start", "registered"}
	}
	return []string{"a", "s"}
}

func checkLine(t *testing.T, line string) (Scope, diag.List) {
	t.Helper()
	ln, err := ParseLine(line)
	if err != nil {
		t.Fatalf("%s: %v", line, err)
	}
	_, scope, ds := CheckLine(ln, fakeResolver{})
	return scope, ds
}

func TestCheckLine(t *testing.T) {
	for line, want := range map[string]Scope{
		"send(NGSetupRequest)":                                              ScopeNode,
		"send(InitialUEMessageRegistrationRequest)":                         ScopeUE,
		"target(208930000000001).send(InitialUEMessageRegistrationRequest)": ScopeUE,
		"target(208930000000001).move(registered)":                          ScopeUE,
		`snap("%d %s", a, s)`:                                               ScopeUE,
		`target(208930000000001).snap("%v", $state)`:                        ScopeUE,
		"help()":                           ScopeConsole,
		"linkup":                           ScopeConsole,
		"reload":                           ScopeConsole,
		"help(send)":                       ScopeConsole,
		"ping(10.10.10.20, 0, 0.2s, 1400)": ScopeUE,
		"ping(10.10.10.20, a)":             ScopeUE,
		"iperf(10.10.10.20)":               ScopeUE,
		"iperf(10.10.10.20, 13531, 100M, 10s, dl)":      ScopeUE,
		`iperf(10.10.10.20, 13531, "1.5G", 5s, ul)`:     ScopeUE,
		"iperf(10.10.10.20, 13531, 500k, 1s, ul, 1200)": ScopeUE,
		"stop":                  ScopeUE,
		"stop(iperf)":           ScopeUE,
		"iperftcp(10.10.10.20)": ScopeUE,
		"iperftcp(10.10.10.20, 13531, 0, 5s, dl)":       ScopeUE,
		"iperftcp(10.10.10.20, 13531, 1G, 5s, ul, 100)": ScopeUE,
	} {
		scope, ds := checkLine(t, line)
		if ds.HasErrors() || scope != want {
			t.Fatalf("%s: scope %v, diags %v", line, scope, ds)
		}
	}
	for line, want := range map[string]string{
		"target(208930000000001).send(NGSetupRequest)":                      "is a node signal, so it cannot take target",
		"target(208930000000001).help()":                                    "cannot take target",
		"target(208930000000099).send(InitialUEMessageRegistrationRequest)": "no subscriber with IMSI",
		"target(12345).send(InitialUEMessageRegistrationRequest)":           "15-digit",
		"send(NGSetupReqest)":                                               "did you mean NGSetupRequest?",
		"move(regstered)":                                                   "did you mean registered?",
		"sned(X)":                                                           "did you mean send?",
		`snap("%d", a, s)`:                                                  "1 conversion(s) in the format but 2 value(s)",
		`snap("%d %[1]d", a)`:                                               "[n] are not supported",
		"snap(a)":                                                           "must be a string",
		`snap("%v", undefined)`:                                             "variable undefined",
		`snap("%v", $nope)`:                                                 "unknown builtin",
		"send(NGSetupRequest, X)":                                           "wrong number of arguments",
		"help(nope)":                                                        "unknown command nope",
		"ping(10.10.10.20, -1)":                                             "from 0 to 1000000",
		"ping(10.10.10.20, 1, 0s)":                                          "interval",
		"ping(10.10.10.20, 1, 1s, 70000)":                                   "size",
		"ping(nowhere)":                                                     "variable nowhere",
		`ping("x")`:                                                         "IPv4",
		"iperf(10.10.10.20, 0)":                                             "port",
		"iperf(10.10.10.20, 5201, 100m)":                                    "100 minutes",
		"iperf(10.10.10.20, 5201, fast)":                                    "cannot read the rate",
		"iperf(10.10.10.20, 5201, 1M, 1500ms)":                              "whole seconds",
		"iperf(10.10.10.20, 5201, 1M, 1s, up)":                              "one of ul, dl",
		"iperf(10.10.10.20, 5201, 1M, 1s, ul, 8)":                           "len",
		"stop(all)":                                                         "one of ping, iperf",
		"iperf(10.10.10.20, 5201, 0)":                                       "must be positive",
		"iperftcp(10.10.10.20, 5201, 0, 1s, ul, 0)":                         "iperftcp len",
		"iperftcp(10.10.10.20, 5201, 0, 1s, ul, 2000000)":                   "iperftcp len",
	} {
		_, ds := checkLine(t, line)
		if !ds.HasErrors() || !strings.Contains(ds[0].Msg, want) {
			t.Fatalf("%s: want error %q, got %v", line, want, ds)
		}
	}
	// A type mismatch is only a warning: fmt still prints something.
	if _, ds := checkLine(t, `snap("%f", a)`); ds.HasErrors() || ds.Count(diag.Warning) != 1 {
		t.Fatalf("type mismatch: %v", ds)
	}
}

func TestCheckScenarioScopes(t *testing.T) {
	node := Context{Where: InScenario, Scope: ScopeNode}
	ue := Context{Where: InScenario, Scope: ScopeUE}
	for _, tc := range []struct {
		ctx  Context
		line string
		want string // "" = no error
	}{
		{node, "send(NGSetupRequest)", ""},
		{node, `snap("%d", a)`, ""},
		{ue, "send(InitialUEMessageRegistrationRequest)", ""},
		{ue, "move(start)", ""},
		{ue, `snap("%s %d %s", s, $ue, $imsi)`, ""},
		{node, "send(InitialUEMessageRegistrationRequest)", "is a UE signal, so the node block cannot send it"},
		{ue, "send(NGSetupRequest)", "is a node signal"},
		{node, "move(start)", "cannot be used in the node block"},
		{node, `snap("%s", s)`, "private variable s cannot be used in node"},
		{node, `snap("%d", $ue)`, "so node cannot use it"},
		{ue, "help()", "cannot be used in a scenario"},
		{ue, "target(1)", "target() cannot be used in a scenario"},
	} {
		c, err := ParseCall(NewLexer("", tc.line))
		if err != nil {
			t.Fatal(err)
		}
		_, _, ds := Check(c, tc.ctx, fakeResolver{})
		switch {
		case tc.want == "" && ds.HasErrors():
			t.Fatalf("%s: unexpected %v", tc.line, ds)
		case tc.want != "" && (!ds.HasErrors() || !strings.Contains(ds[0].Msg, tc.want)):
			t.Fatalf("%s: want %q, got %v", tc.line, tc.want, ds)
		}
	}
}

type fakeEnv struct {
	scope  Scope
	log    []string
	failOn string
}

func (e *fakeEnv) Scope() Scope { return e.scope }
func (e *fakeEnv) Send(s string) error {
	if s == e.failOn {
		return errors.New("boom")
	}
	e.log = append(e.log, "send "+s)
	return nil
}
func (e *fakeEnv) Move(s string) error { e.log = append(e.log, "move "+s); return nil }
func (e *fakeEnv) Print(s string)      { e.log = append(e.log, "print "+s) }
func (e *fakeEnv) Var(n string) (Value, error) {
	if n == "a" {
		return IntVal(7), nil
	}
	return StrVal("test"), nil
}
func (e *fakeEnv) SetVar(string, Value) error { return nil }
func (e *fakeEnv) LinkUp() error              { e.log = append(e.log, "linkup"); return nil }
func (e *fakeEnv) LinkDown() error            { e.log = append(e.log, "linkdown"); return nil }
func (e *fakeEnv) Reload() error              { e.log = append(e.log, "reload"); return nil }
func (e *fakeEnv) Ping(a models.PingArgs) error {
	e.log = append(e.log, fmt.Sprintf("ping %s %d %s %d", a.Dst, a.Count, a.Interval, a.Size))
	return nil
}
func (e *fakeEnv) Iperf(a models.IperfArgs) error {
	e.log = append(e.log, fmt.Sprintf("iperf %s:%d %d %s %v %d %v", a.Dst, a.Port, a.Rate, a.Duration, a.Reverse, a.Length, a.TCP))
	return nil
}
func (e *fakeEnv) Stop(kind string) error {
	e.log = append(e.log, "stop "+kind)
	return nil
}
func (e *fakeEnv) Sleep(d time.Duration) error {
	e.log = append(e.log, "sleep "+d.String())
	return nil
}
func (e *fakeEnv) Builtin(n string) (Value, error) {
	if n == "imsi" {
		return StrVal("208930000000001"), nil
	}
	return Value{}, errors.New("no builtin " + n)
}

func TestExec(t *testing.T) {
	env := &fakeEnv{scope: ScopeUE, failOn: "Bad"}
	for _, line := range []string{"send(X)", "move(start)", `snap("a=%d s=%s imsi=%s d=%v f=%.1f", a, s, $imsi, 5s, 1.25)`} {
		ln, _ := ParseLine(line)
		if err := Exec(env, ln.Call); err != nil {
			t.Fatal(err)
		}
	}
	for _, line := range []string{"ping(10.10.10.20, 3)", "ping(10.10.10.20, 0, 0.2s, 1400)", "iperf(10.10.10.20)",
		"iperf(10.10.10.20, 13531, 100M, 5s, dl, 1200)", `iperf(10.10.10.20, 13531, "1.5G")`, "stop", "stop(iperf)",
		"iperftcp(10.10.10.20)", "iperftcp(10.10.10.20, 13531, 0, 5s, dl, 65536)", "iperftcp(10.10.10.20, 13531, 2G)"} {
		ln, _ := ParseLine(line)
		if err := Exec(env, ln.Call); err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	want := []string{"send X", "move start", "print a=7 s=test imsi=208930000000001 d=5s f=1.2",
		"ping 10.10.10.20 3 0s -1", "ping 10.10.10.20 0 200ms 1400",
		"iperf 10.10.10.20:5201 0 0s false 0 false", "iperf 10.10.10.20:13531 100000000 5s true 1200 false", "iperf 10.10.10.20:13531 1500000000 0s false 0 false",
		"stop ", "stop iperf",
		"iperf 10.10.10.20:5201 0 0s false 0 true", "iperf 10.10.10.20:13531 0 5s true 65536 true", "iperf 10.10.10.20:13531 2000000000 0s false 0 true"}
	if strings.Join(env.log, "|") != strings.Join(want, "|") {
		t.Fatalf("got %q", env.log)
	}
	ln, _ := ParseLine("send(Bad)")
	if err := Exec(env, ln.Call); err == nil {
		t.Fatal("send error not returned")
	}
	ln, _ = ParseLine(`snap("%v", $nope)`)
	if err := Exec(env, ln.Call); err == nil {
		t.Fatal("builtin error not returned")
	}
	env.log = nil
	ln, _ = ParseLine("help()")
	if err := Exec(env, ln.Call); err != nil || len(env.log) != 1 || !strings.Contains(env.log[0], "send(signal)") {
		t.Fatalf("help: %v %q", err, env.log)
	}
}

func TestRegistry(t *testing.T) {
	if got := strings.Join(Names(), ","); got != "help,iperf,iperftcp,linkdown,linkup,move,ping,reload,send,snap,stop,wait" {
		t.Fatalf("commands: %s", got)
	}
	s, _ := Lookup("snap")
	if s.Usage() != "snap(format, value...)" {
		t.Fatal(s.Usage())
	}
	h, _ := Lookup("help")
	if h.Usage() != "help([command])" {
		t.Fatal(h.Usage())
	}
	i, _ := Lookup("iperf")
	if i.Usage() != "iperf(dst, [port=5201], [rate=1M], [time=10s], [dir=ul], [len=1400])" {
		t.Fatal(i.Usage())
	}
	i, _ = Lookup("iperftcp")
	if i.Usage() != "iperftcp(dst, [port=5201], [rate=0], [time=10s], [dir=ul], [len=131072])" {
		t.Fatal(i.Usage())
	}
	defer func() {
		if recover() == nil {
			t.Fatal("registering target did not panic")
		}
	}()
	Register(Spec{Name: TargetKeyword, Scope: ScopeUE, Where: Interactive, Run: func(Env, []Value) error { return nil }})
}

func TestParseFormat(t *testing.T) {
	for f, n := range map[string]int{"plain": 0, "%d": 1, "100%%": 0, "%-8s|%08.3f|%x|%v": 4, "%+q %#v %T": 3} {
		verbs, err := ParseFormat(f)
		if err != nil || len(verbs) != n {
			t.Fatalf("%q: %d verbs, %v", f, len(verbs), err)
		}
	}
	for _, f := range []string{"%", "%5", "%*d", "%[2]d", "%w", "%p", "%!"} {
		if _, err := ParseFormat(f); err == nil {
			t.Fatalf("%q: want error", f)
		}
	}
}

func TestSuggest(t *testing.T) {
	cands := []string{"DownlinkNASTransport", "UplinkNASTransport", "NGSetupResponse"}
	if s := Suggest("DownlinkNasTransprt", cands); s != "DownlinkNASTransport" {
		t.Fatal(s)
	}
	if s := Suggest("Completely", cands); s != "" {
		t.Fatal(s)
	}
}
