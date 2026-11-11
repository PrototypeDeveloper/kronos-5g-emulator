package engine_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"net/netip"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/common/security"
	"kronos/models"
	"kronos/scenario"
	"kronos/simulator/engine"
	"kronos/simulator/gnb"
	"kronos/state"
	"kronos/transport/sctp"
)

// stubAMF is a loopback AMF for tests: it accepts associations and lets the
// test read requests and send answers.
type stubAMF struct {
	l    *sctp.Listener
	conn *sctp.Conn
}

func newStubAMF(t *testing.T) *stubAMF {
	t.Helper()
	l, err := sctp.Listen(netip.MustParseAddrPort("127.0.0.1:0"), sctp.Options{})
	if err != nil {
		t.Skipf("SCTP is not available here: %v", err)
	}
	t.Cleanup(func() { l.Close() })
	return &stubAMF{l: l}
}

func (s *stubAMF) accept(t *testing.T) {
	t.Helper()
	c, err := s.l.Accept()
	if err != nil {
		t.Fatal(err)
	}
	s.conn = c
	t.Cleanup(func() { c.Close() })
}

func (s *stubAMF) recv(t *testing.T) any {
	t.Helper()
	b, info, err := s.conn.Recv()
	if err != nil {
		t.Fatal(err)
	}
	if info.PPID != models.PPIDNGAP {
		t.Fatalf("ppid %d", info.PPID)
	}
	msg, err := libngap.Decode(b)
	if msg == nil {
		t.Fatal(err)
	}
	// Non-UE messages use stream 0, UE-associated ones the others.
	if ue := !strings.HasPrefix(fmt.Sprintf("%T", msg), "*ngapMessage.NGSetup"); ue == (info.Stream == 0) {
		t.Fatalf("%T on stream %d", msg, info.Stream)
	}
	return msg
}

func recvAs[T any](t *testing.T, s *stubAMF) T {
	t.Helper()
	msg := s.recv(t)
	m, ok := msg.(T)
	if !ok {
		var want T
		t.Fatalf("got %T, want %T", msg, want)
	}
	return m
}

func (s *stubAMF) send(t *testing.T, msg any) {
	t.Helper()
	b, err := libngap.Encode(msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.conn.Send(b, 0, models.PPIDNGAP); err != nil {
		t.Fatal(err)
	}
}

const configTemplate = `
base: { snap: %v }
ue:
  imsi: "208930000000001"
  subscription: %d
  imeisv: "4370816125816151"
  suci: { routingindicator: "0000" }
  credential: { key: 8baf473f2f8fd09487cccbd7097c6862, op: 8e27b6af0e692e750f32667a3b14605d, type: OPC, amf: "8000" }
  integrity: { nia0: true, nia1: true, nia2: true }
  ciphering: { nea0: true, nea1: true, nea2: true }
gnodeb:
  name: kronos-test-gnb
  gnodebid: "000123"
  gnodebidlength: 24
  ngap: { src: { address: 127.0.0.1, port: 0 }, dst: { address: 127.0.0.1, port: %d } }
  gtpu: { src: { address: 127.0.0.1, port: %d }, dst: { address: 127.0.0.1, port: %d } }
  plmn: { mcc: "208", mnc: "93" }
  snssai: [ { sst: 1, sd: "010203" }, { sst: 2 } ]
  tac: 7
`

const nodeScenario = `
var { public count = 0 }
node {
	entry {
		send(NGSetupRequest)
	}
	receive(NGSetupResponse) {
		snap("setup ok via %s, count=%d", $msg, count)
	}
	receive(NGSetupFailure) {
		snap("setup failed")
	}
}
`

type harness struct {
	eng    *engine.Engine
	events <-chan models.Event
	amf    *stubAMF
	dir    string // holds config.yaml and a copy of the sample signal.txt
}

func start(t *testing.T, snap bool, subscribers int, scen string) *harness {
	t.Helper()
	return startN3(t, snap, subscribers, scen, 2152, 2152)
}

// startN3 also sets the GTP-U ports: the gNB's (gtpu.src) and the stub UPF's (gtpu.dst).
func startN3(t *testing.T, snap bool, subscribers int, scen string, gtpuPort, upfPort int) *harness {
	t.Helper()
	eng, amf, dir := buildEngine(t, snap, subscribers, scen, gtpuPort, upfPort)
	events, unsubscribe := eng.Events().Subscribe(1024)
	ctx, cancel := context.WithCancel(context.Background())
	eng.Start(ctx)
	t.Cleanup(func() {
		cancel()
		eng.Wait()
		unsubscribe()
	})
	return &harness{eng: eng, events: events, amf: amf, dir: dir}
}

// buildEngine makes an engine (not started) for a stub AMF, with the
// config, a copy of the sample signal.txt and scen in a temporary directory.
func buildEngine(t *testing.T, snap bool, subscribers int, scen string, gtpuPort, upfPort int) (*engine.Engine, *stubAMF, string) {
	t.Helper()
	amf := newStubAMF(t)
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(fmt.Sprintf(configTemplate, snap, subscribers, amf.l.Addr().Port(), gtpuPort, upfPort)), 0o644); err != nil {
		t.Fatal(err)
	}
	sample, err := os.ReadFile("../../../config/signal.txt")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "signal.txt"), sample, 0o644); err != nil {
		t.Fatal(err)
	}
	cfg, ds := state.Load(path)
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	sigs, ds := scenario.ReadSignals(cfg.SignalPath())
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	f, ds := scenario.ParseScenario("test.txt", []byte(scen))
	ds = append(ds, scenario.Validate(f, sigs, models.ValidateOptions{SnapEnabled: snap})...)
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	role, err := gnb.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return engine.New(cfg, sigs, f, role), amf, dir
}

// linkup connects and answers the NG Setup that the node entry sends.
func (h *harness) linkup(t *testing.T) *ngapMessage.NGSetupRequest {
	t.Helper()
	if err := h.eng.LinkUp(); err != nil {
		t.Fatal(err)
	}
	h.amf.accept(t)
	req := recvAs[*ngapMessage.NGSetupRequest](t, h.amf)
	h.amf.send(t, setupResponse())
	h.wait(t, "NG Setup", func(e models.Event) bool { return strings.Contains(e.Text, "NG Setup succeeded") })
	return req
}

func (h *harness) wait(t *testing.T, what string, match func(models.Event) bool) models.Event {
	t.Helper()
	timeout := time.After(5 * time.Second)
	for {
		select {
		case e := <-h.events:
			if match(e) {
				return e
			}
		case <-timeout:
			t.Fatalf("timed out waiting for %s", what)
		}
	}
}

var plmn20893 = []byte{0x02, 0xf8, 0x39}

func setupResponse() *ngapMessage.NGSetupResponse {
	return &ngapMessage.NGSetupResponse{
		AMFName:             ngapType.AMFName{Value: "stub-amf"},
		ServedGUAMIList:     ngapType.ServedGUAMIList{List: []ngapType.ServedGUAMIItem{{GUAMI: guami()}}},
		RelativeAMFCapacity: ngapType.RelativeAMFCapacity{Value: 255},
		PLMNSupportList: ngapType.PLMNSupportList{List: []ngapType.PLMNSupportItem{{
			PLMNIdentity:     ngapType.PLMNIdentity{Value: plmn20893},
			SliceSupportList: ngapType.SliceSupportList{List: []ngapType.SliceSupportItem{{SNSSAI: snssai()}}},
		}}},
	}
}

func guami() ngapType.GUAMI {
	return ngapType.GUAMI{
		PLMNIdentity: ngapType.PLMNIdentity{Value: plmn20893},
		AMFRegionID:  ngapType.AMFRegionID{Value: gnb.BitString(0xca, 8)},
		AMFSetID:     ngapType.AMFSetID{Value: gnb.BitString(0x3f8, 10)},
		AMFPointer:   ngapType.AMFPointer{Value: gnb.BitString(0, 6)},
	}
}

func snssai() ngapType.SNSSAI {
	return ngapType.SNSSAI{SST: ngapType.SST{Value: []byte{1}}, HasSD: true, SD: ngapType.SD{Value: []byte{1, 2, 3}}}
}

func TestNGSetupAndLinkControl(t *testing.T) {
	h := start(t, true, 1, nodeScenario)
	if _, err := h.eng.Exec("send(NGSetupRequest)"); err == nil || !strings.Contains(err.Error(), "linkup") {
		t.Fatalf("send before linkup: %v", err)
	}
	if out, err := h.eng.Exec("linkdown"); err == nil {
		t.Fatalf("linkdown while down: %q", out)
	}
	req := h.linkup(t)

	// The node entry sent NGSetupRequest, patched from config.
	g := req.GlobalRANNodeID.GlobalGNBID
	if !bytes.Equal(g.PLMNIdentity.Value, plmn20893) || g.GNBID.GNBID.BitLength != 24 || gnb.BitValue(g.GNBID.GNBID) != 0x123 {
		t.Fatalf("GlobalGNBID %+v", g)
	}
	if !req.HasRANNodeName || req.RANNodeName.Value != "kronos-test-gnb" {
		t.Fatalf("RANNodeName %+v", req.RANNodeName)
	}
	ta := req.SupportedTAList.List
	if len(ta) != 1 || !bytes.Equal(ta[0].TAC.Value, []byte{0, 0, 7}) {
		t.Fatalf("SupportedTAList %+v", ta)
	}
	if s := ta[0].BroadcastPLMNList.List[0].TAISliceSupportList.List; len(s) != 2 || s[1].SNSSAI.HasSD {
		t.Fatalf("slices %+v", s)
	}
	snap := h.wait(t, "snap", func(e models.Event) bool { return e.Kind == models.KindSnap })
	if snap.Text != "setup ok via NGSetupResponse, count=0" {
		t.Fatalf("snap %q", snap.Text)
	}
	st := h.eng.Status()
	if st.Setup != models.SetupOK || st.PeerName != "stub-amf" || !st.Connected || st.PeerInfo[0] != "GUAMI 208/93-cafe00" {
		t.Fatalf("status %+v", st)
	}

	// CUI/GUI commands go through the same command system.
	if _, err := h.eng.Exec("send(NGSetupRequest)"); err != nil {
		t.Fatal(err)
	}
	recvAs[*ngapMessage.NGSetupRequest](t, h.amf)
	for line, want := range map[string]string{
		"target(208930000000001).send(NGSetupRequest)": "cannot take target",
		"send(NoSuchSignal)":                           "not in signal.txt",
		"linkup":                                       "already connected",
		"bogus(":                                       "an argument is needed",
	} {
		if _, err := h.eng.Exec(line); err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("%s: want %q, got %v", line, want, err)
		}
	}
	if out, err := h.eng.Exec("help"); err != nil || !strings.Contains(out, "linkup()") {
		t.Fatalf("help: %q %v", out, err)
	}

	// NG Setup Failure is recorded with its cause.
	h.amf.send(t, &ngapMessage.NGSetupFailure{
		Cause:         ngapType.Cause{Present: ngapType.CausePresentMisc, Misc: ngapType.CauseMisc{Value: 4}},
		HasTimeToWait: true, TimeToWait: ngapType.TimeToWait{Value: 2},
	})
	h.wait(t, "failure snap", func(e models.Event) bool { return e.Kind == models.KindSnap && e.Text == "setup failed" })
	if st := h.eng.Status(); st.Setup != models.SetupFailed || st.Cause != "misc(4) unknown-PLMN-or-SNPN, TimeToWait 5s" {
		t.Fatalf("status %+v", st)
	}

	// A UE-associated message for an unknown UE is reported.
	h.amf.send(t, &ngapMessage.UEContextReleaseCommand{
		UENGAPIDs: ngapType.UENGAPIDs{Present: ngapType.UENGAPIDsPresentAMFUENGAPID, AMFUENGAPID: ngapType.AMFUENGAPID{Value: 1}},
		Cause:     ngapType.Cause{Present: ngapType.CausePresentNas, Nas: ngapType.CauseNas{Value: 0}}})
	h.wait(t, "unknown UE", func(e models.Event) bool {
		return e.Kind == models.KindWarn && strings.Contains(e.Text, "no subscriber for")
	})

	// The AMF closing the association leaves the engine running; linkup reconnects.
	h.amf.conn.Close()
	h.wait(t, "peer close", func(e models.Event) bool { return e.Text == "the peer disconnected SCTP" })
	if st := h.eng.Status(); st.Connected {
		t.Fatalf("status after close %+v", st)
	}
	h.linkup(t)
	if _, err := h.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	h.wait(t, "linkdown", func(e models.Event) bool { return e.Text == "SCTP disconnected" })
}

func TestSnapDisabled(t *testing.T) {
	h := start(t, false, 1, nodeScenario)
	if err := h.eng.LinkUp(); err != nil {
		t.Fatal(err)
	}
	h.amf.accept(t)
	recvAs[*ngapMessage.NGSetupRequest](t, h.amf)
	h.amf.send(t, setupResponse())
	h.wait(t, "NG Setup result", func(e models.Event) bool {
		if e.Kind == models.KindSnap {
			t.Fatal("snap ran although base.snap is false")
		}
		return e.Kind == models.KindLink && strings.Contains(e.Text, "NG Setup succeeded")
	})
	if _, err := h.eng.Exec(`snap("x")`); err == nil || !strings.Contains(err.Error(), "snap is disabled") {
		t.Fatalf("interactive snap: %v", err)
	}
}

func TestDialFailure(t *testing.T) {
	h := start(t, true, 1, nodeScenario)
	h.amf.l.Close()
	if err := h.eng.LinkUp(); err == nil {
		t.Fatal("linkup succeeded without an AMF")
	}
	h.wait(t, "error event", func(e models.Event) bool {
		return e.Kind == models.KindError && strings.Contains(e.Text, "SCTP connection failed")
	})
}

const registrationScenario = `
node {
	entry { send(NGSetupRequest) }
}
start {
	entry { send(InitialUEMeesageRegistrationRequest) }
	receive(DownLinkNasTransport:AuthenticationRequest) {
		send(UpLinkNASTransportAuthenticationResponse)
	}
	receive(DownLinkNasTransport:SecurityModeCommand) {
		send(UpLinkNASTransportSecurityModeComplete)
	}
	receive(InitialContextSetupRequest:RegistrationAccept) {
		send(InitialContextSetupResponse)
		send(UplinkNASTransportRegistrationComplete)
		move(registered)
	}
}
registered {
	entry { snap("%s registered as UE%d", $imsi, $ue) }
}
`

func h(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// TestRegistrationWithStubAMF plays the AMF side of 5G AKA, the Security
// Mode procedure (NIA2/NEA2) and Registration Accept against the scenario.
// The stub derives its keys with common/security as well, so
// this checks the NAS mapping end to end; a real core checks the derivations.
func TestRegistrationWithStubAMF(t *testing.T) {
	hs := start(t, true, 1, registrationScenario)
	hs.linkup(t) // NG Setup success starts the subscriber: start/entry sends the registration
	amfSec, _ := register(t, hs)
	hs.wait(t, "registered", func(e models.Event) bool {
		return e.Kind == models.KindSnap && e.Text == "208930000000001 registered as UE1"
	})

	// A command for one subscriber: protected with the next uplink COUNT.
	if _, err := hs.eng.Exec("target(208930000000001).send(UplinkNASTransportRegistrationComplete)"); err != nil {
		t.Fatal(err)
	}
	ul := recvAs[*ngapMessage.UplinkNASTransport](t, hs.amf)
	if _, err := libnas.Decode(ul.NASPDU.Value, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink}); err != nil {
		t.Fatalf("second RegistrationComplete: %v", err)
	}

	// linkdown puts the subscriber back to its initial state.
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "reset", func(e models.Event) bool { return strings.Contains(e.Text, "back to the initial state") })
}

// register plays the AMF side of the registration the start entry begins:
// 5G AKA, Security Mode (NIA2/NEA2) and Registration Accept. It returns the
// AMF's NAS security context and the RAN-UE-NGAP-ID.
func register(t *testing.T, hs *harness) (*libnas.SecurityContext, int64) {
	t.Helper()
	amf := hs.amf
	k, opc := h("8baf473f2f8fd09487cccbd7097c6862"), h("8e27b6af0e692e750f32667a3b14605d")
	snn := security.ServingNetworkName("208", "93")
	suci, _ := security.NullSUCI("208", "93", "0000", "0000000001")

	ium := recvAs[*ngapMessage.InitialUEMessage](t, amf)
	ran := ium.RANUENGAPID.Value
	if uli := ium.UserLocationInformation.UserLocationInformationNR; !bytes.Equal(uli.TAI.TAC.Value, []byte{0, 0, 7}) ||
		gnb.BitValue(uli.NRCGI.NRCellIdentity.Value) != 0x123<<12 {
		t.Fatalf("ULI %+v", uli)
	}
	nas, err := libnas.Decode(ium.NASPDU.Value, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := nas.(*nasMessage.RegistrationRequest)
	if !bytes.Equal(rr.MobileIdentity, suci) || rr.RegistrationTypeAndNgKSI>>4 != 7 || !bytes.Equal(rr.UESecurityCapability.Value[:2], []byte{0xe0, 0xe0}) {
		t.Fatalf("RegistrationRequest identity %x ngKSI %x seccap %x", rr.MobileIdentity, rr.RegistrationTypeAndNgKSI, rr.UESecurityCapability.Value)
	}

	// 5G AKA: the network side computes AUTN and XRES*.
	rand, sqn, amfField := bytes.Repeat([]byte{0x5a}, 16), h("000000000021"), h("8000")
	m, _ := security.NewMilenage(k, opc)
	macA, _ := m.F1(rand, sqn, amfField)
	res, ck, ik, ak := m.F2345(rand)
	sqnXorAK := make([]byte, 6)
	for i := range sqnXorAK {
		sqnXorAK[i] = sqn[i] ^ ak[i]
	}
	autn := append(append(append([]byte(nil), sqnXorAK...), amfField...), macA...)
	authReq, err := libnas.Encode(&nasMessage.AuthenticationRequest{NgKSI: 1, ABBA: []byte{0, 0},
		HasAuthenticationParameterRAND: true, AuthenticationParameterRAND: nasType.AuthenticationParameterRAND{Value: rand},
		HasAuthenticationParameterAUTN: true, AuthenticationParameterAUTN: nasType.AuthenticationParameterAUTN{Value: autn}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	ids := func(nas []byte) *ngapMessage.DownlinkNASTransport {
		return &ngapMessage.DownlinkNASTransport{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran}, NASPDU: ngapType.NASPDU{Value: nas}}
	}
	amf.send(t, ids(authReq))

	ul := recvAs[*ngapMessage.UplinkNASTransport](t, amf)
	if ul.AMFUENGAPID.Value != 100 || ul.RANUENGAPID.Value != ran {
		t.Fatalf("IDs amf %d ran %d", ul.AMFUENGAPID.Value, ul.RANUENGAPID.Value)
	}
	nas, err = libnas.Decode(ul.NASPDU.Value, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := nas.(*nasMessage.AuthenticationResponse).AuthenticationResponseParameter.Value; !bytes.Equal(got, security.ResStar(ck, ik, snn, rand, res)) {
		t.Fatalf("RES* %x", got)
	}

	// Security Mode Command with NIA2 / NEA2, protected with the new keys (type 3).
	kamf := security.Kamf(security.Kseaf(security.Kausf(ck, ik, snn, sqnXorAK), snn), "208930000000001", []byte{0, 0})
	enc, integ := security.NASKeys(kamf, 2, 2)
	amfSec := &libnas.SecurityContext{IntegrityAlgorithm: 2, CipheringAlgorithm: 2, KnasInt: integ, KnasEnc: enc}
	smc, err := libnas.Encode(&nasMessage.SecurityModeCommand{SelectedNASAlgorithms: 0x22, NgKSI: 1,
		ReplayedUESecurityCapabilities: rr.UESecurityCapability.Value,
		HasIMEISVRequest:               true, IMEISVRequest: nasType.IMEISVRequest{Value: 1}},
		&libnas.EncodeContext{Security: amfSec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedWithNewContext, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	amf.send(t, ids(smc))

	ul = recvAs[*ngapMessage.UplinkNASTransport](t, amf)
	if ul.NASPDU.Value[1] != 4 {
		t.Fatalf("SecurityModeComplete header type %d", ul.NASPDU.Value[1])
	}
	nas, err = libnas.Decode(ul.NASPDU.Value, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink})
	if err != nil {
		t.Fatalf("SecurityModeComplete: %v", err)
	}
	smcc := nas.(*nasMessage.SecurityModeComplete)
	if !smcc.HasIMEISV || !bytes.Equal(smcc.IMEISV.Value, h("4573806121856151f1")) || !smcc.HasNASMessageContainer {
		t.Fatalf("SecurityModeComplete %+v", smcc)
	}
	inner, err := libnas.Decode(smcc.NASMessageContainer.Value, nil)
	if err != nil || !bytes.Equal(inner.(*nasMessage.RegistrationRequest).MobileIdentity, suci) {
		t.Fatalf("NAS message container: %v", err)
	}

	// Registration Accept, ciphered (type 2), inside InitialContextSetupRequest.
	guti := h("f202f839cafe0000000001")
	acc, err := libnas.Encode(&nasMessage.RegistrationAccept{RegistrationResult: []byte{0x01}, HasGUTI5GS: true, GUTI5GS: nasType.GUTI5GS{Value: guti}},
		&libnas.EncodeContext{Security: amfSec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	amf.send(t, &ngapMessage.InitialContextSetupRequest{
		AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran},
		GUAMI:        guami(),
		AllowedNSSAI: ngapType.AllowedNSSAI{List: []ngapType.AllowedNSSAIItem{{SNSSAI: snssai()}}},
		UESecurityCapabilities: ngapType.UESecurityCapabilities{
			NRencryptionAlgorithms:             ngapType.NRencryptionAlgorithms{Value: gnb.BitString(0xe000, 16)},
			NRintegrityProtectionAlgorithms:    ngapType.NRintegrityProtectionAlgorithms{Value: gnb.BitString(0xe000, 16)},
			EUTRAencryptionAlgorithms:          ngapType.EUTRAencryptionAlgorithms{Value: gnb.BitString(0, 16)},
			EUTRAintegrityProtectionAlgorithms: ngapType.EUTRAintegrityProtectionAlgorithms{Value: gnb.BitString(0, 16)},
		},
		SecurityKey: ngapType.SecurityKey{Value: ngapType.BitString{Bytes: security.KgNB(kamf, 0), BitLength: 256}},
		HasNASPDU:   true, NASPDU: ngapType.NASPDU{Value: acc},
	})

	icsr := recvAs[*ngapMessage.InitialContextSetupResponse](t, amf)
	if icsr.AMFUENGAPID.Value != 100 || icsr.RANUENGAPID.Value != ran {
		t.Fatalf("InitialContextSetupResponse IDs %+v", icsr)
	}
	ul = recvAs[*ngapMessage.UplinkNASTransport](t, amf)
	if nas, err = libnas.Decode(ul.NASPDU.Value, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink}); err != nil {
		t.Fatalf("RegistrationComplete: %v", err)
	}
	if _, ok := nas.(*nasMessage.RegistrationComplete); !ok {
		t.Fatalf("got %T", nas)
	}
	return amfSec, ran
}

func TestCommandFanOut(t *testing.T) {
	hs := start(t, true, 3, "node { entry { send(NGSetupRequest) } }\nstart { }")
	hs.linkup(t)
	if _, err := hs.eng.Exec("send(InitialUEMeesageRegistrationRequest)"); err != nil {
		t.Fatal(err)
	}
	seen := map[string]int64{}
	for i := 0; i < 3; i++ {
		ium := recvAs[*ngapMessage.InitialUEMessage](t, hs.amf)
		nas, _ := libnas.Decode(ium.NASPDU.Value, nil)
		seen[hex.EncodeToString(nas.(*nasMessage.RegistrationRequest).MobileIdentity)] = ium.RANUENGAPID.Value
	}
	for _, msin := range []string{"0000000001", "0000000002", "0000000003"} {
		suci, _ := security.NullSUCI("208", "93", "0000", msin)
		if _, ok := seen[hex.EncodeToString(suci)]; !ok {
			t.Fatalf("no registration from MSIN %s: %v", msin, seen)
		}
	}
	hs.wait(t, "summary", func(e models.Event) bool {
		return e.Kind == models.KindCommand && e.Text == "send(InitialUEMeesageRegistrationRequest): 3 succeeded, 0 failed"
	})

	if _, err := hs.eng.Exec("target(208930000000002).send(InitialUEMeesageRegistrationRequest)"); err != nil {
		t.Fatal(err)
	}
	ium := recvAs[*ngapMessage.InitialUEMessage](t, hs.amf)
	nas, _ := libnas.Decode(ium.NASPDU.Value, nil)
	if suci, _ := security.NullSUCI("208", "93", "0000", "0000000002"); !bytes.Equal(nas.(*nasMessage.RegistrationRequest).MobileIdentity, suci) {
		t.Fatal("target sent from the wrong subscriber")
	}

	// Without an AMF-UE-NGAP-ID nobody can answer: three errors and a summary.
	if _, err := hs.eng.Exec("send(UpLinkNASTransportAuthenticationResponse)"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "failures", func(e models.Event) bool {
		return e.Kind == models.KindCommand && strings.HasSuffix(e.Text, "0 succeeded, 3 failed")
	})
	if _, err := hs.eng.Exec("target(208930000000003).send(UpLinkNASTransportAuthenticationResponse)"); err == nil ||
		!strings.Contains(err.Error(), "AMF-UE-NGAP-ID") {
		t.Fatalf("single target error: %v", err)
	}
}

func TestBusHistory(t *testing.T) {
	b := engine.NewBus(3)
	for i := 0; i < 5; i++ {
		b.Publish(models.Event{Text: fmt.Sprint(i)})
	}
	got := b.History(0)
	if len(got) != 3 || got[0].Text != "2" || got[2].Seq != 5 {
		t.Fatalf("history %+v", got)
	}
	if len(b.History(4)) != 1 {
		t.Fatal("since")
	}
}

// Commands may run while the engine starts: Start sets the context that
// wait and the traffic runs read (the race detector checks this).
func TestStartWhileCommandsRun(t *testing.T) {
	eng, _, _ := buildEngine(t, true, 1, sessionScenario, 2152, 2152)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		_, err := eng.Exec("wait(0.05s)")
		done <- err
	}()
	eng.Start(ctx)
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	// A wait after the start ends when the engine does.
	go func() {
		_, err := eng.Exec("wait(10s)")
		done <- err
	}()
	time.Sleep(50 * time.Millisecond)
	cancel()
	eng.Wait()
	if err := <-done; err == nil || !strings.Contains(err.Error(), "interrupted") {
		t.Fatalf("wait after cancel: %v", err)
	}
}
