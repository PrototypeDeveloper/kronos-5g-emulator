package command

import (
	"fmt"
	"net/netip"

	"kronos/common/diag"
	"kronos/models"
)

// init registers ping.
func init() {
	Register(Spec{
		Name: "ping",
		Args: []Arg{
			{Name: "dst", Kind: KindValue},
			{Name: "count", Kind: KindValue, Optional: true, Default: "1"},
			{Name: "interval", Kind: KindValue, Optional: true, Default: "1s"},
			{Name: "size", Kind: KindValue, Optional: true, Default: "56"},
		},
		Scope: ScopeUE,
		Where: InScenario | Interactive,
		Help: "send ICMP echo requests through the PDU session (GTP-U): count of them (default 1; 0 until stop), one every interval (default 1s), " +
			"with a payload of size octets (default 56). The summary gives the loss and the RTT min/avg/max/mdev",
		Check: checkPing,
		Run: func(env Env, args []Value) error {
			a, err := pingArgs(args)
			if err != nil {
				return err
			}
			return env.Ping(a)
		},
	})
}

// pingArgs reads the arguments of ping and checks their ranges.
func pingArgs(args []Value) (models.PingArgs, error) {
	a := models.PingArgs{Count: 1, Size: -1}
	if err := checkIPv4("the ping destination", args[0]); err != nil {
		return a, err
	}
	a.Dst = args[0].Str
	if len(args) > 1 {
		if args[1].Type != TypeInt || args[1].Int < 0 || args[1].Int > pingMaxCount {
			return a, fmt.Errorf("the ping count is an integer from 0 to %d (0: until stop; found %s)", pingMaxCount, args[1].Literal())
		}
		a.Count = int(args[1].Int)
	}
	if len(args) > 2 {
		if args[2].Type != TypeDuration || args[2].Dur < pingMinInterval || args[2].Dur > pingMaxInterval {
			return a, fmt.Errorf("the ping interval is a duration from %s to %s (e.g. 0.2s; found %s)", pingMinInterval, pingMaxInterval, args[2].Literal())
		}
		a.Interval = args[2].Dur
	}
	if len(args) > 3 {
		if args[3].Type != TypeInt || args[3].Int < 0 || args[3].Int > pingMaxSize {
			return a, fmt.Errorf("the ping size is 0 to %d octets (found %s)", pingMaxSize, args[3].Literal())
		}
		a.Size = int(args[3].Int)
	}
	return a, nil
}

// checkIPv4 returns an error unless v is a string holding an IPv4 address.
func checkIPv4(what string, v Value) error {
	if a, err := netip.ParseAddr(v.Str); v.Type != TypeString || err != nil || !a.Is4() {
		return fmt.Errorf("%s must be an IPv4 address (found %s)", what, v.Literal())
	}
	return nil
}

// checkLiterals runs the run-time checks when every argument is known
// before running: literals, and the words of rate and choice arguments.
func checkLiterals(c Call, check func([]Value) error) diag.List {
	var ds diag.List
	s, _ := Lookup(c.Name)
	args := make([]Value, len(c.Args))
	for i, e := range c.Args {
		switch k := s.argAt(i).Kind; {
		case k == KindChoice:
			args[i] = StrVal(e.Name)
		case k == KindRate || e.Kind == ExprLiteral:
			args[i] = EvalLiteral(e)
		default:
			return nil // a variable or a builtin: checked when it runs
		}
	}
	if err := check(args); err != nil {
		ds.Errorf(c.Pos, "%v", err)
	}
	return ds
}

// checkPing is the static check of ping.
func checkPing(c Call, _ Resolver, _ Scope) diag.List {
	return checkLiterals(c, func(args []Value) error {
		_, err := pingArgs(args)
		return err
	})
}
