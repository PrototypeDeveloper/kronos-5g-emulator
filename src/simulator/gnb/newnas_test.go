package gnb

import (
	"bytes"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
)

// sm5GSM sends a 5GSM message up in a ULNASTransport and returns it as sent.
func sm5GSM(t *testing.T, u *UE, m any) any {
	t.Helper()
	b, err := libnas.Encode(m, nil)
	if err != nil {
		t.Fatal(err)
	}
	nas, err := libnas.Decode(ulNAS(t, u, &nasMessage.ULNASTransport{PayloadContainerType: payloadN1SMInfo, PayloadContainer: b}), nil)
	if err != nil {
		t.Fatal(err)
	}
	inner, err := libnas.Decode(nas.(*nasMessage.ULNASTransport).PayloadContainer, nil)
	if err != nil {
		t.Fatal(err)
	}
	return inner
}

// dl5GSM hands u a 5GSM message in a DLNASTransport.
func dl5GSM(t *testing.T, u *UE, m any) []string {
	t.Helper()
	b, err := libnas.Encode(m, nil)
	if err != nil {
		t.Fatal(err)
	}
	return dlNAS(t, u, &nasMessage.DLNASTransport{PayloadContainerType: payloadN1SMInfo, PayloadContainer: b})
}

// ulPlain sends a 5GMM message up and decodes it as sent.
func ulPlain(t *testing.T, u *UE, m any) any {
	t.Helper()
	out, err := libnas.Decode(ulNAS(t, u, m), nil)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func hasNote(notes []string, s string) bool {
	for _, n := range notes {
		if strings.Contains(n, s) {
			return true
		}
	}
	return false
}

// The 5GMM messages added to libnas: what the UE takes from the network
// and what the mapping writes into its answers.
func TestNew5GMMMessages(t *testing.T) {
	u := established(t)

	// Configuration Update Command: the new 5G-GUTI is taken.
	guti := h("f202f839cafe00000000aa")
	notes := dlNAS(t, u, &nasMessage.ConfigurationUpdateCommand{HasConfigurationUpdateIndication: true, ConfigurationUpdateIndication: nasType.ConfigurationUpdateIndication{Value: cuiACKBit},
		HasGUTI5GS: true, GUTI5GS: nasType.GUTI5GS{Value: guti}})
	if !bytes.Equal(u.guti, guti) || !hasNote(notes, "acknowledgement requested") {
		t.Fatalf("after ConfigurationUpdateCommand: guti %x, notes %q", u.guti, notes)
	}

	// Control Plane Service Request: ngKSI and the session status.
	cp := ulPlain(t, u, &nasMessage.ControlPlaneServiceRequest{ControlPlaneServiceTypeAndNgKSI: 0x01,
		HasPDUSessionStatus: true, PDUSessionStatus: nasType.PDUSessionStatus{Value: []byte{0, 0}},
		HasUplinkDataStatus: true, UplinkDataStatus: nasType.UplinkDataStatus{Value: []byte{0xff, 0xff}}}).(*nasMessage.ControlPlaneServiceRequest)
	if cp.ControlPlaneServiceTypeAndNgKSI != noKey<<ngKSIShift|0x01 || !bytes.Equal(cp.PDUSessionStatus.Value, []byte{0x02, 0}) || !bytes.Equal(cp.UplinkDataStatus.Value, []byte{0x02, 0}) {
		t.Fatalf("ControlPlaneServiceRequest %+v", cp)
	}

	// Notification Response: the session status.
	nr := ulPlain(t, u, &nasMessage.NotificationResponse{HasPDUSessionStatus: true, PDUSessionStatus: nasType.PDUSessionStatus{Value: []byte{0, 0}}}).(*nasMessage.NotificationResponse)
	if !bytes.Equal(nr.PDUSessionStatus.Value, []byte{0x02, 0}) {
		t.Fatalf("NotificationResponse %+v", nr)
	}

	// NSSAA: the complete names the slice of the command.
	dlNAS(t, u, &nasMessage.NetworkSliceSpecificAuthenticationCommand{SNSSAI: []byte{1, 0, 0, 9}, EAPMessage: make([]byte, 4)})
	nc := ulPlain(t, u, &nasMessage.NetworkSliceSpecificAuthenticationComplete{SNSSAI: []byte{1}, EAPMessage: make([]byte, 4)}).(*nasMessage.NetworkSliceSpecificAuthenticationComplete)
	if !bytes.Equal(nc.SNSSAI, []byte{1, 0, 0, 9}) {
		t.Fatalf("NSSAA complete S-NSSAI %x", nc.SNSSAI)
	}

	// Service Reject #10: de-registered, the 5G-GUTI stays.
	u.registered = true
	dlNAS(t, u, &nasMessage.ServiceReject{Cause: causeImplicitlyDeregistered})
	if u.registered || len(u.sessions) != 0 || u.guti == nil {
		t.Fatalf("after Service Reject #10: %+v", u)
	}
	// #9: the 5G-GUTI and the security context go too.
	u.sec, u.ngKSI = &libnas.SecurityContext{}, 1
	dlNAS(t, u, &nasMessage.ServiceReject{Cause: causeIdentityCannotBeDerived})
	if u.guti != nil || u.sec != nil || u.ngKSI != noKey {
		t.Fatalf("after Service Reject #9: %+v", u)
	}

	// Network-initiated de-registration: the sessions are released locally.
	u = established(t)
	notes = dlNAS(t, u, &nasMessage.DeregistrationRequestNW{DeregistrationType: reRegistrationBit | 1, HasCause5GMM: true, Cause5GMM: nasType.Cause5GMM{Value: 22}})
	if len(u.sessions) != 0 || !hasNote(notes, "5GMM cause 22, re-registration required") {
		t.Fatalf("DeregistrationRequestNW: sessions %v, notes %q", u.sessions, notes)
	}
}

// The 5GSM messages added to libnas: the PDU session ID and PTI.
func TestNew5GSMMessages(t *testing.T) {
	u := established(t)
	// A UE-requested modification gets a new PTI.
	mr := sm5GSM(t, u, &nasMessage.PDUSessionModificationRequest{PDUSessionID: 1, ProcedureTransactionID: 200}).(*nasMessage.PDUSessionModificationRequest)
	if mr.PDUSessionID != 1 || mr.ProcedureTransactionID == 200 || mr.ProcedureTransactionID == 0 || u.sessions[1].pti != mr.ProcedureTransactionID {
		t.Fatalf("PDUSessionModificationRequest %+v", mr)
	}
	// The complete and the command reject answer with the PTI of the command.
	notes := dl5GSM(t, u, &nasMessage.PDUSessionModificationCommand{PDUSessionID: 1, ProcedureTransactionID: mr.ProcedureTransactionID})
	if !hasNote(notes, "PDU session 1 modification commanded") {
		t.Fatalf("notes %q", notes)
	}
	mc := sm5GSM(t, u, &nasMessage.PDUSessionModificationComplete{PDUSessionID: 1, ProcedureTransactionID: 99}).(*nasMessage.PDUSessionModificationComplete)
	if mc.ProcedureTransactionID != mr.ProcedureTransactionID {
		t.Fatalf("PDUSessionModificationComplete %+v", mc)
	}
	dl5GSM(t, u, &nasMessage.PDUSessionModificationCommand{PDUSessionID: 1}) // network-initiated: PTI 0
	cr := sm5GSM(t, u, &nasMessage.PDUSessionModificationCommandReject{PDUSessionID: 1, ProcedureTransactionID: 99, Cause: 43}).(*nasMessage.PDUSessionModificationCommandReject)
	if cr.ProcedureTransactionID != 0 || cr.Cause != 43 {
		t.Fatalf("PDUSessionModificationCommandReject %+v", cr)
	}
	// PDU session authentication.
	dl5GSM(t, u, &nasMessage.PDUSessionAuthenticationCommand{PDUSessionID: 1, ProcedureTransactionID: 9, EAPMessage: make([]byte, 4)})
	ac := sm5GSM(t, u, &nasMessage.PDUSessionAuthenticationComplete{PDUSessionID: 1, ProcedureTransactionID: 1, EAPMessage: make([]byte, 4)}).(*nasMessage.PDUSessionAuthenticationComplete)
	if ac.ProcedureTransactionID != 9 {
		t.Fatalf("PDUSessionAuthenticationComplete %+v", ac)
	}
	if notes := dl5GSM(t, u, &nasMessage.PDUSessionModificationReject{PDUSessionID: 1, ProcedureTransactionID: 3, Cause: 26}); !hasNote(notes, "PDU session 1 modification rejected (cause 26)") {
		t.Fatalf("notes %q", notes)
	}
}

// The signal editor marks what the mapping writes into the new messages.
func TestAutoFieldsOfNewNASMessages(t *testing.T) {
	r := autoRole(t, "")
	for _, c := range []struct {
		nas, outer, path string
		auto             bool
		note             string
	}{
		{"PDUSessionModificationRequest", "ULNASTransport", "ProcedureTransactionID", true, "allocated when sent"},
		{"PDUSessionModificationComplete", "ULNASTransport", "ProcedureTransactionID", true, "the PTI of the command it answers"},
		{"PDUSessionAuthenticationComplete", "ULNASTransport", "PDUSessionID", true, "the PDU session of ue.session"},
		{"ControlPlaneServiceRequest", "", "ControlPlaneServiceTypeAndNgKSI", false, "ngKSI"},
		{"ControlPlaneServiceRequest", "", "UplinkDataStatus/Value", true, "the established PDU sessions"},
		{"NotificationResponse", "", "PDUSessionStatus/Value", true, "the established PDU sessions"},
		{"NetworkSliceSpecificAuthenticationComplete", "", "SNSSAI", false, "Network Slice-Specific Authentication Command"},
		{"NetworkSliceSpecificAuthenticationComplete", "", "EAPMessage", false, ""},
	} {
		auto, note := r.autoNAS(c.nas, c.outer, c.path)
		if auto != c.auto || !strings.Contains(note, c.note) || c.note == "" && note != "" {
			t.Errorf("%s %s: %v %q", c.nas, c.path, auto, note)
		}
	}
}
