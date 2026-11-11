package engine_test

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/netip"
	"os"
	"path/filepath"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
	"kronos/scenario"
	"kronos/transport/peer"
)

// The target pseudo-call: no registration of its own, it waits for the handover.
const hoTargetScenario = `
node {
	entry { send(NGSetupRequest) }
	sync_receive(SYNC) { snap("node sync") }
}
start {
	sync_receive(HO) { snap("ho in %s", $state) move(handover) }
}
handover {
	sync_receive(SYNC) { snap("sync in %s", $state) }
}
`

// httpLine is the config line of gnodeb.http.src on the loopback.
func httpLine(port int) string {
	return fmt.Sprintf("http: { src: { address: 127.0.0.1, port: %d } }", port)
}

func TestHandoverBetweenPseudoCalls(t *testing.T) {
	// The target listens for the other pseudo-calls on its gnodeb.http.src.
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	dst := fmt.Sprintf("127.0.0.1:%d", port)
	tgt := startN3(t, true, 1, hoTargetScenario, freeUDPPort(t), freeUDPPort(t), httpLine(port))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go peer.NewServer(tgt.eng.DeliverSync).Serve(ctx, l)

	// The node takes SYNC only while it runs; an unknown IMSI is refused.
	cli := peer.NewClient(netip.MustParseAddrPort("127.0.0.1:0"))
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: models.SyncKindSync}); err == nil || !strings.Contains(err.Error(), "503") {
		t.Fatalf("SYNC before linkup: %v", err)
	}
	tgt.linkup(t)
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: models.SyncKindSync}); err != nil {
		t.Fatal(err)
	}
	tgt.wait(t, "node sync", func(e models.Event) bool { return e.Kind == models.KindSnap && e.Text == "node sync" })
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: models.SyncKindSync, IMSI: "208930000000099"}); err == nil || !strings.Contains(err.Error(), "404") {
		t.Fatalf("unknown IMSI: %v", err)
	}
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: "XN"}); err == nil || !strings.Contains(err.Error(), "400") {
		t.Fatalf("unknown kind: %v", err)
	}
	// A SYNC the subscriber's state has no handler for; a broken handover context.
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: models.SyncKindSync, IMSI: "208930000000001"}); err != nil {
		t.Fatal(err)
	}
	tgt.wait(t, "unhandled", func(e models.Event) bool {
		return e.Kind == models.KindWarn && e.Text == "sync_receive(SYNC) from 127.0.0.1:0: no sync_receive(SYNC) in start or any"
	})
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: models.SyncKindHO, IMSI: "208930000000001", Handover: &models.HOContext{GUTI: "zz"}}); err != nil {
		t.Fatal(err)
	}
	tgt.wait(t, "broken context", func(e models.Event) bool {
		return e.Kind == models.KindError && strings.Contains(e.Text, "the handover context: guti")
	})

	// The source registers its subscriber and sets up a PDU session.
	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	src := startN3(t, true, 1, sessionScenario, freeUDPPort(t), upf.LocalAddr().(*net.UDPAddr).Port, httpLine(freeUDPPort(t)))
	src.linkup(t)
	amfSec, ran := register(t, src)
	establish(t, src, amfSec, ran)

	// HO hands the mapping context over; then the target's sync_receive(HO) runs.
	if _, err := src.eng.Exec("target(208930000000001).sync_send(HO, " + dst + ")"); err != nil {
		t.Fatal(err)
	}
	src.wait(t, "accepted", func(e models.Event) bool { return e.Text == "sync_send(HO) to "+dst+": accepted" })
	tgt.wait(t, "taken over", func(e models.Event) bool {
		return e.Text == "handover context taken over: registered, NAS security on, 1 PDU session(s)"
	})
	tgt.wait(t, "the handler", func(e models.Event) bool { return e.Kind == models.KindSnap && e.Text == "ho in start" })
	st, _ := tgt.eng.UE("208930000000001")
	if st.RM != models.RMRegistered || len(st.Sessions) != 1 || st.Sessions[0].UEIP != "10.60.0.7" || st.State != "handover" {
		t.Fatalf("target subscriber %+v", st)
	}

	// The target goes on with the same NAS security context: its Service
	// Request has the next uplink COUNT and the 5G-S-TMSI of the 5G-GUTI.
	if _, err := tgt.eng.Exec("target(208930000000001).send(InitialUEMeesageServiceRequest)"); err != nil {
		t.Fatal(err)
	}
	ium := recvAs[*ngapMessage.InitialUEMessage](t, tgt.amf)
	nas, err := libnas.Decode(ium.NASPDU.Value, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink})
	if err != nil {
		t.Fatalf("Service Request at the target: %v", err)
	}
	if sr := nas.(*nasMessage.ServiceRequest); !bytes.HasSuffix(sr.FiveGSTMSI, []byte{0, 0, 0, 1}) {
		t.Fatalf("5G-S-TMSI %x", sr.FiveGSTMSI)
	}

	// SYNC for the subscriber runs the handler of the state it is in.
	if _, err := src.eng.Exec("target(208930000000001).sync_send(SYNC, " + dst + ")"); err != nil {
		t.Fatal(err)
	}
	tgt.wait(t, "sync", func(e models.Event) bool { return e.Kind == models.KindSnap && e.Text == "sync in handover" })
}

// appendSignals adds new signals (minimal messages of their types) to the
// signal.txt of h and reloads it.
func appendSignals(t *testing.T, h *harness, names ...string) {
	t.Helper()
	var b strings.Builder
	for _, n := range names {
		raw, err := scenario.NewSignal(n)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(&b, "\n%s {\n  %x\n}\n", n, raw)
	}
	f, err := os.OpenFile(filepath.Join(h.dir, "signal.txt"), os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(b.String())
	f.Close()
	if _, err := h.eng.Exec("reload"); err != nil {
		t.Fatal(err)
	}
}

// TestHandoverRequestAtTheTarget: after sync_send(HO) the AMF's
// HandoverRequest (a new AMF-UE-NGAP-ID, no RAN-UE-NGAP-ID) reaches the
// subscriber that took the context over, by the uplink TEID of its session.
func TestHandoverRequestAtTheTarget(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	dst := fmt.Sprintf("127.0.0.1:%d", port)
	tgt := startN3(t, true, 1, hoTargetScenario, freeUDPPort(t), freeUDPPort(t), httpLine(port))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go peer.NewServer(tgt.eng.DeliverSync).Serve(ctx, l)
	tgt.linkup(t)
	appendSignals(t, tgt, "HandoverRequestAcknowledge")

	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	src := startN3(t, true, 1, sessionScenario, freeUDPPort(t), upf.LocalAddr().(*net.UDPAddr).Port, httpLine(freeUDPPort(t)))
	src.linkup(t)
	amfSec, ran := register(t, src)
	establish(t, src, amfSec, ran)
	if _, err := src.eng.Exec("target(208930000000001).sync_send(HO, " + dst + ")"); err != nil {
		t.Fatal(err)
	}
	tgt.wait(t, "the handler", func(e models.Event) bool { return e.Kind == models.KindSnap && e.Text == "ho in start" })

	// The AMF of the target: HandoverRequest for the session (its uplink TEID is the source's).
	raw, _ := scenario.NewSignal("HandoverRequest")
	msg, _ := libngap.Decode(raw)
	req := msg.(*ngapMessage.HandoverRequest)
	req.AMFUENGAPID.Value = 200
	req.PDUSessionResourceSetupListHOReq.List = []ngapType.PDUSessionResourceSetupItemHOReq{{PDUSessionID: ngapType.PDUSessionID{Value: 1}, SNSSAI: snssai(), HandoverRequestTransfer: setupTransfer(t)}}
	tgt.amf.send(t, req)
	tgt.wait(t, "admitted", func(e models.Event) bool {
		return e.Source == "208930000000001" && strings.HasPrefix(e.Text, "handover request: RAN-UE-NGAP-ID ")
	})
	if _, err := tgt.eng.Exec("target(208930000000001).send(HandoverRequestAcknowledge)"); err != nil {
		t.Fatal(err)
	}
	ack := recvAs[*ngapMessage.HandoverRequestAcknowledge](t, tgt.amf)
	items := ack.PDUSessionResourceAdmittedList.List
	if ack.AMFUENGAPID.Value != 200 || ack.RANUENGAPID.Value == 0 || len(items) != 1 || items[0].PDUSessionID.Value != 1 {
		t.Fatalf("HandoverRequestAcknowledge %+v", ack)
	}
	var tr ngapType.HandoverRequestAcknowledgeTransfer
	if err := ngapcodec.DecodeParameter("HandoverRequestAcknowledgeTransfer", items[0].HandoverRequestAcknowledgeTransfer, &tr); ngapcodec.IsFatal(err) || len(tr.DLNGUUPTNLInformation.GTPTunnel.GTPTEID.Value) != 4 {
		t.Fatalf("admitted transfer %+v %v", tr, err)
	}
	st, _ := tgt.eng.UE("208930000000001")
	if len(st.Sessions) != 1 || !st.Sessions[0].Active {
		t.Fatalf("target sessions %+v", st.Sessions)
	}
	// The new association is known by its IDs: a partial NG Reset naming it releases it.
	raw, _ = scenario.NewSignal("NGReset")
	msg, _ = libngap.Decode(raw)
	part := msg.(*ngapMessage.NGReset)
	part.ResetType.Present = ngapType.ResetTypePresentPartOfNGInterface
	part.ResetType.PartOfNGInterface.List = []ngapType.UEAssociatedLogicalNGConnectionItem{{
		HasRANUENGAPID: true, RANUENGAPID: ack.RANUENGAPID, HasAMFUENGAPID: true, AMFUENGAPID: ngapType.AMFUENGAPID{Value: 200}}}
	tgt.amf.send(t, part)
	tgt.wait(t, "released", func(e models.Event) bool {
		return e.Source == "208930000000001" && e.Text == "UE association released (NG Reset from the AMF); the NAS state stays"
	})

	// At the source, the AMF's NG Reset of the whole interface ends the UE association.
	raw, _ = scenario.NewSignal("NGReset")
	msg, _ = libngap.Decode(raw)
	reset := msg.(*ngapMessage.NGReset)
	reset.ResetType.Present = ngapType.ResetTypePresentNGInterface
	src.amf.send(t, reset)
	src.wait(t, "reset", func(e models.Event) bool {
		return e.Source == "208930000000001" && e.Text == "UE association released (NG Reset from the AMF); the NAS state stays"
	})
}
