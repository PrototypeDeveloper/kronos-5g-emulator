package libnas

import (
	"bytes"
	"encoding/hex"
	"libnas/nasMessage"
	"testing"
)

func TestServiceRequestWireAndSecurity(t *testing.T) {
	raw, _ := hex.DecodeString("7e004c220007f40040123456784002020050020200")
	msg, err := Decode(raw, nil)
	if err != nil {
		t.Fatal(err)
	}
	req := msg.(*nasMessage.ServiceRequest)
	if req.ServiceTypeAndNgKSI != 0x22 || !req.HasUplinkDataStatus || !bytes.Equal(req.PDUSessionStatus, []byte{2, 0}) {
		t.Fatal(req)
	}
	back, err := Encode(req, nil)
	if err != nil || !bytes.Equal(back, raw) {
		t.Fatal(err, back)
	}
	for _, n := range []int{0, 1, 2, 3, 4, 5, 6, 12, 14, 15, 18} {
		if _, err := Decode(raw[:n], nil); err == nil {
			t.Fatalf("truncated length %d accepted", n)
		}
	}
	for _, header := range []SecurityHeaderType{SecurityHeaderIntegrityProtected, SecurityHeaderIntegrityProtectedAndCiphered} {
		tx := &SecurityContext{IntegrityAlgorithm: 2, CipheringAlgorithm: 2}
		rx := *tx
		protected, err := Encode(req, &EncodeContext{Security: tx, Direction: Uplink, SecurityHeaderType: header})
		if err != nil {
			t.Fatal(err)
		}
		bad := bytes.Clone(protected)
		bad[2] ^= 1
		if _, err = Decode(bad, &DecodeContext{Security: &rx, Direction: Uplink}); err == nil || rx.UplinkCount != 0 {
			t.Fatal("bad MAC accepted")
		}
		if _, err = Decode(protected, &DecodeContext{Security: &rx, Direction: Uplink}); err != nil {
			t.Fatal(err)
		}
		if _, err = Decode(protected, &DecodeContext{Security: &rx, Direction: Uplink}); err == nil {
			t.Fatal("replay accepted")
		}
	}
	accept := &nasMessage.ServiceAccept{HasPDUSessionStatus: true, PDUSessionStatus: []byte{2, 0}, HasReactivationResult: true, ReactivationResult: []byte{4, 0}}
	b, err := Encode(accept, nil)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := hex.DecodeString("7e004e5002020026020400")
	if !bytes.Equal(b, want) {
		t.Fatal(b)
	}
	if _, err := Decode(b, nil); err != nil {
		t.Fatal(err)
	}
}
