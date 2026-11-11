package libnas_test

import (
	"bytes"
	"libnas"
	"libnas/nasMessage"
	"testing"
)

func TestNASOptionalIECompatibility(t *testing.T) {
	for _, msg := range []any{
		&nasMessage.RegistrationRequest{RegistrationTypeAndNgKSI: 0x71, MobileIdentity: []byte{1, 2, 3, 4}},
		&nasMessage.SecurityModeComplete{},
		&nasMessage.PDUSessionEstablishmentRequest{PDUSessionID: 1, ProcedureTransactionID: 1, IntegrityProtectionMaximumDataRate: []byte{0xff, 0xff}},
	} {
		wire, err := libnas.Encode(msg, nil)
		if err != nil {
			t.Fatal(err)
		}
		// These IEIs are unknown in all three messages; their values must not
		// be interpreted as subsequent IEs or alter known mandatory fields.
		for _, extra := range [][]byte{{0xe1}, {0x41, 2, 0x7f, 0xff}, {0x60, 0, 2, 0x7f, 0xff}} {
			input := append(append([]byte(nil), wire...), extra...)
			v, err := libnas.Decode(input, nil)
			if err != nil {
				t.Fatalf("%T %x: %v", msg, extra, err)
			}
			got, err := libnas.Encode(v, nil)
			if err != nil || !bytes.Equal(got, wire) {
				t.Fatalf("known fields changed: %T %v", msg, err)
			}
		}
		input := append(append([]byte(nil), wire...), 0x7e, 0, 0)
		if _, err := libnas.Decode(input, nil); err == nil {
			t.Fatalf("%T ignored comprehension-required IE", msg)
		}
	}
}
