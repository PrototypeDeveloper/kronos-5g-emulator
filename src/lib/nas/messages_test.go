package libnas_test

import (
	"bytes"
	"fmt"
	nas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	"testing"
)

func TestMessageWireFixtures(t *testing.T) {
	fixtures := []struct{ hex, name string }{
		{"7e00420101", "RegistrationAccept"},
		{"7e0043", "RegistrationComplete"},
		{"7e00440b", "RegistrationReject"},
		{"7e00560002000021000102030405060708090a0b0c0d0e0f2010000102030405060708090a0b0c0d0e0f", "AuthenticationRequest"},
		{"7e00572d10000102030405060708090a0b0c0d0e0f", "AuthenticationResponse"},
		{"7e0058", "AuthenticationReject"},
		{"7e005914300e000102030405060708090a0b0c0d", "AuthenticationFailure"},
		{"7e005b01", "IdentityRequest"},
		{"7e005c000d0199f959000000000100000060", "IdentityResponse"},
		{"7e005d2200022020e1360101", "SecurityModeCommand"},
		{"7e005e7100037e0043", "SecurityModeComplete"},
		{"7e005f18", "SecurityModeReject"},
		{"7e006460", "Status5GMM"},
		{"7e00670100062e0101c1ffff120181", "ULNASTransport"},
		{"7e00680100052e0101c31a1201", "DLNASTransport"},
		{"2e0101c1ffff91a1", "PDUSessionEstablishmentRequest"},
		{"2e0101c21100040100012006010001010001", "PDUSessionEstablishmentAccept"},
		{"2e0101c31a", "PDUSessionEstablishmentReject"},
		{"2e0101d61a", "Status5GSM"},
	}
	for _, tc := range fixtures {
		t.Run(tc.name, func(t *testing.T) {
			wire := unhex(tc.hex)
			msg, e := nas.Decode(wire, nil)
			if e != nil {
				t.Fatal(e)
			}
			if got := fmt.Sprintf("%T", msg); got != "*nasMessage."+tc.name {
				t.Fatalf("type %s", got)
			}
			b, e := nas.Encode(msg, nil)
			if e != nil || !bytes.Equal(b, wire) {
				t.Fatalf("encoded %x: %v", b, e)
			}
		})
	}
}
func TestOptionalIEBoundaries(t *testing.T) {
	for _, msg := range []*nasMessage.AuthenticationRequest{
		{ABBA: []byte{0, 0}, HasAuthenticationParameterRAND: true, AuthenticationParameterRAND: nasType.AuthenticationParameterRAND{Value: make([]byte, 15)}},
		{ABBA: []byte{0, 0}, HasAuthenticationParameterAUTN: true, AuthenticationParameterAUTN: nasType.AuthenticationParameterAUTN{Value: make([]byte, 17)}},
		{ABBA: []byte{0, 0}, HasEAPMessage: true, EAPMessage: nasType.EAPMessage{Value: make([]byte, 1501)}},
	} {
		if _, e := nas.Encode(msg, nil); e == nil {
			t.Fatal("invalid IE accepted")
		}
	}
	if _, e := nas.Encode(&nasMessage.SecurityModeCommand{ReplayedUESecurityCapabilities: []byte{0x20, 0x20}, HasIMEISVRequest: true, IMEISVRequest: nasType.IMEISVRequest{Value: 16}}, nil); e == nil {
		t.Fatal("half-octet overflow")
	}
	for _, wire := range []string{"7e00560002000021", "7e005600020000201000", "7e00560002000078010000", "7e0056000200007f"} {
		if _, e := nas.Decode(unhex(wire), nil); e == nil {
			t.Fatalf("invalid IE accepted %s", wire)
		}
	}
	payload := bytes.Repeat([]byte{0x42}, 300)
	msg := &nasMessage.SecurityModeComplete{HasNASMessageContainer: true, NASMessageContainer: nasType.NASMessageContainer{Value: payload}}
	wire, e := nas.Encode(msg, nil)
	if e != nil {
		t.Fatal(e)
	}
	if wire[4] != 1 || wire[5] != 44 {
		t.Fatal("TLV-E length is not big endian")
	}
	got, e := nas.Decode(wire, nil)
	if e != nil {
		t.Fatal(e)
	}
	if !bytes.Equal(got.(*nasMessage.SecurityModeComplete).NASMessageContainer.Value, payload) {
		t.Fatal("long TLV-E value")
	}
}

func TestNamedOptionalPresenceAndEditing(t *testing.T) {
	// A zero-valued optional nibble is distinct from an absent IE.
	msg := &nasMessage.SecurityModeCommand{
		ReplayedUESecurityCapabilities: []byte{0x20, 0x20},
		HasIMEISVRequest:               true,
		IMEISVRequest:                  nasType.IMEISVRequest{Value: 0},
	}
	wire, err := nas.Encode(msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := nas.Decode(wire, nil)
	if err != nil {
		t.Fatal(err)
	}
	got := decoded.(*nasMessage.SecurityModeCommand)
	if !got.HasIMEISVRequest || got.IMEISVRequest.Value != 0 {
		t.Fatal("present zero was lost")
	}
	got.IMEISVRequest.Value = 1
	edited, err := nas.Encode(got, nil)
	if err != nil {
		t.Fatal(err)
	}
	if edited[len(edited)-1] != 0xe1 {
		t.Fatal("named field edit was not encoded")
	}
	// A false presence flag suppresses even an otherwise invalid value.
	got.HasIMEISVRequest = false
	got.IMEISVRequest.Value = 255
	omitted, err := nas.Encode(got, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(omitted) != len(wire)-1 {
		t.Fatal("absent IE was emitted")
	}
	decoded, err = nas.Decode(omitted, nil)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.(*nasMessage.SecurityModeCommand).HasIMEISVRequest {
		t.Fatal("absent flag was set")
	}
}

func TestNamedAuthenticationIEsInAnyOrder(t *testing.T) {
	// AUTN before RAND: decode by IEI, then encode in the message's canonical order.
	wire := unhex("7e0056000200002010101112131415161718191a1b1c1d1e1f21000102030405060708090a0b0c0d0e0f")
	value, err := nas.Decode(wire, nil)
	if err != nil {
		t.Fatal(err)
	}
	msg := value.(*nasMessage.AuthenticationRequest)
	if !msg.HasAuthenticationParameterRAND || !msg.HasAuthenticationParameterAUTN || msg.HasEAPMessage {
		t.Fatal("presence flags")
	}
	if msg.AuthenticationParameterRAND.Value[0] != 0 || msg.AuthenticationParameterAUTN.Value[0] != 0x10 {
		t.Fatal("IE values assigned to wrong fields")
	}
	wire[10] = 0xff
	if msg.AuthenticationParameterAUTN.Value[0] != 0x10 {
		t.Fatal("typed value aliases input")
	}
	msg.AuthenticationParameterRAND.Value[0] = 0x42
	edited, err := nas.Encode(msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	if edited[7] != 0x21 || edited[8] != 0x42 {
		t.Fatal("named RAND edit not encoded")
	}
	// TS 24.501 7.6.3: first singleton occurrence wins, even if later values differ.
	duplicate := append(append([]byte(nil), edited...), edited[7:24]...)
	duplicate[len(duplicate)-16] = 0x99
	value, err = nas.Decode(duplicate, nil)
	if err != nil || value.(*nasMessage.AuthenticationRequest).AuthenticationParameterRAND.Value[0] != 0x42 {
		t.Fatal("duplicate RAND replaced the first occurrence", err)
	}
}

func TestRegistrationAcceptNamedServiceAreaAndTimer(t *testing.T) {
	msg := &nasMessage.RegistrationAccept{
		RegistrationResult: []byte{1},
		HasServiceAreaList: true, ServiceAreaList: nasType.ServiceAreaList{Value: []byte{0x40, 0x99, 0xf9, 0x59}},
		HasT3324Value: true, T3324Value: nasType.T3324Value{Value: 0},
	}
	wire, err := nas.Encode(msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wire, unhex("7e0042010127044099f9596a0100")) {
		t.Fatalf("wrong service area/timer encoding %x", wire)
	}
	value, err := nas.Decode(wire, nil)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(*nasMessage.RegistrationAccept)
	if !got.HasServiceAreaList || !got.HasT3324Value || got.T3324Value.Value != 0 {
		t.Fatal("named fields missing")
	}
	msg.ServiceAreaList.Value = []byte{0}
	if _, err = nas.Encode(msg, nil); err == nil {
		t.Fatal("short ServiceAreaList accepted")
	}
}

func TestPresentEmptyOptionalValue(t *testing.T) {
	msg := &nasMessage.RegistrationRequest{
		RegistrationTypeAndNgKSI: 0x79, MobileIdentity: unhex("0199f959000000000100000060"),
		HasLADNIndication: true, LADNIndication: nasType.LADNIndication{},
	}
	wire, err := nas.Encode(msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	value, err := nas.Decode(wire, nil)
	if err != nil {
		t.Fatal(err)
	}
	got := value.(*nasMessage.RegistrationRequest)
	if !got.HasLADNIndication || len(got.LADNIndication.Value) != 0 {
		t.Fatal("empty but present LADN indication lost")
	}
}
