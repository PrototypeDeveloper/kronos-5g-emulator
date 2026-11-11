package engine_test

import (
	"bytes"
	"encoding/binary"
	"net"
	"strings"
	"testing"
	"time"

	libgtp "libgtp"
	"libgtp/gtpMessage"
	"libgtp/gtpType"
	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
)

const sessionScenario = `
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
		move(sessionestablishment)
	}
}
sessionestablishment {
	entry {
		send(UplinkNASTransportPDUSessionEstablishmentRequest)
	}
	receive(PDUSessionResourceSetupRequest:PDUSessionEstablishmentAccept) {
		send(PDUSessionResourceSetupResponse)
	}
}
`

func freeUDPPort(t *testing.T) int {
	t.Helper()
	c, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	return c.LocalAddr().(*net.UDPAddr).Port
}

const ulTEID = 0x1001

func setupTransfer(t *testing.T) []byte {
	t.Helper()
	reject := ngapType.Criticality{Value: ngapType.CriticalityPresentReject}
	b, err := ngapcodec.EncodeParameter("PDUSessionResourceSetupRequestTransfer", ngapType.PDUSessionResourceSetupRequestTransfer{
		ProtocolIEs: []ngapType.PDUSessionResourceSetupRequestTransferIEsEntry{
			{ID: ngapType.ProtocolIEID{Value: 139}, Criticality: reject,
				Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentULNGUUPTNLInformation,
				ULNGUUPTNLInformation: ngapType.UPTransportLayerInformation{Present: ngapType.UPTransportLayerInformationPresentGTPTunnel,
					GTPTunnel: ngapType.GTPTunnel{
						TransportLayerAddress: ngapType.TransportLayerAddress{Value: ngapType.BitString{Bytes: []byte{127, 0, 0, 1}, BitLength: 32}},
						GTPTEID:               ngapType.GTPTEID{Value: binary.BigEndian.AppendUint32(nil, ulTEID)},
					}}},
			{ID: ngapType.ProtocolIEID{Value: 134}, Criticality: reject,
				Present:        ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentPDUSessionType,
				PDUSessionType: ngapType.PDUSessionType{Value: ngapType.PDUSessionTypePresentIpv4}},
			{ID: ngapType.ProtocolIEID{Value: 136}, Criticality: reject,
				Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentQosFlowSetupRequestList,
				QosFlowSetupRequestList: ngapType.QosFlowSetupRequestList{List: []ngapType.QosFlowSetupRequestItem{{
					QosFlowIdentifier: ngapType.QosFlowIdentifier{Value: 1},
					QosFlowLevelQosParameters: ngapType.QosFlowLevelQosParameters{
						QosCharacteristics: ngapType.QosCharacteristics{Present: ngapType.QosCharacteristicsPresentNonDynamic5QI,
							NonDynamic5QI: ngapType.NonDynamic5QIDescriptor{FiveQI: ngapType.FiveQI{Value: 9}}},
						AllocationAndRetentionPriority: ngapType.AllocationAndRetentionPriority{
							PriorityLevelARP:        ngapType.PriorityLevelARP{Value: 8},
							PreEmptionCapability:    ngapType.PreEmptionCapability{Value: ngapType.PreEmptionCapabilityPresentShallNotTriggerPreEmption},
							PreEmptionVulnerability: ngapType.PreEmptionVulnerability{Value: ngapType.PreEmptionVulnerabilityPresentNotPreEmptable},
						},
					},
				}}}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func inetChecksum(b []byte) uint16 {
	var sum uint32
	for i := 0; i+1 < len(b); i += 2 {
		sum += uint32(binary.BigEndian.Uint16(b[i:]))
	}
	if len(b)%2 == 1 {
		sum += uint32(b[len(b)-1]) << 8
	}
	for sum>>16 != 0 {
		sum = sum&0xffff + sum>>16
	}
	return ^uint16(sum)
}

// stubUPF answers ICMP echo requests that arrive through the tunnel and
// reports Echo Responses on echoes.
func stubUPF(t *testing.T, upf *net.UDPConn, dlTEID uint32, gnb *net.UDPAddr, echoes chan<- struct{}) {
	buf := make([]byte, 65536)
	for {
		n, err := upf.Read(buf)
		if err != nil {
			return
		}
		msg, err := libgtp.Decode(buf[:n])
		if err != nil {
			t.Errorf("UPF: %v", err)
			return
		}
		switch m := msg.(type) {
		case *gtpMessage.EchoResponse:
			echoes <- struct{}{}
		case *gtpMessage.GPDU:
			if m.Header.TEID != ulTEID || m.PDUSessionContainer.PDUType != gtpType.PDUSessionInformationUL || m.PDUSessionContainer.QFI != 1 {
				t.Errorf("UPF: TEID %#x PDU type %d QFI %d", m.Header.TEID, m.PDUSessionContainer.PDUType, m.PDUSessionContainer.QFI)
				return
			}
			ip := append([]byte(nil), m.TPDU...)
			if !bytes.Equal(ip[12:16], []byte{10, 60, 0, 7}) || !bytes.Equal(ip[16:20], []byte{10, 10, 10, 20}) || ip[20] != 8 {
				t.Errorf("UPF: unexpected packet %x", ip[:24])
				return
			}
			// Turn the echo request around.
			copy(ip[12:16], m.TPDU[16:20])
			copy(ip[16:20], m.TPDU[12:16])
			ip[20] = 0
			binary.BigEndian.PutUint16(ip[22:], 0)
			binary.BigEndian.PutUint16(ip[22:], inetChecksum(ip[20:]))
			binary.BigEndian.PutUint16(ip[10:], 0)
			binary.BigEndian.PutUint16(ip[10:], inetChecksum(ip[:20]))
			g, _ := gtpMessage.NewGPDUWithPDUSessionContainer(dlTEID, ip, gtpType.PDUSessionInformationDL, 1)
			b, _ := g.Encode()
			upf.WriteToUDP(b, gnb)
		}
	}
}

// establish answers the PDU Session Establishment Request that the
// sessionestablishment entry sends (session 1, UE address 10.60.0.7) and
// returns the DL TEID the gNB allocated.
func establish(t *testing.T, hs *harness, amfSec *libnas.SecurityContext, ran int64) uint32 {
	t.Helper()

	ul := recvAs[*ngapMessage.UplinkNASTransport](t, hs.amf)
	nas, err := libnas.Decode(ul.NASPDU.Value, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink})
	if err != nil {
		t.Fatalf("PDU session establishment request: %v", err)
	}
	tr := nas.(*nasMessage.ULNASTransport)
	if !tr.HasPDUSessionID || tr.PDUSessionID.Value != 1 || !tr.HasDNN || string(tr.DNN.Value) != "\x08internet" {
		t.Fatalf("ULNASTransport %+v", tr)
	}
	inner, err := libnas.Decode(tr.PayloadContainer, nil)
	if err != nil {
		t.Fatal(err)
	}
	req := inner.(*nasMessage.PDUSessionEstablishmentRequest)
	if req.PDUSessionID != 1 || req.ProcedureTransactionID == 0 {
		t.Fatalf("PDUSessionEstablishmentRequest %+v", req)
	}

	accept, err := libnas.Encode(&nasMessage.PDUSessionEstablishmentAccept{PDUSessionID: 1, ProcedureTransactionID: req.ProcedureTransactionID,
		SelectedSSCModeAndPDUSessionType: 0x11, AuthorizedQoSRules: h("0100063101010109"), SessionAMBR: h("060001060001"),
		HasPDUAddress: true, PDUAddress: nasType.PDUAddress{Value: []byte{1, 10, 60, 0, 7}}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	dl, err := libnas.Encode(&nasMessage.DLNASTransport{PayloadContainerType: 1, PayloadContainer: accept,
		HasPDUSessionID: true, PDUSessionID: nasType.PDUSessionID{Value: 1}},
		&libnas.EncodeContext{Security: amfSec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	hs.amf.send(t, &ngapMessage.PDUSessionResourceSetupRequest{
		AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran},
		PDUSessionResourceSetupListSUReq: ngapType.PDUSessionResourceSetupListSUReq{List: []ngapType.PDUSessionResourceSetupItemSUReq{{
			PDUSessionID:        ngapType.PDUSessionID{Value: 1},
			HasPDUSessionNASPDU: true, PDUSessionNASPDU: ngapType.NASPDU{Value: dl},
			SNSSAI:                                 snssai(),
			PDUSessionResourceSetupRequestTransfer: setupTransfer(t),
		}}},
	})

	resp := recvAs[*ngapMessage.PDUSessionResourceSetupResponse](t, hs.amf)
	items := resp.PDUSessionResourceSetupListSURes.List
	if !resp.HasPDUSessionResourceSetupListSURes || len(items) != 1 || items[0].PDUSessionID.Value != 1 {
		t.Fatalf("setup response %+v", resp)
	}
	var rt ngapType.PDUSessionResourceSetupResponseTransfer
	if err := ngapcodec.DecodeParameter("PDUSessionResourceSetupResponseTransfer", items[0].PDUSessionResourceSetupResponseTransfer, &rt); ngapcodec.IsFatal(err) {
		t.Fatal(err)
	}
	g := rt.DLQosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel
	flows := rt.DLQosFlowPerTNLInformation.AssociatedQosFlowList.List
	if !bytes.Equal(g.TransportLayerAddress.Value.Bytes, []byte{127, 0, 0, 1}) || len(g.GTPTEID.Value) != 4 || len(flows) != 1 || flows[0].QosFlowIdentifier.Value != 1 {
		t.Fatalf("response transfer %+v", rt)
	}
	dlTEID := binary.BigEndian.Uint32(g.GTPTEID.Value)
	hs.wait(t, "session up", func(e models.Event) bool {
		return strings.Contains(e.Text, "PDU session 1 established: UE IP 10.60.0.7")
	})
	return dlTEID
}

// TestPDUSessionAndPing establishes a PDU session after registration (the
// sessionestablishment entry sends the request), checks the N3 information
// on both sides and pings through a stub UPF.
func TestPDUSessionAndPing(t *testing.T) {
	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	gtpuPort := freeUDPPort(t)
	hs := startN3(t, true, 1, sessionScenario, gtpuPort, upf.LocalAddr().(*net.UDPAddr).Port)
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	dlTEID := establish(t, hs, amfSec, ran)

	echoes := make(chan struct{}, 1)
	gnb := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: gtpuPort}
	go stubUPF(t, upf, dlTEID, gnb, echoes)
	if _, err := hs.eng.Exec("target(208930000000001).ping(10.10.10.20, 2)"); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"seq=1 reply", "seq=2 reply", "2 transmitted, 2 received, 0% loss"} {
		hs.wait(t, want, func(e models.Event) bool { return strings.Contains(e.Text, "ping 10.10.10.20: "+want) })
	}

	// The gNB answers GTP-U Echo Request.
	b, _ := gtpMessage.NewEchoRequest(7).Encode()
	upf.WriteToUDP(b, gnb)
	select {
	case <-echoes:
	case <-time.After(3 * time.Second):
		t.Fatal("no GTP-U Echo Response")
	}

	// ping needs a PDU session; the scenario check accepts it, the subscriber
	// reports the missing session only when it runs.
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "reset", func(e models.Event) bool { return strings.Contains(e.Text, "back to the initial state") })
	time.Sleep(50 * time.Millisecond) // let the subscriber process the reset
	if _, err := hs.eng.Exec("target(208930000000001).ping(10.10.10.20)"); err == nil || !strings.Contains(err.Error(), "no usable PDU session") {
		t.Fatalf("ping without a session: %v", err)
	}
}

const releaseScenario = sessionScenario + `
any {
	receive(PDUSessionResourceReleaseCommand:PDUSessionReleaseCommand) {
		send(UplinkNASTransportPDUSessionReleaseComplete)
		send(PDUSessionResourceReleaseResponse)
	}
	receive(UEContextReleaseCommand) {
		send(UEContextReleaseComplete)
	}
}
`

// TestReleases runs the release procedures: a UE requested PDU session
// release, establishing the session again, and a gNB requested UE context
// release.
func TestReleases(t *testing.T) {
	hs := startN3(t, true, 1, releaseScenario, freeUDPPort(t), freeUDPPort(t))
	ft := useFakeTUN(hs) // a TUN device follows the session
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	establish(t, hs, amfSec, ran)
	const imsi = "target(208930000000001)."
	tunEvent := func(what string) {
		t.Helper()
		hs.wait(t, "TUN "+what, func(e models.Event) bool { return strings.HasPrefix(e.Text, "TUN kronos-ue1 "+what) })
	}
	tunEvent("up")
	upNAS := func() any {
		t.Helper()
		ul := recvAs[*ngapMessage.UplinkNASTransport](t, hs.amf)
		nas, err := libnas.Decode(ul.NASPDU.Value, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink})
		if err != nil {
			t.Fatal(err)
		}
		inner, err := libnas.Decode(nas.(*nasMessage.ULNASTransport).PayloadContainer, nil)
		if err != nil {
			t.Fatal(err)
		}
		return inner
	}

	// UE requested PDU session release.
	if _, err := hs.eng.Exec(imsi + "send(UplinkNASTransportPDUSessionReleaseRequest)"); err != nil {
		t.Fatal(err)
	}
	req := upNAS().(*nasMessage.PDUSessionReleaseRequest)
	if req.PDUSessionID != 1 || req.ProcedureTransactionID == 0 {
		t.Fatalf("PDUSessionReleaseRequest %+v", req)
	}
	cmd, _ := libnas.Encode(&nasMessage.PDUSessionReleaseCommand{PDUSessionID: 1, ProcedureTransactionID: req.ProcedureTransactionID, Cause: 36}, nil)
	dl, err := libnas.Encode(&nasMessage.DLNASTransport{PayloadContainerType: 1, PayloadContainer: cmd, HasPDUSessionID: true, PDUSessionID: nasType.PDUSessionID{Value: 1}},
		&libnas.EncodeContext{Security: amfSec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	transfer, err := ngapcodec.EncodeParameter("PDUSessionResourceReleaseCommandTransfer",
		ngapType.PDUSessionResourceReleaseCommandTransfer{Cause: ngapType.Cause{Present: ngapType.CausePresentNas, Nas: ngapType.CauseNas{Value: 0}}})
	if err != nil {
		t.Fatal(err)
	}
	hs.amf.send(t, &ngapMessage.PDUSessionResourceReleaseCommand{
		AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran},
		HasNASPDU: true, NASPDU: ngapType.NASPDU{Value: dl},
		PDUSessionResourceToReleaseListRelCmd: ngapType.PDUSessionResourceToReleaseListRelCmd{List: []ngapType.PDUSessionResourceToReleaseItemRelCmd{{
			PDUSessionID: ngapType.PDUSessionID{Value: 1}, PDUSessionResourceReleaseCommandTransfer: transfer}}},
	})
	cpl := upNAS().(*nasMessage.PDUSessionReleaseComplete)
	if cpl.PDUSessionID != 1 || cpl.ProcedureTransactionID != req.ProcedureTransactionID {
		t.Fatalf("PDUSessionReleaseComplete %+v", cpl)
	}
	rr := recvAs[*ngapMessage.PDUSessionResourceReleaseResponse](t, hs.amf)
	if l := rr.PDUSessionResourceReleasedListRelRes.List; len(l) != 1 || l[0].PDUSessionID.Value != 1 || rr.RANUENGAPID.Value != ran {
		t.Fatalf("PDUSessionResourceReleaseResponse %+v", rr)
	}
	if _, err := hs.eng.Exec(imsi + "ping(10.10.10.20)"); err == nil || !strings.Contains(err.Error(), "no usable PDU session") {
		t.Fatalf("ping after the release: %v", err)
	}
	tunEvent("removed")
	if !ft.last().isClosed() {
		t.Fatal("the TUN device is open after the session release")
	}

	// move re-enters the state, so its entry establishes the session again.
	if _, err := hs.eng.Exec(imsi + "move(sessionestablishment)"); err != nil {
		t.Fatal(err)
	}
	establish(t, hs, amfSec, ran)
	tunEvent("up")
	if ft.count() != 2 {
		t.Fatalf("%d TUN devices opened", ft.count())
	}

	// gNB requested UE context release, listing the established session.
	if _, err := hs.eng.Exec(imsi + "send(AN_UEContextReleaseRequest)"); err != nil {
		t.Fatal(err)
	}
	crq := recvAs[*ngapMessage.UEContextReleaseRequest](t, hs.amf)
	if l := crq.PDUSessionResourceListCxtRelReq.List; crq.AMFUENGAPID.Value != 100 || crq.RANUENGAPID.Value != ran ||
		!crq.HasPDUSessionResourceListCxtRelReq || len(l) != 1 || l[0].PDUSessionID.Value != 1 {
		t.Fatalf("UEContextReleaseRequest %+v", crq)
	}
	hs.amf.send(t, &ngapMessage.UEContextReleaseCommand{
		UENGAPIDs: ngapType.UENGAPIDs{Present: ngapType.UENGAPIDsPresentUENGAPIDPair,
			UENGAPIDPair: ngapType.UENGAPIDPair{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran}}},
		Cause: ngapType.Cause{Present: ngapType.CausePresentRadioNetwork, RadioNetwork: ngapType.CauseRadioNetwork{Value: 20}},
	})
	ccpl := recvAs[*ngapMessage.UEContextReleaseComplete](t, hs.amf)
	if ccpl.AMFUENGAPID.Value != 100 || ccpl.RANUENGAPID.Value != ran ||
		!bytes.Equal(ccpl.UserLocationInformation.UserLocationInformationNR.TAI.PLMNIdentity.Value, []byte{0x02, 0xf8, 0x39}) {
		t.Fatalf("UEContextReleaseComplete %+v", ccpl)
	}
	hs.wait(t, "release", func(e models.Event) bool { return strings.Contains(e.Text, "UE context released") })

	// CM-IDLE: no NGAP IDs and no tunnel until the UE comes back.
	if _, err := hs.eng.Exec(imsi + "ping(10.10.10.20)"); err == nil || !strings.Contains(err.Error(), "no usable PDU session") {
		t.Fatalf("ping after the UE context release: %v", err)
	}
	if _, err := hs.eng.Exec(imsi + "send(UplinkNASTransportRegistrationComplete)"); err == nil || !strings.Contains(err.Error(), "no RAN-UE-NGAP-ID") {
		t.Fatalf("uplink NAS after the UE context release: %v", err)
	}
	// The session is kept while idle, and so is its TUN device.
	if ft.count() != 2 || ft.last().isClosed() {
		t.Fatalf("TUN after the UE context release: %d opened, closed %v", ft.count(), ft.last().isClosed())
	}
}

// TestEchoRequestToUE checks that an ICMP echo request from the DN to the
// UE's address is answered through the tunnel, and that one to another
// address is not.
func TestEchoRequestToUE(t *testing.T) {
	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	gtpuPort := freeUDPPort(t)
	hs := startN3(t, true, 1, sessionScenario, gtpuPort, upf.LocalAddr().(*net.UDPAddr).Port)
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	dlTEID := establish(t, hs, amfSec, ran)
	gnbAddr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: gtpuPort}

	request := func(dst net.IP, id, seq uint16) {
		t.Helper()
		icmp := make([]byte, 8+24)
		icmp[0] = 8
		binary.BigEndian.PutUint16(icmp[4:], id)
		binary.BigEndian.PutUint16(icmp[6:], seq)
		copy(icmp[8:], "kronos echo test payload")
		binary.BigEndian.PutUint16(icmp[2:], inetChecksum(icmp))
		ip := make([]byte, 20, 20+len(icmp))
		ip[0], ip[8], ip[9] = 0x45, 63, 1
		binary.BigEndian.PutUint16(ip[2:], uint16(20+len(icmp)))
		copy(ip[12:], []byte{10, 10, 10, 20})
		copy(ip[16:], dst.To4())
		binary.BigEndian.PutUint16(ip[10:], inetChecksum(ip))
		g, _ := gtpMessage.NewGPDUWithPDUSessionContainer(dlTEID, append(ip, icmp...), gtpType.PDUSessionInformationDL, 1)
		b, _ := g.Encode()
		if _, err := upf.WriteToUDP(b, gnbAddr); err != nil {
			t.Fatal(err)
		}
	}

	request(net.IPv4(10, 60, 0, 7), 0x77, 3)
	upf.SetReadDeadline(time.Now().Add(3 * time.Second))
	buf := make([]byte, 2048)
	n, err := upf.Read(buf)
	if err != nil {
		t.Fatalf("no echo reply: %v", err)
	}
	msg, err := libgtp.Decode(buf[:n])
	if err != nil {
		t.Fatal(err)
	}
	g, ok := msg.(*gtpMessage.GPDU)
	if !ok || g.Header.TEID != ulTEID || g.PDUSessionContainer.PDUType != gtpType.PDUSessionInformationUL || g.PDUSessionContainer.QFI != 1 {
		t.Fatalf("uplink %T %+v", msg, msg)
	}
	ip := g.TPDU
	icmp := ip[20:]
	if !bytes.Equal(ip[12:16], []byte{10, 60, 0, 7}) || !bytes.Equal(ip[16:20], []byte{10, 10, 10, 20}) || inetChecksum(ip[:20]) != 0 ||
		icmp[0] != 0 || binary.BigEndian.Uint16(icmp[4:]) != 0x77 || binary.BigEndian.Uint16(icmp[6:]) != 3 ||
		string(icmp[8:]) != "kronos echo test payload" || inetChecksum(icmp) != 0 {
		t.Fatalf("echo reply %x", ip)
	}
	hs.wait(t, "echo event", func(e models.Event) bool {
		return strings.Contains(e.Text, "answered an ICMP Echo Request (10.10.10.20 → 10.60.0.7, id=119 seq=3)")
	})

	// Not the UE's address: no answer.
	request(net.IPv4(10, 60, 0, 99), 0x78, 1)
	upf.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
	if n, err := upf.Read(buf); err == nil {
		t.Fatalf("unexpected uplink packet %x", buf[:n])
	}
}
