package nasType

import (
	"bytes"
	"testing"
)

func TestUEParameters(t *testing.T) {
	raw, e := EncodeNullSUCI([]byte{0, 0xf1, 0x10}, "001010000000001", "0")
	if e != nil {
		t.Fatal(e)
	}
	id, e := DecodeMobileIdentity(raw)
	if e != nil || id.SUCI != "suci-0-001-01-0-0-0-0000000001" {
		t.Fatalf("identity %s %v", id.SUCI, e)
	}
	d, e := EncodeDNN("ims.example")
	if e != nil || !bytes.Equal(d, []byte{3, 'i', 'm', 's', 7, 'e', 'x', 'a', 'm', 'p', 'l', 'e'}) {
		t.Fatal("DNN", e)
	}
	for _, bad := range []string{"", "a..b", "a/b"} {
		if _, e = EncodeDNN(bad); e == nil {
			t.Fatal("invalid DNN accepted")
		}
	}
	if _, e = EncodeNullSUCI([]byte{0, 0xf1, 0x10}, "999010000000001", "0"); e == nil {
		t.Fatal("mismatched PLMN accepted")
	}
	if _, e = EncodeSNSSAI(1, "abcd"); e == nil {
		t.Fatal("short SD accepted")
	}
	for _, raw := range [][]byte{{1, 10, 0, 0, 1}, {2, 0, 0, 0, 0, 0, 0, 0, 1}, {3, 0, 0, 0, 0, 0, 0, 0, 1, 10, 0, 0, 1}} {
		kind, ip, iid, e := DecodePDUAddress(raw)
		if e != nil || kind != raw[0] {
			t.Fatal(e)
		}
		if kind != 2 && !bytes.Equal(ip, []byte{10, 0, 0, 1}) {
			t.Fatal("IPv4 address")
		}
		if kind != 1 && len(iid) != 8 {
			t.Fatal("IPv6 interface ID")
		}
		if _, _, _, e = DecodePDUAddress(raw[:len(raw)-1]); e == nil {
			t.Fatal("truncated address accepted")
		}
	}
}
