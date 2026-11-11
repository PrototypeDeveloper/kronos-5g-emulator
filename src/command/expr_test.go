package command

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestLexerOperators(t *testing.T) {
	lx := NewLexer("", "a==1 != <= >= && || ! + - * / % < > = -3 a-5")
	var got []string
	for tok := lx.Next(); tok.Kind != TokEOF; tok = lx.Next() {
		got = append(got, tok.Text)
	}
	want := "a == 1 != <= >= && || ! + - * / % < > = - 3 a - 5"
	if strings.Join(got, " ") != want {
		t.Fatalf("tokens %q", got)
	}
	for src, want := range map[string]string{"a & b": "and is written &&", "a | b": "or is written ||"} {
		lx := NewLexer("", src)
		lx.Next()
		if tok := lx.Next(); tok.Kind != TokIllegal || !strings.Contains(tok.Text, want) {
			t.Fatalf("%s: %+v", src, tok)
		}
	}
}

// evalString parses and evaluates src on env.
func evalString(t *testing.T, env Env, src string) (Value, error) {
	t.Helper()
	lx := NewLexer("", src)
	e, err := ParseExpr(lx)
	if err != nil {
		t.Fatalf("%s: %v", src, err)
	}
	if tok := lx.Next(); tok.Kind != TokEOF {
		t.Fatalf("%s: left %s", src, tok.Describe())
	}
	return EvalExpr(env, e)
}

func TestEvalExpr(t *testing.T) {
	env := &fakeEnv{} // a is 7, s is "test"
	for src, want := range map[string]Value{
		"1 + 2 * 3":                  IntVal(7),
		"(1 + 2) * 3":                IntVal(9),
		"a - 5 - 1":                  IntVal(1),
		"a-5":                        IntVal(2),
		"-a + 1":                     IntVal(-6),
		"-(a + 0.5)":                 FloatVal(-7.5),
		"-(5s + 1s)":                 DurVal(-6 * time.Second),
		"!(a > 1)":                   BoolVal(false),
		"7 / 2":                      IntVal(3),
		"7 % 3":                      IntVal(1),
		"7 / 2.0":                    FloatVal(3.5),
		"a + 0.5":                    FloatVal(7.5),
		`s + "add"`:                  StrVal("testadd"),
		`"ue" + a`:                   StrVal("ue7"),
		`$imsi + "-" + 1.5`:          StrVal("208930000000001-1.5"),
		"5s + 500ms":                 DurVal(5500 * time.Millisecond),
		"5s * 2":                     DurVal(10 * time.Second),
		"2 * 5s":                     DurVal(10 * time.Second),
		"5s / 2":                     DurVal(2500 * time.Millisecond),
		"5s / 2s":                    FloatVal(2.5),
		"-5s":                        DurVal(-5 * time.Second),
		"a == 7":                     BoolVal(true),
		"a == 7.0":                   BoolVal(true),
		"a != 7":                     BoolVal(false),
		"a > 1 && a < 10":            BoolVal(true),
		"a < 1 || a >= 7":            BoolVal(true),
		"!(a == 7)":                  BoolVal(false),
		`s == "test"`:                BoolVal(true),
		`s < "u"`:                    BoolVal(true),
		"1s < 500ms":                 BoolVal(false),
		"a == 1 || a == 7 && s == s": BoolVal(true), // && binds tighter
		"1 + 2 == 3":                 BoolVal(true),
	} {
		got, err := evalString(t, env, src)
		if err != nil || got != want {
			t.Errorf("%s = %v (%s), %v; want %v (%s)", src, got, got.Type, err, want, want.Type)
		}
	}
	// && and || do not evaluate a right side that does not matter.
	for _, src := range []string{"a == 1 && 1 / 0 == 0", "a == 7 || 1 / 0 == 0"} {
		if _, err := evalString(t, env, src); err != nil {
			t.Errorf("%s: %v", src, err)
		}
	}
	for src, want := range map[string]string{
		"a / 0":   "division by zero",
		"a % 0":   "division by zero",
		"1.5 / 0": "division by zero",
		"5s / 0":  "division by zero",
		"$nope":   "no builtin",
	} {
		if _, err := evalString(t, env, src); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%s: want %q, got %v", src, want, err)
		}
	}
}

func TestExprString(t *testing.T) {
	for src, want := range map[string]string{
		"a==1&&(s!=\"x\"||!(a>2))": `a == 1 && (s != "x" || !(a > 2))`,
		"-1 + -a":                  "-1 + -a",
	} {
		e, err := ParseExpr(NewLexer("", src))
		if err != nil || e.String() != want {
			t.Errorf("%s: %q %v", src, e, err)
		}
	}
	ln, err := ParseLine("set(a = a + 5)")
	if err != nil || ln.Call.String() != "set(a = a + 5)" || ln.Call.Args[0].Kind != ExprAssign {
		t.Fatalf("set: %+v %v", ln.Call, err)
	}
	for in, want := range map[string]string{
		"set($state = 1)":      "only a variable can be assigned ($state is not one)",
		"set(a + 1 = 2)":       "only a variable can be assigned (a + 1 is not one)",
		"snap(\"%d\", (a + 1)": "an argument must be followed by , or )",
		"set(a = (1 + 2":       "( needs a closing )",
		"set(a = )":            "a value is needed",
	} {
		if _, err := ParseLine(in); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%s: want %q, got %v", in, want, err)
		}
	}
}

func TestSet(t *testing.T) {
	env := &fakeEnv{scope: ScopeUE}
	for _, line := range []string{"set(a = a + 5)", "set(a = 2 * a)", `set(s = s + "add")`, "set(f = a)"} {
		ln, err := ParseLine(line)
		if err != nil {
			t.Fatal(err)
		}
		if line == "set(f = a)" {
			env.vars["f"] = FloatVal(0)
		}
		if err := Exec(env, ln.Call); err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	if got := strings.Join(env.log, "; "); got != `set a = 12; set a = 24; set s = "testadd"; set f = 24` {
		t.Fatalf("log: %s", got)
	}
	if env.vars["f"] != FloatVal(24) {
		t.Fatalf("an int stored in a float variable: %+v", env.vars["f"])
	}
	ln, _ := ParseLine("set(a = a / 0)")
	if err := Exec(env, ln.Call); err == nil || env.vars["a"] != IntVal(24) {
		t.Fatalf("division by zero: %v, a = %v", err, env.vars["a"])
	}
}

func TestCheckExpr(t *testing.T) {
	// The fake resolver: a is a public int, s a private string, t a private
	// duration and n a public duration.
	for line, want := range map[string]Scope{
		"set(a = a + 5)":       ScopeConsole,
		"set(a = 2)":           ScopeConsole,
		"set(s = s + \"add\")": ScopeUE,
		"target(208930000000001).set(s = \"mod\")": ScopeUE,
		"set(t = n * 2)":              ScopeUE,
		`snap("%v %t", a + 1, a > 2)`: ScopeUE,
		"wait(n + 1s)":                ScopeConsole,
	} {
		scope, ds := checkLine(t, line)
		if ds.HasErrors() || scope != want {
			t.Errorf("%s: scope %v, diags %v", line, scope, ds)
		}
	}
	for line, want := range map[string]string{
		"target(208930000000001).set(a = 1)": "a is public (one value for all), so set cannot take target",
		"set(a = \"x\")":                     "a is of type int, so \"x\" (string) cannot be assigned",
		"set(a = 1.5)":                       "a is of type int, so 1.5 (float)",
		"set(a = s)":                         "private variable s cannot be used here",
		"set(a = $imsi)":                     "$imsi cannot be used here",
		"set(x = 1)":                         "variable x is not declared",
		"set(1)":                             "set takes an assignment",
		"set(a = a + s)":                     "private variable s cannot be used here",
		"set(s = s - \"x\")":                 "- cannot be applied to string and string",
		"set(s = s + (a > 1))":               "+ cannot be applied to string and bool",
		"snap(\"%v\", a = 1)":                "= assigns a variable only in set(a = …)",
		"snap(\"%v\", a && 1)":               "&& joins comparisons",
		"snap(\"%v\", !a)":                   "! takes a comparison",
		"snap(\"%v\", -s)":                   "- cannot be applied to string",
		"snap(\"%v\", a % 1.5)":              "% cannot be applied to int and float",
	} {
		_, ds := checkLine(t, line)
		if !ds.HasErrors() || !strings.Contains(fmt.Sprint(ds), want) {
			t.Errorf("%s: want %q, got %v", line, want, ds)
		}
	}
	// In a scenario set runs in its block; node sets public variables only.
	for line, want := range map[string]string{
		"set(s = \"x\")": "set in node: s is private",
		"set(a = $ue)":   "$ue belongs to each subscriber",
	} {
		c, err := ParseCall(NewLexer("", line))
		if err != nil {
			t.Fatal(err)
		}
		_, _, ds := Check(c, Context{Where: InScenario, Scope: ScopeNode}, fakeResolver{})
		if !ds.HasErrors() || !strings.Contains(fmt.Sprint(ds), want) {
			t.Errorf("%s: want %q, got %v", line, want, ds)
		}
	}
	c, _ := ParseCall(NewLexer("", "set(a = a + 1)"))
	if _, subject, ds := Check(c, Context{Where: InScenario, Scope: ScopeUE}, fakeResolver{}); ds.HasErrors() || subject != ScopeUE {
		t.Fatalf("a subscriber sets a public variable: %v %v", subject, ds)
	}
	// Conditions.
	for src, want := range map[string]string{"a": "the condition of if must be a comparison such as a == 1 (a is of type int)", "a == 1": ""} {
		e, _ := ParseExpr(NewLexer("", src))
		ds := CheckCond(e, fakeResolver{}, ScopeUE)
		if want == "" && ds.HasErrors() || want != "" && !strings.Contains(fmt.Sprint(ds), want) {
			t.Errorf("if (%s): %v", src, ds)
		}
	}
}
