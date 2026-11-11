package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestSyncSend(t *testing.T) {
	// address:port needs no quotes.
	tok := NewLexer("", "127.0.0.1:12345)").Next()
	if tok.Kind != TokString || tok.Val.Str != "127.0.0.1:12345" {
		t.Fatalf("token %+v", tok)
	}
	if ln, err := ParseLine("target(208930000000001).sync_send(HO, 127.0.0.1:12345)"); err != nil || ln.String() != "target(208930000000001).sync_send(HO, 127.0.0.1:12345)" {
		t.Fatalf("%v %v", ln, err)
	}
	for line, want := range map[string]Scope{
		"sync_send(SYNC, 127.0.0.1:12345)":                       ScopeUE,
		`target(208930000000001).sync_send(HO, "10.0.0.1:8080")`: ScopeUE,
		"sync_send(SYNC, s)":                                     ScopeUE,
	} {
		scope, ds := checkLine(t, line)
		if ds.HasErrors() || scope != want {
			t.Errorf("%s: %v %v", line, scope, ds)
		}
	}
	for line, want := range map[string]string{
		"sync_send(XN, 127.0.0.1:12345)": "kind must be one of SYNC, HO",
		"sync_send(SYNC, 127.0.0.1)":     "dst must be address:port",
		`sync_send(SYNC, "x:1")`:         "dst must be address:port",
		"sync_send(SYNC, 127.0.0.1:0)":   "dst must be address:port",
		"sync_send(SYNC, a)":             "dst must be a string address:port (a is of type int)",
		"sync_send(SYNC)":                "wrong number of arguments",
	} {
		if _, ds := checkLine(t, line); !strings.Contains(fmt.Sprint(ds), want) {
			t.Errorf("%s: want %q, got %v", line, want, ds)
		}
	}
	c, _ := ParseCall(NewLexer("", "sync_send(HO, 127.0.0.1:12345)"))
	if _, _, ds := Check(c, Context{Where: InScenario, Scope: ScopeNode}, fakeResolver{}); !strings.Contains(fmt.Sprint(ds), "the node cannot send it") {
		t.Errorf("HO in node: %v", ds)
	}
	c, _ = ParseCall(NewLexer("", "sync_send(SYNC, 127.0.0.1:12345)"))
	if _, _, ds := Check(c, Context{Where: InScenario, Scope: ScopeNode}, fakeResolver{}); ds.HasErrors() {
		t.Errorf("SYNC in node: %v", ds)
	}
	env := &fakeEnv{scope: ScopeUE}
	if err := Exec(env, c); err != nil || strings.Join(env.log, ";") != "sync_send SYNC 127.0.0.1:12345" {
		t.Fatalf("%v %v", env.log, err)
	}
}
