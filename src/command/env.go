package command

import (
	"fmt"
	"kronos/models"
	"time"
)

// Env is the execution environment a command sees. The engine implements it
// for a subscriber, the node or the console; command never imports the engine.
// The dispatcher resolves target(), so Run always works on a single subject.
type Env interface {
	Scope() Scope // ScopeUE, ScopeNode or ScopeConsole
	Send(signal string) error
	Move(state string) error
	Print(text string) // console output (snap, help)
	Var(name string) (Value, error)
	SetVar(name string, v Value) error
	Builtin(name string) (Value, error) // $imsi etc., without the "$"
	LinkUp() error                      // set up the SCTP association (linkup)
	LinkDown() error                    // shut the association down (linkdown)
	Reload() error                      // read signal.txt again (reload)
	Ping(a models.PingArgs) error       // ICMP echo through the subscriber's PDU session
	Iperf(a models.IperfArgs) error     // iperf3 UDP test through the subscriber's PDU session
	Stop(kind string) error             // stop the subscriber's ping / iperf ("" for both)
	Sleep(d time.Duration) error        // wait (command line and batches only)
}

var builtins = map[string]BuiltinInfo{
	"ue":    {TypeInt, ScopeUE, "the number of the subscriber (from 1)"},
	"imsi":  {TypeString, ScopeUE, "the IMSI of the subscriber"},
	"state": {TypeString, ScopeUE, "the current state of the subscriber"},
	"msg":   {TypeString, ScopeNode | ScopeUE, "the last trigger received (NGAP:NAS)"},
	"time":  {TypeString, ScopeNode | ScopeUE, "the current time"},
}

// LookupBuiltin returns the builtin named name (without "$").
func LookupBuiltin(name string) (BuiltinInfo, bool) {
	b, ok := builtins[name]
	return b, ok
}

// Eval evaluates the arguments of c for the parameters of s.
func Eval(env Env, s Spec, c Call) ([]Value, error) {
	out := make([]Value, 0, len(c.Args))
	for i, e := range c.Args {
		switch s.argAt(i).Kind {
		case KindValue:
			switch e.Kind {
			case ExprLiteral:
				out = append(out, e.Val)
			case ExprBuiltin:
				v, err := env.Builtin(e.Name)
				if err != nil {
					return nil, err
				}
				out = append(out, v)
			default:
				v, err := env.Var(e.Name)
				if err != nil {
					return nil, err
				}
				out = append(out, v)
			}
		case KindFormat:
			out = append(out, e.Val)
		case KindRate:
			if e.Kind == ExprIdent {
				out = append(out, StrVal(e.Name)) // 100M is read as a word
			} else {
				out = append(out, e.Val)
			}
		default: // names of signals, states, commands and choices
			out = append(out, StrVal(e.Name))
		}
	}
	return out, nil
}

// Exec looks up, evaluates and runs one call on env.
func Exec(env Env, c Call) error {
	s, ok := Lookup(c.Name)
	if !ok {
		return fmt.Errorf("unknown command %s", c.Name)
	}
	args, err := Eval(env, s, c)
	if err != nil {
		return err
	}
	return s.Run(env, args)
}
