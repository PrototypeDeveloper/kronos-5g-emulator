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
	// Assign stores the value eval returns in a variable. For a public
	// variable eval runs under the lock of the public variables, so
	// set(a = a + 1) from many subscribers adds every one.
	Assign(name string, eval func() (Value, error)) error
	Builtin(name string) (Value, error) // $imsi etc., without the "$"
	LinkUp() error                      // set up the SCTP association (linkup)
	LinkDown() error                    // shut the association down (linkdown)
	Reload() error                      // read signal.txt again (reload)
	// SetMapping turns the mapping of items on or off: of the subscriber,
	// of the node, or on the console of the node and every subscriber.
	SetMapping(items models.MapOff, on bool) error
	// Mapping returns of how many of its subjects (1, or on the console the
	// node and the subscribers) the mapping of item is off.
	Mapping(item models.MapItem) (off, of int)
	Info(text string)               // a result line (mapget): an info event, printed on the console
	Timer(name string) error        // start or restart the timer of a duration variable
	StopTimer(name string) error    // stop it
	Ping(a models.PingArgs) error   // ICMP echo through the subscriber's PDU session
	Iperf(a models.IperfArgs) error // iperf3 UDP test through the subscriber's PDU session
	Stop(kind string) error         // stop the subscriber's ping / iperf ("" for both)
	Sleep(d time.Duration) error    // wait (command line and batches only)
	// Control waits for the subscriber's turn at the control() written at
	// key: turns are interval apart, shared by all subscribers. It returns
	// ErrInterrupted when the wait is cut short (the link went down).
	Control(key string, interval time.Duration) error
	// SyncSend sends a message of kind (SYNC, HO) to the pseudo-call at dst
	// (address:port): for the subscriber, or from the node.
	SyncSend(kind, dst string) error
}

var builtins = map[string]BuiltinInfo{
	"ue":    {TypeInt, ScopeUE, "the number of the subscriber (from 1)"},
	"imsi":  {TypeString, ScopeUE, "the IMSI of the subscriber"},
	"state": {TypeString, ScopeUE, "the current state of the subscriber"},
	"msg":   {TypeString, ScopeNode | ScopeUE, "the last trigger received (NGAP:NAS)"},
	"time":  {TypeString, ScopeNode | ScopeUE | ScopeConsole, "the current time"},
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
			v, err := EvalExpr(env, e)
			if err != nil {
				return nil, err
			}
			out = append(out, v)
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
	if s.RunCall != nil {
		return s.RunCall(env, c)
	}
	args, err := Eval(env, s, c)
	if err != nil {
		return err
	}
	return s.Run(env, args)
}
