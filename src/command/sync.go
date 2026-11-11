package command

import (
	"fmt"
	"net/netip"

	"kronos/common/diag"
	"kronos/models"
)

// init registers sync_send.
func init() {
	Register(Spec{
		Name: "sync_send",
		Args: []Arg{
			{Name: "kind", Kind: KindChoice, Choices: []string{models.SyncKindSync, models.SyncKindHO}},
			{Name: "dst", Kind: KindValue},
		},
		Scope: ScopeUE | ScopeNode,
		Where: InScenario | Interactive,
		Help: "send a message to another pseudo-call (JSON over HTTP, from gnodeb.http.src) at dst (address:port), " +
			"where sync_receive(kind) of the subscriber with the same IMSI (of the node, from node) runs. " +
			"SYNC: a plain synchronization; HO: a handover, which also hands the subscriber's mapping context over",
		Check: checkSyncSend,
		Run:   func(env Env, args []Value) error { return env.SyncSend(args[0].Str, args[1].Str) },
	})
}

// SyncDst checks the destination of sync_send: an IPv4 address and a port.
func SyncDst(v Value) error {
	if a, err := netip.ParseAddrPort(v.Str); v.Type != TypeString || err != nil || !a.Addr().Is4() || a.Port() == 0 {
		return fmt.Errorf("dst must be address:port, e.g. 127.0.0.1:12345 (found %s)", v.Literal())
	}
	return nil
}

// checkSyncSend: HO is a subscriber's; a literal dst must be address:port.
func checkSyncSend(c Call, r Resolver, scope Scope) diag.List {
	var ds diag.List
	if c.Args[0].Name == models.SyncKindHO && scope == ScopeNode {
		ds.Errorf(c.Args[0].Pos, "sync_send(HO) hands a subscriber over, so the node cannot send it (write it in a state block or any)")
	}
	switch e := c.Args[1]; {
	case e.Kind == ExprLiteral:
		if err := SyncDst(e.Val); err != nil {
			ds.Errorf(e.Pos, "sync_send: %v", err)
		}
	default:
		if t := exprType(e, r); t != TypeNone && t != TypeString {
			ds.Errorf(e.Pos, "sync_send: dst must be a string address:port (%s is of type %s)", e, t)
		}
	}
	return ds
}
