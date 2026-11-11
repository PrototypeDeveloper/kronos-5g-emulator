package libnas

import (
	"bytes"
	"encoding/hex"
	"libnas/nasMessage"
	"testing"
)

func TestDeregistrationVectorsAndSecurity(t *testing.T) {
	for _, typ := range []byte{0x21, 0x29} {
		wire, _ := hex.DecodeString("7e004521000bf202f83901004012345678")
		wire[3] = typ
		value, err := Decode(wire, nil)
		if err != nil {
			t.Fatal(err)
		}
		msg := value.(*nasMessage.DeregistrationRequestUE)
		if msg.DeregistrationTypeAndNgKSI != typ || len(msg.MobileIdentity) != 11 {
			t.Fatal(msg)
		}
		encoded, err := Encode(msg, nil)
		if err != nil || !bytes.Equal(encoded, wire) {
			t.Fatal(encoded, err)
		}
		for _, header := range []SecurityHeaderType{SecurityHeaderIntegrityProtected, SecurityHeaderIntegrityProtectedAndCiphered} {
			tx := &SecurityContext{IntegrityAlgorithm: 2, CipheringAlgorithm: 2}
			rx := *tx
			protected, err := Encode(msg, &EncodeContext{Security: tx, Direction: Uplink, SecurityHeaderType: header})
			if err != nil {
				t.Fatal(err)
			}
			corrupt := append([]byte(nil), protected...)
			corrupt[2] ^= 1
			if _, err := Decode(corrupt, &DecodeContext{Security: &rx, Direction: Uplink}); err == nil {
				t.Fatal("bad MAC accepted")
			}
			if _, err := Decode(protected, &DecodeContext{Security: &rx, Direction: Uplink}); err != nil {
				t.Fatal(err)
			}
			if _, err := Decode(protected, &DecodeContext{Security: &rx, Direction: Uplink}); err == nil {
				t.Fatal("replay accepted")
			}
		}
		for i := 0; i < len(wire); i++ {
			if _, err := Decode(wire[:i], nil); err == nil {
				t.Fatalf("truncation %d accepted", i)
			}
		}
	}
	accept, _ := Encode(&nasMessage.DeregistrationAcceptUE{}, nil)
	if !bytes.Equal(accept, []byte{0x7e, 0, 0x46}) {
		t.Fatal(accept)
	}
	if _, err := Decode(append(accept, 0), nil); err == nil {
		t.Fatal("extra Accept body accepted")
	}
}
