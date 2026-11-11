package libnas_test

import (
	"bytes"
	nas "libnas"
	"libnas/nasMessage"
	"testing"
)

func TestObservationDoesNotConsumeExtraCount(t *testing.T) {
	tx, rx := securityContext(), securityContext()
	calls := 0
	var plainCopy []byte
	observe := func(wire, plain []byte) {
		calls++
		plainCopy = bytes.Clone(plain)
		if bytes.Equal(wire, plain) {
			t.Error("expected protected wire")
		}
	}
	wire, err := nas.Encode(&nasMessage.RegistrationComplete{}, &nas.EncodeContext{Security: &tx, Direction: nas.Uplink, SecurityHeaderType: 2, Observe: observe})
	if err != nil {
		t.Fatal(err)
	}
	saved := bytes.Clone(wire)
	if _, err = nas.Decode(wire, &nas.DecodeContext{Security: &rx, Direction: nas.Uplink, Observe: observe}); err != nil {
		t.Fatal(err)
	}
	if calls != 2 || tx.UplinkCount != 1 || rx.UplinkCount != 1 || !bytes.Equal(wire, saved) || !bytes.Equal(plainCopy, []byte{0x7e, 0, 0x43}) {
		t.Fatal("observation changed NAS/security state")
	}
	bad := bytes.Clone(wire)
	bad[2] ^= 1
	if _, err = nas.Decode(bad, &nas.DecodeContext{Security: &rx, Direction: nas.Uplink, Observe: observe}); err == nil || calls != 2 || rx.UplinkCount != 1 {
		t.Fatal("invalid NAS observed or consumed count")
	}
}
