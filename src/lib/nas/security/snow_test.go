package security

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// Byte-aligned standard vectors also published in free5gc/nas v1.2.2
// security/security_test.go: NEA1 set 3, NIA1 sets 1 and 4.
func TestSnowStandardVectors(t *testing.T) {
	input := unhex("ad9c441f890b38c457a49d421407e8")
	k := key("5acb1d644c0d51204ea5f1451010d852")
	got, err := Cipher(NEA1, k, 0xfa556b26, 3, 1, input)
	if err != nil || !bytes.Equal(got, unhex("ba0f31300334c56b52a7497cbac046")) {
		t.Fatalf("NEA1: %x %v", got, err)
	}
	back, err := Cipher(NEA1, k, 0xfa556b26, 3, 1, got)
	if err != nil || !bytes.Equal(back, input) {
		t.Fatal("NEA1 decrypt", err)
	}
	for _, tc := range []struct {
		k                 string
		count             uint32
		bearer, direction uint8
		msg               string
		mac               uint32
	}{
		{"2bd6459f82c5b300952c49104881ff48", 0x38a6f056, 0x1f, 0, "3332346263393861373479", 0x731f1165},
		{"83fd23a244a74cf358da3019f1722635", 0x36af6144, 0x0f, 1, "35c68716633c66fb750c266865d53c11ea05b1e9fa49c8398d48e1efa5909d3947902837f5ae96d5a05bc8d61ca8dbef1b13a4b4abfe4fb1006045b674bb54729304c382be53a5af05556176f6eaa2ef1d05e4b083181ee674cda5a485f74d7a", 0xbba74492},
	} {
		t.Run(tc.k, func(t *testing.T) {
			t.Parallel()
			got, err := Integrity(NIA1, key(tc.k), tc.count, tc.bearer, tc.direction, unhex(tc.msg))
			if err != nil || binary.BigEndian.Uint32(got[:]) != tc.mac {
				t.Fatalf("NIA1: %x %v", got, err)
			}
		})
	}
}
