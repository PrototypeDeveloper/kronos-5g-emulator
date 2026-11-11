package nasType

import (
	"encoding/hex"
	"testing"
)

func TestMobileIdentity(t *testing.T) {
	tests := []struct {
		name, wire, want string
		bad              bool
	}{
		{"null", "0199f999f0ff00000000000010", "suci-0-999-99-0-0-0-0000000001", false},
		{"profile-a", "0199f999f0ff0101abcd", "suci-0-999-99-0-1-1-abcd", false},
		{"guti", "f299f99901004001020304", "", false},
		{"truncated", "0199", "", true},
		{"invalid-plmn", "019af999f0ff00000000000010", "", true},
		{"invalid-routing", "0199f999f01f00000000000010", "", true},
		{"null-key-id", "0199f999f0ff00010000000010", "", true},
		{"invalid-msin", "0199f999f0ff00000a", "", true},
		{"non-imsi", "1199f999f0ff00000000000010", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.wire)
			id, err := DecodeMobileIdentity(b)
			if (err != nil) != tt.bad {
				t.Fatalf("error=%v", err)
			}
			if err == nil && id.SUCI != tt.want {
				t.Fatalf("SUCI=%s", id.SUCI)
			}
		})
	}
}
func TestPLMNThreeDigitMNC(t *testing.T) {
	mcc, mnc, err := DecodePLMN([]byte{0x13, 0x00, 0x62})
	if err != nil || mcc != "310" || mnc != "260" {
		t.Fatalf("%s %s %v", mcc, mnc, err)
	}
}
