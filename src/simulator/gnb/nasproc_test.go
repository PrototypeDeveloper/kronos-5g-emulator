package gnb

import (
	"bytes"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/common/security"
	"kronos/scenario"
)

// ulNAS builds a NAS message inside an UplinkNASTransport on u and returns
// the NAS-PDU as sent.
func ulNAS(t *testing.T, u *UE, m any) []byte {
	t.Helper()
	nb, err := libnas.Encode(m, nil)
	if err != nil {
		t.Fatal(err)
	}
	var ul ngapMessage.UplinkNASTransport
	ul.AMFUENGAPID.Value, ul.RANUENGAPID.Value, ul.NASPDU.Value = 100, 1, nb
	ul.UserLocationInformation.Present = ngapType.UserLocationInformationPresentUserLocationInformationNR
	nr := &ul.UserLocationInformation.UserLocationInformationNR
	nr.NRCGI.PLMNIdentity.Value, nr.TAI.PLMNIdentity.Value, nr.TAI.TAC.Value = []byte{0, 0xf1, 0x10}, []byte{0, 0xf1, 0x10}, []byte{0, 0, 1}
	nr.NRCGI.NRCellIdentity.Value = ngapType.BitString{Bytes: make([]byte, 5), BitLength: 36}
	b, err := libngap.Encode(&ul)
	if err != nil {
		t.Fatal(err)
	}
	built, err := u.Build(&scenario.Signal{Name: "x", Bytes: b, Mode: scenario.ModeAuto}, 0)
	if err != nil {
		t.Fatal(err)
	}
	out, _ := libngap.Decode(built.PDU)
	return out.(*ngapMessage.UplinkNASTransport).NASPDU.Value
}

// dlNAS hands u a downlink NAS message.
func dlNAS(t *testing.T, u *UE, m any) []string {
	t.Helper()
	nb, err := libnas.Encode(m, nil)
	if err != nil {
		t.Fatal(err)
	}
	var dl ngapMessage.DownlinkNASTransport
	dl.AMFUENGAPID.Value, dl.RANUENGAPID.Value, dl.NASPDU.Value = 100, 1, nb
	in := u.Receive(&dl)
	return append(in.Notes, in.Errors...)
}

func TestIdentityResponse(t *testing.T) {
	r := autoRole(t, `, imeisv: "4901542032375199"`)
	u := r.NewUE(0, "001010000000001").(*UE)
	u.ranID, u.hasRAN, u.amfID, u.hasAMF = 1, true, 100, true
	u.guti = h("f202f839cafe0000000001")
	suci, _ := security.NullSUCI("001", "01", r.ri, u.msin)
	for typ, want := range map[uint8][]byte{
		nasType.MobileIdentitySUCI: suci, nasType.MobileIdentity5GGUTI: u.guti, identityIMEI: r.imei, identityIMEISV: r.imeisv,
		identity5GSTMSI: h("f4fe0000000001"), // AMF set ID and pointer, 5G-TMSI of the 5G-GUTI
	} {
		dlNAS(t, u, &nasMessage.IdentityRequest{IdentityType: typ})
		nas, err := libnas.Decode(ulNAS(t, u, &nasMessage.IdentityResponse{MobileIdentity: []byte{0x01, 0x02, 0xf8, 0x39}}), nil)
		if err != nil {
			t.Fatal(err)
		}
		if got := nas.(*nasMessage.IdentityResponse).MobileIdentity; !bytes.Equal(got, want) {
			t.Errorf("identity type %d: %x, want %x", typ, got, want)
		}
	}
}

// h decodes hex for the tests.
func h(s string) []byte {
	b := make([]byte, len(s)/2)
	for i := range b {
		var v byte
		for _, c := range s[2*i : 2*i+2] {
			v <<= 4
			switch {
			case c >= '0' && c <= '9':
				v |= byte(c - '0')
			default:
				v |= byte(c-'a') + 10
			}
		}
		b[i] = v
	}
	return b
}

func TestAuthenticationFailureAUTS(t *testing.T) {
	r := autoRole(t, "")
	u := r.NewUE(0, "001010000000001").(*UE)
	u.ranID, u.hasRAN, u.amfID, u.hasAMF = 1, true, 100, true
	// The network's AuthenticationRequest: RAND and AUTN with SQN 0x21.
	m, _ := security.NewMilenage(r.k, r.opc)
	rand, sqn, amf := bytes.Repeat([]byte{0x5a}, 16), h("000000000021"), h("8000")
	macA, _ := m.F1(rand, sqn, amf)
	_, _, _, ak := m.F2345(rand)
	autn := make([]byte, 0, 16)
	for i := range sqn {
		autn = append(autn, sqn[i]^ak[i])
	}
	autn = append(append(autn, amf...), macA...)
	dlNAS(t, u, &nasMessage.AuthenticationRequest{NgKSI: 1, ABBA: []byte{0, 0},
		HasAuthenticationParameterRAND: true, AuthenticationParameterRAND: nasType.AuthenticationParameterRAND{Value: rand},
		HasAuthenticationParameterAUTN: true, AuthenticationParameterAUTN: nasType.AuthenticationParameterAUTN{Value: autn}})
	// The network reads SQN_MS out of the AUTS and checks MAC-S.
	nas, _ := libnas.Decode(ulNAS(t, u, &nasMessage.AuthenticationFailure{Cause: causeSynchFailure}), nil)
	af := nas.(*nasMessage.AuthenticationFailure)
	auts := af.AuthenticationFailureParameter.Value
	if !af.HasAuthenticationFailureParameter || len(auts) != 14 {
		t.Fatalf("AUTS %x", auts)
	}
	akStar := m.F5Star(rand)
	sqnMS := make([]byte, 6)
	for i := range sqnMS {
		sqnMS[i] = auts[i] ^ akStar[i]
	}
	_, macS := m.F1(rand, sqnMS, []byte{0, 0})
	if !bytes.Equal(sqnMS, sqn) || !bytes.Equal(macS, auts[6:]) {
		t.Fatalf("SQN_MS %x MAC-S %x (AUTS %x)", sqnMS, macS, auts)
	}
	// A MAC failure carries no parameter.
	nas, _ = libnas.Decode(ulNAS(t, u, &nasMessage.AuthenticationFailure{Cause: 20, HasAuthenticationFailureParameter: true,
		AuthenticationFailureParameter: nasType.AuthenticationFailureParameter{Value: auts}}), nil)
	if nas.(*nasMessage.AuthenticationFailure).HasAuthenticationFailureParameter {
		t.Fatal("MAC failure with a parameter")
	}
}

func TestNASProceduresOfTheUE(t *testing.T) {
	u := established(t)
	// SecurityModeReject goes in clear, and the context before the command comes back.
	old := &libnas.SecurityContext{IntegrityAlgorithm: 2}
	u.prevSec, u.sec = old, &libnas.SecurityContext{IntegrityAlgorithm: 1}
	if b := ulNAS(t, u, &nasMessage.SecurityModeReject{Cause: 24}); b[1]&shtMask != 0 || u.sec != old {
		t.Fatalf("SecurityModeReject %x (context %+v)", b, u.sec)
	}
	u.sec = nil
	// Status5GSM in ULNASTransport gets the session of ue.session.
	st, _ := libnas.Encode(&nasMessage.Status5GSM{PDUSessionID: 9, ProcedureTransactionID: 3, Cause: 98}, nil)
	conf, _ := u.r.sessionConf(9)
	nas, _ := libnas.Decode(ulNAS(t, u, &nasMessage.ULNASTransport{PayloadContainerType: payloadN1SMInfo, PayloadContainer: st}), nil)
	inner, _ := libnas.Decode(nas.(*nasMessage.ULNASTransport).PayloadContainer, nil)
	if s := inner.(*nasMessage.Status5GSM); s.PDUSessionID != conf.id || s.ProcedureTransactionID != 3 {
		t.Fatalf("Status5GSM %+v (session %d)", s, conf.id)
	}
	// A rejected establishment leaves no session behind.
	u.session(2)
	rej, _ := libnas.Encode(&nasMessage.PDUSessionEstablishmentReject{PDUSessionID: 2, ProcedureTransactionID: 1, Cause: 26}, nil)
	dlNAS(t, u, &nasMessage.DLNASTransport{PayloadContainerType: payloadN1SMInfo, PayloadContainer: rej})
	if u.sessions[2] != nil || u.sessions[1] == nil {
		t.Fatalf("sessions %v", u.sessions)
	}
	// AuthenticationReject: the 5G-GUTI and the security context are gone.
	u.registered, u.guti, u.sec = true, h("f202f839cafe0000000001"), old
	dlNAS(t, u, &nasMessage.AuthenticationReject{})
	if u.registered || u.guti != nil || u.sec != nil || u.ngKSI != noKey || len(u.sessions) != 0 {
		t.Fatalf("after AuthenticationReject %+v", u)
	}
}
