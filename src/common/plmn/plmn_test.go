package plmn

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestCodec(t *testing.T) {
	for _, tc := range []struct{ mcc, mnc, hex string }{{"208", "93", "02f839"}, {"001", "001", "001100"}, {"999", "70", "99f907"}} {
		want, _ := hex.DecodeString(tc.hex)
		b, ok := Encode(tc.mcc, tc.mnc)
		if !ok || !bytes.Equal(b, want) {
			t.Fatalf("encode %s/%s: %x", tc.mcc, tc.mnc, b)
		}
		if mcc, mnc, ok := Decode(b); !ok || mcc != tc.mcc || mnc != tc.mnc {
			t.Fatalf("decode %x: %s/%s", b, mcc, mnc)
		}
	}
	broken, _ := hex.DecodeString("2f8390")
	if _, _, ok := Decode(broken); ok || String(broken) != "2f8390" {
		t.Fatal("2f8390 accepted")
	}
	if _, ok := Encode("20", "93"); ok {
		t.Fatal("2-digit MCC accepted")
	}
}
