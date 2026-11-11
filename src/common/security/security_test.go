package security

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func h(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// TS 35.208 §4.3, test set 1.
func TestMilenageTestSet1(t *testing.T) {
	k, rand := h("465b5ce8b199b49faa5f0a2ee238a6bc"), h("23553cbe9637a89d218ae64dae47bf35")
	sqn, amf := h("ff9bb4d0b607"), h("b9b9")
	opc, err := OPc(k, h("cdc202d5123e20f62b6d676ac72cb318"))
	if err != nil || !bytes.Equal(opc, h("cd63cb71954a9f4e48a5994e37a02baf")) {
		t.Fatalf("OPc %x %v", opc, err)
	}
	m, err := NewMilenage(k, opc)
	if err != nil {
		t.Fatal(err)
	}
	macA, macS := m.F1(rand, sqn, amf)
	res, ck, ik, ak := m.F2345(rand)
	for name, c := range map[string][2][]byte{
		"f1":  {macA, h("4a9ffac354dfafb3")},
		"f1*": {macS, h("01cfaf9ec4e871e9")},
		"f2":  {res, h("a54211d5e3ba50bf")},
		"f3":  {ck, h("b40ba9a3c58b2a05bbf0d987b21bf8cb")},
		"f4":  {ik, h("f769bcd751044604127672711c6d3441")},
		"f5":  {ak, h("aa689c648370")},
		"f5*": {m.F5Star(rand), h("451e8beca43b")},
	} {
		if !bytes.Equal(c[0], c[1]) {
			t.Errorf("%s = %x, want %x", name, c[0], c[1])
		}
	}
}

func TestKDF(t *testing.T) {
	key := h("000102030405060708090a0b0c0d0e0f")
	mac := hmac.New(sha256.New, key)
	mac.Write(h("6b" + "616263" + "0003" + "01" + "0001"))
	if got := KDF(key, 0x6b, []byte("abc"), []byte{1}); !bytes.Equal(got, mac.Sum(nil)) {
		t.Fatalf("KDF %x", got)
	}
	if s := ServingNetworkName("208", "93"); s != "5G:mnc093.mcc208.3gppnetwork.org" {
		t.Fatal(s)
	}
	if s := ServingNetworkName("001", "001"); s != "5G:mnc001.mcc001.3gppnetwork.org" {
		t.Fatal(s)
	}
	enc, integ := NASKeys(make([]byte, 32), 2, 2)
	full := KDF(make([]byte, 32), 0x69, []byte{2}, []byte{2})
	if !bytes.Equal(integ[:], full[16:]) || bytes.Equal(enc[:], integ[:]) {
		t.Fatal("NAS keys must be the low 128 bits of distinct derivations")
	}
}

func TestNullSUCI(t *testing.T) {
	// The sample signal.txt carries this SUCI for IMSI 208930000000003.
	b, err := NullSUCI("208", "93", "0000", "0000000003")
	if err != nil || !bytes.Equal(b, h("0102f839000000000000000030")) {
		t.Fatalf("%x %v", b, err)
	}
	b, _ = NullSUCI("001", "01", "", "123456789")
	if !bytes.Equal(b, h("0100f110"+"f0ff"+"0000"+"2143658709"[:8]+"f9")) {
		t.Fatalf("%x", b)
	}
}

func TestIMEI(t *testing.T) {
	// 490154203237518 (a known valid IMEI): the check digit 8 follows the first 14 digits.
	b, err := IMEI("4901542032375199")
	if err != nil || !bytes.Equal(b, h("4b09512430325781")) {
		t.Fatalf("%x %v", b, err)
	}
	if _, err := IMEI("12"); err == nil {
		t.Fatal("a short IMEISV")
	}
}

func TestIMEISV(t *testing.T) {
	// The sample SecurityModeComplete carries IMEISV 4370816125816151.
	b, err := IMEISV("4370816125816151")
	if err != nil || !bytes.Equal(b, h("4573806121856151f1")) {
		t.Fatalf("%x %v", b, err)
	}
}
