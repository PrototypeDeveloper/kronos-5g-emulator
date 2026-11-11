package command

import (
	"fmt"
	"strconv"
	"time"

	"kronos/common/diag"
	"kronos/models"
)

// init registers iperf, iperftcp and stop.
func init() {
	// iperf: a UDP test through the PDU session.
	Register(Spec{
		Name: "iperf",
		Args: []Arg{
			{Name: "dst", Kind: KindValue},
			{Name: "port", Kind: KindValue, Optional: true, Default: "5201"},
			{Name: "rate", Kind: KindRate, Optional: true, Default: "1M"},
			{Name: "time", Kind: KindValue, Optional: true, Default: "10s"},
			{Name: "dir", Kind: KindChoice, Optional: true, Choices: []string{"ul", "dl"}, Default: "ul"},
			{Name: "len", Kind: KindValue, Optional: true, Default: "1400"},
		},
		Scope: ScopeUE,
		Where: InScenario | Interactive,
		Help: "run a UDP test with an iperf3 server (iperf3 -s) on the DN through the PDU session: port (default 5201), rate (default 1M; 100M, 1G …), " +
			"time (default 10s), dir ul (uplink, default) or dl (downlink, iperf3 -R), len the UDP payload (default 1400)",
		Check: checkIperf(false),
		Run: func(env Env, args []Value) error {
			a, err := iperfArgs(args, false)
			if err != nil {
				return err
			}
			return env.Iperf(a)
		},
	})
	// iperftcp: a TCP test through the UE's TUN device.
	Register(Spec{
		Name: "iperftcp",
		Args: []Arg{
			{Name: "dst", Kind: KindValue},
			{Name: "port", Kind: KindValue, Optional: true, Default: "5201"},
			{Name: "rate", Kind: KindRate, Optional: true, Default: "0", ZeroRate: true},
			{Name: "time", Kind: KindValue, Optional: true, Default: "10s"},
			{Name: "dir", Kind: KindChoice, Optional: true, Choices: []string{"ul", "dl"}, Default: "ul"},
			{Name: "len", Kind: KindValue, Optional: true, Default: strconv.Itoa(models.IperfDefaultTCPLength)},
		},
		Scope: ScopeUE,
		Where: InScenario | Interactive,
		Help: "run a TCP test with an iperf3 server (iperf3 -s) on the DN through the UE's TUN device (gnodeb needs CAP_NET_ADMIN: sudo, or make setcap): " +
			"port (default 5201), rate (default 0, as fast as TCP goes; 100M, 1G …), time (default 10s), " +
			"dir ul (uplink, default) or dl (downlink, iperf3 -R), len the octets per write (default 131072)",
		Check: checkIperf(true),
		Run: func(env Env, args []Value) error {
			a, err := iperfArgs(args, true)
			if err != nil {
				return err
			}
			return env.Iperf(a)
		},
	})
	// stop: ends a running ping or iperf.
	Register(Spec{
		Name:  "stop",
		Args:  []Arg{{Name: "kind", Kind: KindChoice, Optional: true, Choices: []string{"ping", "iperf"}}},
		Scope: ScopeUE,
		Where: InScenario | Interactive,
		Help:  "stop the running ping / iperf / iperftcp (all without kind; iperf stops iperftcp too); each reports what it has",
		Run: func(env Env, args []Value) error {
			kind := ""
			if len(args) > 0 {
				kind = args[0].Str
			}
			return env.Stop(kind)
		},
	})
}

// iperfArgs reads the arguments of iperf, or of iperftcp with tcp.
func iperfArgs(args []Value, tcp bool) (models.IperfArgs, error) {
	a := models.IperfArgs{Port: models.IperfDefaultPort, TCP: tcp}
	if err := checkIPv4("the iperf destination", args[0]); err != nil {
		return a, err
	}
	a.Dst = args[0].Str
	if len(args) > 1 {
		if args[1].Type != TypeInt || args[1].Int < 1 || args[1].Int > iperfMaxPort {
			return a, fmt.Errorf("the iperf port is 1 to 65535 (found %s)", args[1].Literal())
		}
		a.Port = int(args[1].Int)
	}
	if len(args) > 2 && !(tcp && args[2].Type == TypeInt && args[2].Int == 0) { // 0: no limit (TCP)
		r, err := RateOf(args[2])
		if err != nil {
			return a, fmt.Errorf("iperf rate: %v", err)
		}
		if r > models.IperfMaxRate {
			return a, fmt.Errorf("the iperf rate is at most 100G (%s)", args[2].Literal())
		}
		a.Rate = r
	}
	if len(args) > 3 {
		d := args[3]
		if d.Type != TypeDuration || d.Dur < time.Second || d.Dur > models.IperfMaxDuration || d.Dur%time.Second != 0 {
			return a, fmt.Errorf("the iperf time is whole seconds from 1s to %s (found %s)", models.IperfMaxDuration, d.Literal())
		}
		a.Duration = d.Dur
	}
	if len(args) > 4 {
		a.Reverse = args[4].Str == "dl"
	}
	if len(args) > 5 && tcp {
		if args[5].Type != TypeInt || args[5].Int < 1 || args[5].Int > models.IperfMaxTCPLength {
			return a, fmt.Errorf("the iperftcp len is 1 to %d octets (found %s)", models.IperfMaxTCPLength, args[5].Literal())
		}
		a.Length = int(args[5].Int)
	} else if len(args) > 5 {
		if args[5].Type != TypeInt || args[5].Int < models.IperfMinLength || args[5].Int > models.IperfMaxLength {
			return a, fmt.Errorf("the iperf len is %d to %d octets (found %s)", models.IperfMinLength, models.IperfMaxLength, args[5].Literal())
		}
		a.Length = int(args[5].Int)
	}
	return a, nil
}

// checkIperf returns the static check of iperf, or of iperftcp with tcp.
func checkIperf(tcp bool) func(Call, Resolver, Scope) diag.List {
	return func(c Call, _ Resolver, _ Scope) diag.List {
		return checkLiterals(c, func(args []Value) error {
			_, err := iperfArgs(args, tcp)
			return err
		})
	}
}
