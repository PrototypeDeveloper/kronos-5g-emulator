package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestFieldExpressions(t *testing.T) {
	// The path: names joined by dots, with [index].
	for src, want := range map[string]string{
		"$nas.Cause":      "$nas.Cause",
		"$ngap.Cause.Nas": "$ngap.Cause.Nas",
		"$ngap.PDUSessionResourceSetupListSUReq[0].PDUSessionID": "$ngap.PDUSessionResourceSetupListSUReq[0].PDUSessionID",
		"$nas.Cause==22&&$ngap.AMFUENGAPID>0":                    "$nas.Cause == 22 && $ngap.AMFUENGAPID > 0",
	} {
		e, err := ParseExpr(NewLexer("", src))
		if err != nil || e.String() != want {
			t.Errorf("%s: %q %v", src, e, err)
		}
	}
	for src, want := range map[string]string{
		"$nas.":      "must be followed by the name of an IE",
		"$nas.X[":    "takes the index of a list item",
		"$nas.X[-1]": "takes the index of a list item",
		"$nas.X[0":   "needs a closing ]",
		"$ngap.3":    "must be followed by the name of an IE",
	} {
		if _, err := ParseExpr(NewLexer("", src)); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%s: want %q, got %v", src, want, err)
		}
	}
	// Checks: the type shows when it runs, so most uses pass.
	ue := Context{Where: InScenario, Scope: ScopeUE}
	for line, want := range map[string]string{
		`snap("%d %s", $nas.Cause, $ngap.Cause.Nas)`: "",
		"set(a = $nas.Cause)":                        "",
		`set(s = "cause " + $nas.Cause)`:             "",
		"set(a = $nas.Cause + 1)":                    "",
		`snap("%v", $nas)`:                           "$nas needs the name of an IE, e.g. $nas.Cause",
		`snap("%v", $imsi.Cause)`:                    "$imsi has no IEs",
		`snap("%v", $nas.Cause && 1)`:                "&& joins comparisons",
		`snap("%v", $nas.Cause - "x")`:               "- cannot be applied to IE value and string",
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
	e, _ := ParseExpr(NewLexer("", "$nas.Cause == 22"))
	if ds := CheckCond(e, fakeResolver{}, ScopeUE); ds.HasErrors() {
		t.Fatal(ds)
	}
	e, _ = ParseExpr(NewLexer("", "$nas.Cause"))
	if ds := CheckCond(e, fakeResolver{}, ScopeUE); !strings.Contains(fmt.Sprint(ds), "must be a comparison") {
		t.Fatal(ds)
	}
	if _, ds := checkLine(t, `wait($nas.T)`); !strings.Contains(fmt.Sprint(ds), "cannot be used here") {
		t.Fatalf("console: %v", ds)
	}

	// Values: a missing IE equals nothing and cannot be computed with.
	env := &fakeEnv{scope: ScopeUE, fields: map[string]Value{"$nas.Cause": IntVal(22), "$ngap.Cause.Nas": StrVal("normal-release")}}
	for src, want := range map[string]Value{
		"$nas.Cause == 22":                    BoolVal(true),
		"$nas.Cause + 1":                      IntVal(23),
		`$ngap.Cause.Nas == "normal-release"`: BoolVal(true),
		`"c" + $nas.Cause`:                    StrVal("c22"),
		"$nas.T3512 == 22":                    BoolVal(false),
		"$nas.T3512 != 22":                    BoolVal(true),
		"$nas.T3512 > 0 || $nas.Cause > 0":    BoolVal(true),
	} {
		if got, err := evalString(t, env, src); err != nil || got != want {
			t.Errorf("%s = %v, %v; want %v", src, got, err, want)
		}
	}
	for _, src := range []string{"$nas.T3512 + 1", "-$nas.T3512", `$ngap.Cause.Nas == 1`} {
		if _, err := evalString(t, env, src); err == nil {
			t.Errorf("%s: no error", src)
		}
	}
	ln, _ := ParseLine("set(a = $nas.T3512)")
	if err := Exec(env, ln.Call); err == nil || !strings.Contains(err.Error(), "no such IE") {
		t.Fatalf("set from a missing IE: %v", err)
	}
	ln, _ = ParseLine("set(a = $nas.Cause)")
	if err := Exec(env, ln.Call); err != nil || env.vars["a"] != IntVal(22) {
		t.Fatalf("set from an IE: %v %v", err, env.vars["a"])
	}
}
