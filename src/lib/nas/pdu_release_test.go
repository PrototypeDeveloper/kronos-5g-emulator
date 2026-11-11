package libnas_test

import (
	"bytes"
	"encoding/hex"
	"libnas"
	"testing"
)

func TestPDUReleaseWire(t *testing.T) {
	for _, s := range []string{"2e0102d1", "2e0102d15924", "2e0102d224", "2e0102d324", "2e0102d4", "2e0102d459247b000180"} {
		raw, _ := hex.DecodeString(s)
		v, err := libnas.Decode(raw, nil)
		if err != nil {
			t.Fatal(s, err)
		}
		got, err := libnas.Encode(v, nil)
		if err != nil || !bytes.Equal(got, raw) {
			t.Fatal(s, err, got)
		}
	}
	for _, s := range []string{"2e0102d3", "2e0102d2", "2e0102d159", "2e0102d47b000280"} {
		raw, _ := hex.DecodeString(s)
		if _, err := libnas.Decode(raw, nil); err == nil {
			t.Fatal("malformed release accepted", s)
		}
	}
}
