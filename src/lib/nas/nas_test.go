package libnas_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	nas "libnas"
	"libnas/nasMessage"
	"libnas/nasType"
	"libnas/security"
	"reflect"
	"testing"
)

func unhex(s string) []byte {
	b, e := hex.DecodeString(s)
	if e != nil {
		panic(e)
	}
	return b
}
func securityContext() nas.SecurityContext {
	s := nas.SecurityContext{IntegrityAlgorithm: 2, CipheringAlgorithm: 2}
	copy(s.KnasInt[:], unhex("2b7e151628aed2a6abf7158809cf4f3c"))
	copy(s.KnasEnc[:], unhex("d3c5d592327fb11c4035c6680af8c6d1"))
	return s
}
func TestPlainRegistrationFixture(t *testing.T) {
	// Same registration PDU as lib/ngap's InitialUEMessage wire fixture.
	wire := unhex("7e004179000d0199f9590000000001000000602e04f0f0f0f0")
	value, e := nas.Decode(wire, nil)
	if e != nil {
		t.Fatal(e)
	}
	msg, ok := value.(*nasMessage.RegistrationRequest)
	if !ok {
		t.Fatalf("type %T", value)
	}
	if msg.RegistrationTypeAndNgKSI != 0x79 || len(msg.MobileIdentity) != 13 || !msg.HasUESecurityCapability || !bytes.Equal(msg.UESecurityCapability.Value, []byte{0xf0, 0xf0, 0xf0, 0xf0}) {
		t.Fatalf("decoded %+v", msg)
	}
	encoded, e := nas.Encode(msg, nil)
	if e != nil || !bytes.Equal(encoded, wire) {
		t.Fatalf("encoded %x: %v", encoded, e)
	}
	wire[6] = 0xff
	if msg.MobileIdentity[0] != 1 {
		t.Fatal("decoded identity aliases input")
	}
}
func TestSecurityHeadersAndDirections(t *testing.T) {
	for _, d := range []nas.Direction{nas.Uplink, nas.Downlink} {
		for _, sht := range []nas.SecurityHeaderType{0, 1, 2, 3, 4} {
			var msg any = &nasMessage.RegistrationComplete{}
			if sht == 3 {
				msg = &nasMessage.SecurityModeCommand{SelectedNASAlgorithms: 0x22, ReplayedUESecurityCapabilities: []byte{0x20, 0x20}}
			}
			if sht == 4 {
				msg = &nasMessage.SecurityModeComplete{}
			}
			tx, rx := securityContext(), securityContext()
			encoded, e := nas.Encode(msg, &nas.EncodeContext{Security: &tx, SecurityHeaderType: sht, Direction: d})
			if e != nil {
				t.Fatal(e)
			}
			original := append([]byte(nil), encoded...)
			decoded, e := nas.Decode(encoded, &nas.DecodeContext{Security: &rx, Direction: d})
			if e != nil {
				t.Fatalf("header=%d dir=%d: %v", sht, d, e)
			}
			if !reflect.DeepEqual(decoded, msg) {
				t.Fatalf("round trip: %#v != %#v", decoded, msg)
			}
			if !bytes.Equal(encoded, original) {
				t.Fatal("decode mutated input")
			}
			if tx != rx {
				t.Fatal("endpoint counts diverged")
			}
			want := nas.Count(0)
			if sht != 0 {
				want = 1
			}
			if d == nas.Uplink {
				if tx.UplinkCount != want || tx.DownlinkCount != 0 {
					t.Fatal("wrong UL count")
				}
			} else if tx.DownlinkCount != want || tx.UplinkCount != 0 {
				t.Fatal("wrong DL count")
			}
		}
	}
}
func TestTamperingReplayAndRollover(t *testing.T) {
	tx, rx := securityContext(), securityContext()
	tx.UplinkCount = 255
	rx.UplinkCount = 255
	for i := 0; i < 3; i++ {
		wire, e := nas.Encode(&nasMessage.RegistrationComplete{}, &nas.EncodeContext{Security: &tx, SecurityHeaderType: 2})
		if e != nil {
			t.Fatal(e)
		}
		for _, offset := range []int{2, 6, 7} {
			bad := append([]byte(nil), wire...)
			bad[offset] ^= 1
			before := rx
			if msg, e := nas.Decode(bad, &nas.DecodeContext{Security: &rx}); e == nil || msg != nil {
				t.Fatal("tampering accepted")
			}
			if rx != before {
				t.Fatal("count changed after invalid MAC")
			}
		}
		if _, e = nas.Decode(wire, &nas.DecodeContext{Security: &rx}); e != nil {
			t.Fatal(e)
		}
		before := rx
		if _, e = nas.Decode(wire, &nas.DecodeContext{Security: &rx}); e == nil {
			t.Fatal("replay accepted")
		}
		if rx != before {
			t.Fatal("replay changed COUNT")
		}
	}
	if tx.UplinkCount != 258 || rx.UplinkCount != 258 {
		t.Fatalf("rollover: %v %v", tx.UplinkCount, rx.UplinkCount)
	}
}
func TestWrongSecurityAndExhaustion(t *testing.T) {
	tx := securityContext()
	wire, e := nas.Encode(&nasMessage.RegistrationComplete{}, &nas.EncodeContext{Security: &tx, SecurityHeaderType: 2})
	if e != nil {
		t.Fatal(e)
	}
	for _, change := range []func(*nas.SecurityContext){func(s *nas.SecurityContext) { s.KnasInt[0] ^= 1 }, func(s *nas.SecurityContext) { s.NASConnectionID = 2 }, func(s *nas.SecurityContext) { s.IntegrityAlgorithm = 3 }, func(s *nas.SecurityContext) { s.CipheringAlgorithm = 3 }} {
		rx := securityContext()
		change(&rx)
		before := rx
		if _, e = nas.Decode(wire, &nas.DecodeContext{Security: &rx}); e == nil {
			t.Fatal("wrong security accepted")
		}
		if rx != before {
			t.Fatal("failure consumed count")
		}
	}
	rx := securityContext()
	if _, e = nas.Decode(wire, &nas.DecodeContext{Security: &rx, Direction: nas.Downlink}); !errors.Is(e, nas.ErrIntegrity) {
		t.Fatal(e)
	}
	tx = securityContext()
	rx = tx
	tx.UplinkCount = nasType.MaxCount
	rx.UplinkCount = nasType.MaxCount
	wire, e = nas.Encode(&nasMessage.RegistrationComplete{}, &nas.EncodeContext{Security: &tx, SecurityHeaderType: 1})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = nas.Decode(wire, &nas.DecodeContext{Security: &rx}); e != nil {
		t.Fatal(e)
	}
	if _, e = nas.Encode(&nasMessage.RegistrationComplete{}, &nas.EncodeContext{Security: &tx, SecurityHeaderType: 1}); !errors.Is(e, nas.ErrCountExhausted) {
		t.Fatal(e)
	}
}
func TestSMTransportAndSessionID(t *testing.T) {
	sm := &nasMessage.PDUSessionEstablishmentRequest{PDUSessionID: 2, ProcedureTransactionID: 1, IntegrityProtectionMaximumDataRate: []byte{0xff, 0xff}, HasPDUSessionType: true, PDUSessionType: nasType.PDUSessionType{Value: 1}, HasSSCMode: true, SSCMode: nasType.SSCMode{Value: 1}}
	inner, e := nas.Encode(sm, nil)
	if e != nil {
		t.Fatal(e)
	}
	// Session ID=2 must not be interpreted as an encrypted security header.
	decoded, e := nas.Decode(inner, nil)
	if e != nil || !reflect.DeepEqual(decoded, sm) {
		t.Fatalf("5GSM: %#v %v", decoded, e)
	}
	tx, rx := securityContext(), securityContext()
	if _, e = nas.Encode(sm, &nas.EncodeContext{Security: &tx, SecurityHeaderType: 2}); e == nil {
		t.Fatal("direct protected 5GSM accepted")
	}
	mm := &nasMessage.ULNASTransport{PayloadContainerType: 1, PayloadContainer: inner}
	wire, e := nas.Encode(mm, &nas.EncodeContext{Security: &tx, SecurityHeaderType: 2})
	if e != nil {
		t.Fatal(e)
	}
	outer, e := nas.Decode(wire, &nas.DecodeContext{Security: &rx})
	if e != nil {
		t.Fatal(e)
	}
	decoded, e = nas.Decode(outer.(*nasMessage.ULNASTransport).PayloadContainer, nil)
	if e != nil || !reflect.DeepEqual(decoded, sm) {
		t.Fatal("nested round trip")
	}
}
func TestInvalidInputsAndNoStateCommitOnPayloadFailure(t *testing.T) {
	for _, b := range [][]byte{nil, {0x7e}, {0x7e, 2}, {0x7e, 0x10, 0x43}, {0x7e, 5, 0}, {0x2e, 16, 1, 0xd6, 0}, {0x7e, 0, 0xff}, {0x7e, 0, 0x41, 0, 0xff, 0xff}} {
		if v, e := nas.Decode(b, nil); e == nil || v != nil {
			t.Fatalf("invalid accepted %x", b)
		}
	}
	var typedNil *nasMessage.RegistrationRequest
	if _, e := nas.Encode(typedNil, nil); e == nil {
		t.Fatal("nil accepted")
	}
	if _, e := nas.Encode(struct{}{}, nil); e == nil {
		t.Fatal("unknown type accepted")
	}
	s := securityContext()
	before := s
	if _, e := nas.Encode(&nasMessage.RegistrationComplete{}, &nas.EncodeContext{Security: &s, SecurityHeaderType: 3}); e == nil {
		t.Fatal("wrong new context message")
	}
	if s != before {
		t.Fatal("failed encode consumed count")
	}
	// Valid MAC over malformed inner NAS must still leave receive state unchanged.
	wire := []byte{0x7e, 1, 0, 0, 0, 0, 0, 0x7e, 0, 0xff}
	mac, _ := security.Integrity(2, s.KnasInt, 0, 1, 0, wire[6:])
	copy(wire[2:6], mac[:])
	if _, e := nas.Decode(wire, &nas.DecodeContext{Security: &s}); e == nil {
		t.Fatal("malformed authenticated payload accepted")
	}
	if s != before {
		t.Fatal("payload failure consumed count")
	}
}
func FuzzDecode(f *testing.F) {
	for _, s := range []string{"7e0043", "7e004179000d0199f9590000000001000000602e04f0f0f0f0", "2e0201c1ffff91a1", "7e0200000000007e0043"} {
		f.Add(unhex(s))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		s := securityContext()
		before := s
		input := append([]byte(nil), data...)
		value, e := nas.Decode(data, &nas.DecodeContext{Security: &s})
		if !bytes.Equal(input, data) {
			t.Fatal("input mutated")
		}
		if e != nil {
			if value != nil || s != before {
				t.Fatal("failed decode returned a value or mutated state")
			}
			return
		}
		if _, e := nas.Encode(value, nil); e != nil {
			t.Fatal(e)
		}
	})
}
