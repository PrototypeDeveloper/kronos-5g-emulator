package engine_test

import (
	"bytes"
	"net"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	nassec "libnas/security"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/models"
	"kronos/simulator/gnb"
)

const idleScenario = sessionScenario + `
any {
	receive(UEContextReleaseCommand) {
		send(UEContextReleaseComplete)
	}
	receive(InitialContextSetupRequest:ServiceAccept) {
		send(InitialContextSetupResponseSVR)
	}
	receive(Paging) {
		send(InitialUEMeesageServiceRequest)
	}
}
`

// The 5G-GUTI that register assigns: AMF region 0xca, AMF set ID 0x3f8,
// AMF pointer 0, 5G-TMSI 00000001.
var (
	testGUTI     = h("f202f839cafe0000000001")
	testSTMSINAS = h("f4fe0000000001")
)

func pagingFor(setID, pointer uint64, tmsi []byte) *ngapMessage.Paging {
	return &ngapMessage.Paging{
		UEPagingIdentity: ngapType.UEPagingIdentity{Present: ngapType.UEPagingIdentityPresentFiveGSTMSI,
			FiveGSTMSI: ngapType.FiveGSTMSI{
				AMFSetID:   ngapType.AMFSetID{Value: gnb.BitString(setID, 10)},
				AMFPointer: ngapType.AMFPointer{Value: gnb.BitString(pointer, 6)},
				FiveGTMSI:  ngapType.FiveGTMSI{Value: tmsi}}},
		TAIListForPaging: ngapType.TAIListForPaging{List: []ngapType.TAIListForPagingItem{{
			TAI: ngapType.TAI{PLMNIdentity: ngapType.PLMNIdentity{Value: []byte{0x02, 0xf8, 0x39}}, TAC: ngapType.TAC{Value: []byte{0, 0, 7}}}}}},
	}
}

// releaseUE is an AMF initiated UE context release (the UE goes to CM-IDLE).
func releaseUE(t *testing.T, hs *harness, amfID, ran int64) {
	t.Helper()
	hs.amf.send(t, &ngapMessage.UEContextReleaseCommand{
		UENGAPIDs: ngapType.UENGAPIDs{Present: ngapType.UENGAPIDsPresentUENGAPIDPair,
			UENGAPIDPair: ngapType.UENGAPIDPair{AMFUENGAPID: ngapType.AMFUENGAPID{Value: amfID}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran}}},
		Cause: ngapType.Cause{Present: ngapType.CausePresentNas, Nas: ngapType.CauseNas{Value: 0}},
	})
	cpl := recvAs[*ngapMessage.UEContextReleaseComplete](t, hs.amf)
	if cpl.AMFUENGAPID.Value != amfID || cpl.RANUENGAPID.Value != ran {
		t.Fatalf("UEContextReleaseComplete %+v", cpl)
	}
	hs.wait(t, "release", func(e models.Event) bool { return strings.Contains(e.Text, "UE context released") })
}

// initialNAS reads an InitialUEMessage sent from CM-IDLE under the security
// context: a new RAN-UE-NGAP-ID, the 5G-S-TMSI, and an integrity protected
// (not ciphered) NAS message. It returns the NAS message and its COUNT.
func initialNAS(t *testing.T, hs *harness, amfSec *libnas.SecurityContext, oldRAN int64) (*ngapMessage.InitialUEMessage, any, uint32) {
	t.Helper()
	ium := recvAs[*ngapMessage.InitialUEMessage](t, hs.amf)
	if ium.RANUENGAPID.Value == oldRAN {
		t.Fatalf("RAN-UE-NGAP-ID %d is not new", oldRAN)
	}
	s := ium.FiveGSTMSI
	if !ium.HasFiveGSTMSI || gnb.BitValue(s.AMFSetID.Value) != 0x3f8 || gnb.BitValue(s.AMFPointer.Value) != 0 || !bytes.Equal(s.FiveGTMSI.Value, []byte{0, 0, 0, 1}) {
		t.Fatalf("5G-S-TMSI %+v", s)
	}
	pdu := ium.NASPDU.Value
	if pdu[1] != byte(libnas.SecurityHeaderIntegrityProtected) {
		t.Fatalf("security header type %d", pdu[1])
	}
	count := uint32(pdu[6]) // the overflow counter is still 0
	nas, err := libnas.Decode(pdu, &libnas.DecodeContext{Security: amfSec, Direction: libnas.Uplink})
	if err != nil {
		t.Fatalf("initial NAS: %v", err)
	}
	return ium, nas, count
}

// openContainer deciphers a NAS message container with the COUNT of the
// message that carries it (TS 24.501 §4.4.6) and decodes it.
func openContainer(t *testing.T, amfSec *libnas.SecurityContext, count uint32, c []byte) any {
	t.Helper()
	plain, err := nassec.Cipher(amfSec.CipheringAlgorithm, amfSec.KnasEnc, count, 1, uint8(libnas.Uplink), c)
	if err != nil {
		t.Fatal(err)
	}
	inner, err := libnas.Decode(plain, nil)
	if err != nil {
		t.Fatalf("NAS message container %x: %v", plain, err)
	}
	return inner
}

// TestIdleModeProcedures runs the CM-IDLE procedures: Paging and the
// Service Request it triggers (the session's user plane comes back and ping
// works again), a mobility registration update and a switch-off
// de-registration, each from CM-IDLE under the NAS security context.
func TestIdleModeProcedures(t *testing.T) {
	upf, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer upf.Close()
	gtpuPort := freeUDPPort(t)
	hs := startN3(t, true, 1, idleScenario, gtpuPort, upf.LocalAddr().(*net.UDPAddr).Port)
	ft := useFakeTUN(hs)
	hs.linkup(t)
	amfSec, ran := register(t, hs)
	dlTEID := establish(t, hs, amfSec, ran)
	go stubUPF(t, upf, dlTEID, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: gtpuPort}, make(chan struct{}, 1))
	const imsi = "target(208930000000001)."
	hs.wait(t, "TUN up", func(e models.Event) bool { return strings.HasPrefix(e.Text, "TUN kronos-ue1 up") })

	// A Paging for a UE the gNB does not serve is reported, not answered.
	hs.amf.send(t, pagingFor(0x3f8, 0, []byte{0, 0, 0, 9}))
	hs.wait(t, "paging warning", func(e models.Event) bool {
		return e.Kind == models.KindWarn && strings.Contains(e.Text, "no subscriber for Paging (5G-S-TMSI 3f8-00-00000009)")
	})

	releaseUE(t, hs, 100, ran)

	// Paging → Service Request with the NAS message container ciphered.
	hs.amf.send(t, pagingFor(0x3f8, 0, []byte{0, 0, 0, 1}))
	ium, nas, count := initialNAS(t, hs, amfSec, ran)
	ran2 := ium.RANUENGAPID.Value
	sr, ok := nas.(*nasMessage.ServiceRequest)
	if !ok {
		t.Fatalf("got %T", nas)
	}
	// ngKSI 1 (from the network) in the low half octet, the service type of
	// the template (data) in the high one; the status IEs only in the container.
	if sr.ServiceTypeAndNgKSI != 0x11 || !bytes.Equal(sr.FiveGSTMSI, testSTMSINAS) || sr.HasPDUSessionStatus || !sr.HasNASMessageContainer {
		t.Fatalf("ServiceRequest %+v", sr)
	}
	if bytes.Contains(sr.NASMessageContainer, testSTMSINAS) {
		t.Fatal("the NAS message container is not ciphered")
	}
	in := openContainer(t, amfSec, count, sr.NASMessageContainer).(*nasMessage.ServiceRequest)
	if in.ServiceTypeAndNgKSI != 0x11 || !bytes.Equal(in.FiveGSTMSI, testSTMSINAS) || in.HasNASMessageContainer ||
		!bytes.Equal(in.PDUSessionStatus, []byte{0x02, 0}) || !bytes.Equal(in.UplinkDataStatus, []byte{0x02, 0}) {
		t.Fatalf("ServiceRequest in the container %+v", in)
	}

	// Service Accept with the session's user plane in InitialContextSetupRequest.
	acc, err := libnas.Encode(&nasMessage.ServiceAccept{},
		&libnas.EncodeContext{Security: amfSec, SecurityHeaderType: libnas.SecurityHeaderIntegrityProtectedAndCiphered, Direction: libnas.Downlink})
	if err != nil {
		t.Fatal(err)
	}
	hs.amf.send(t, &ngapMessage.InitialContextSetupRequest{
		AMFUENGAPID: ngapType.AMFUENGAPID{Value: 101}, RANUENGAPID: ngapType.RANUENGAPID{Value: ran2},
		GUAMI:        guami(),
		AllowedNSSAI: ngapType.AllowedNSSAI{List: []ngapType.AllowedNSSAIItem{{SNSSAI: snssai()}}},
		UESecurityCapabilities: ngapType.UESecurityCapabilities{
			NRencryptionAlgorithms:             ngapType.NRencryptionAlgorithms{Value: gnb.BitString(0xe000, 16)},
			NRintegrityProtectionAlgorithms:    ngapType.NRintegrityProtectionAlgorithms{Value: gnb.BitString(0xe000, 16)},
			EUTRAencryptionAlgorithms:          ngapType.EUTRAencryptionAlgorithms{Value: gnb.BitString(0, 16)},
			EUTRAintegrityProtectionAlgorithms: ngapType.EUTRAintegrityProtectionAlgorithms{Value: gnb.BitString(0, 16)},
		},
		SecurityKey:                          ngapType.SecurityKey{Value: gnb.BitString(0, 256)},
		HasPDUSessionResourceSetupListCxtReq: true,
		PDUSessionResourceSetupListCxtReq: ngapType.PDUSessionResourceSetupListCxtReq{List: []ngapType.PDUSessionResourceSetupItemCxtReq{{
			PDUSessionID: ngapType.PDUSessionID{Value: 1}, SNSSAI: snssai(), PDUSessionResourceSetupRequestTransfer: setupTransfer(t)}}},
		HasNASPDU: true, NASPDU: ngapType.NASPDU{Value: acc},
	})
	icsr := recvAs[*ngapMessage.InitialContextSetupResponse](t, hs.amf)
	items := icsr.PDUSessionResourceSetupListCxtRes.List
	if icsr.AMFUENGAPID.Value != 101 || icsr.RANUENGAPID.Value != ran2 || !icsr.HasPDUSessionResourceSetupListCxtRes || len(items) != 1 || items[0].PDUSessionID.Value != 1 {
		t.Fatalf("InitialContextSetupResponse %+v", icsr)
	}
	hs.wait(t, "service accept", func(e models.Event) bool { return strings.Contains(e.Text, "Service Accept") })
	// The downlink TEID stays the one of the establishment (stubUPF answers to it).
	if _, err := hs.eng.Exec(imsi + "ping(10.10.10.20, 1)"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "ping", func(e models.Event) bool { return strings.Contains(e.Text, "1 transmitted, 1 received, 0% loss") })

	// Mobility registration update from CM-IDLE: the complete message,
	// with the session status, goes ciphered in the container.
	releaseUE(t, hs, 101, ran2)
	if _, err := hs.eng.Exec(imsi + "send(InitialUEMeesageMobilityUpdate)"); err != nil {
		t.Fatal(err)
	}
	ium, nas, count = initialNAS(t, hs, amfSec, ran2)
	ran3 := ium.RANUENGAPID.Value
	rr := nas.(*nasMessage.RegistrationRequest)
	if rr.RegistrationTypeAndNgKSI != 0x12 || !bytes.Equal(rr.MobileIdentity, testGUTI) || !rr.HasNASMessageContainer {
		t.Fatalf("RegistrationRequest %+v", rr)
	}
	irr := openContainer(t, amfSec, count, rr.NASMessageContainer.Value).(*nasMessage.RegistrationRequest)
	if irr.RegistrationTypeAndNgKSI != 0x12 || !bytes.Equal(irr.MobileIdentity, testGUTI) || irr.HasNASMessageContainer ||
		!bytes.Equal(irr.PDUSessionStatus.Value, []byte{0x02, 0}) || !bytes.Equal(irr.UplinkDataStatus.Value, []byte{0x02, 0}) ||
		!bytes.Equal(irr.UESecurityCapability.Value[:2], []byte{0xe0, 0xe0}) {
		t.Fatalf("RegistrationRequest in the container %+v", irr)
	}

	// Switch-off de-registration from CM-IDLE: the AMF answers with the release only.
	releaseUE(t, hs, 102, ran3)
	if ft.count() != 1 || ft.last().isClosed() {
		t.Fatalf("TUN before the de-registration: %d opened, closed %v", ft.count(), ft.last().isClosed())
	}
	if _, err := hs.eng.Exec(imsi + "send(InitialUEMeesageDeregistrationRequest)"); err != nil {
		t.Fatal(err)
	}
	// The UE releases its PDU sessions locally, and the TUN device goes with them.
	hs.wait(t, "TUN removed", func(e models.Event) bool { return strings.HasPrefix(e.Text, "TUN kronos-ue1 removed") })
	if st, _ := hs.eng.UE("208930000000001"); len(st.Sessions) != 0 || st.RM != models.RMDeregistered || !ft.last().isClosed() {
		t.Fatalf("after the switch-off: %+v", st)
	}
	ium, nas, _ = initialNAS(t, hs, amfSec, ran3)
	dr := nas.(*nasMessage.DeregistrationRequestUE)
	if dr.DeregistrationTypeAndNgKSI != 0x19 || !bytes.Equal(dr.MobileIdentity, testGUTI) {
		t.Fatalf("DeregistrationRequestUE %+v", dr)
	}
	releaseUE(t, hs, 103, ium.RANUENGAPID.Value)
}
