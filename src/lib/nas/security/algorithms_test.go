package security

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func unhex(s string) []byte {
	b, e := hex.DecodeString(s)
	if e != nil {
		panic(e)
	}
	return b
}
func key(s string) (k [16]byte) { copy(k[:], unhex(s)); return }

func TestCMACRFC4493(t *testing.T) {
	k := key("2b7e151628aed2a6abf7158809cf4f3c")
	input := unhex("6bc1bee22e409f96e93d7e117393172aae2d8a571e03ac9c9eb76fac45af8e5130c81c46a35ce411e5fbc1191a0a52eff69f2445df4f9b17ad2b417be66c3710")
	for _, tc := range []struct {
		n   int
		mac string
	}{{0, "bb1d6929e95937287fa37d129b756746"}, {16, "070a16b46b4d4144f79bdd9dd04a287c"}, {40, "dfa66747de9ae63030ca32611497c827"}, {64, "51f0bebf7e3b9d92fc49741779363cfe"}} {
		got := cmac(k, input[:tc.n])
		if !bytes.Equal(got[:], unhex(tc.mac)) {
			t.Fatalf("length %d: %x", tc.n, got)
		}
	}
}

// 3GPP TS 33.401 Annex C.2.2, byte-aligned EIA2/NIA2 test set 2.
func TestNIA2StandardVector(t *testing.T) {
	got, err := Integrity(NIA2, key("d3c5d592327fb11c4035c6680af8c6d1"), 0x398a59b4, 0x1a, 1, unhex("484583d5afe082ae"))
	if err != nil || !bytes.Equal(got[:], unhex("b93787e6")) {
		t.Fatalf("MAC=%x error=%v", got, err)
	}
}

// TS 33.401 Annex C.1.1: the reference vector contains 253 significant bits.
// A byte-oriented CTR computation has the same first 253 output bits.
func TestNEA2StandardVector(t *testing.T) {
	input := unhex("981ba6824c1bfb1ab485472029b71d808ce33e2cc3c0b5fc1f3de8a6dc66b1f0")
	got, err := Cipher(NEA2, key("d3c5d592327fb11c4035c6680af8c6d1"), 0x398a59b4, 0x15, 1, input)
	if err != nil {
		t.Fatal(err)
	}
	recovered, err := Cipher(NEA2, key("d3c5d592327fb11c4035c6680af8c6d1"), 0x398a59b4, 0x15, 1, got)
	if err != nil || !bytes.Equal(recovered, input) {
		t.Fatal("cipher is not reversible")
	}
	got[len(got)-1] &= 0xf8
	if !bytes.Equal(got, unhex("e9fed8a63d155304d71df20bf3e82214b20ed7dad2f233dc3c22d7bdeeed8e78")) {
		t.Fatalf("ciphertext %x", got)
	}
}
func TestUnsupportedAndNullAlgorithms(t *testing.T) {
	input := []byte{1, 2, 3}
	var k [16]byte
	for _, alg := range []uint8{3, 255} {
		if _, e := Cipher(alg, k, 0, 1, 0, input); e == nil {
			t.Fatal("cipher algorithm accepted")
		}
		if _, e := Integrity(alg, k, 0, 1, 0, input); e == nil {
			t.Fatal("integrity algorithm accepted")
		}
	}
	b, e := Cipher(NEA0, k, 0, 1, 0, input)
	if e != nil || !bytes.Equal(b, input) {
		t.Fatal(e)
	}
	b[0] = 9
	if input[0] != 1 {
		t.Fatal("null cipher aliases input")
	}
	mac, e := Integrity(NIA0, k, 0, 1, 0, input)
	if e != nil || mac != [4]byte{} {
		t.Fatal("null integrity")
	}
}
