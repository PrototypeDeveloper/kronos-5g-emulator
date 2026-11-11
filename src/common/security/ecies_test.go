package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/hmac"
	"crypto/sha256"
	"testing"
)

// The SUCI test data of TS 33.501 Annex C.4 (the MSIN 001002086 of IMSI
// 208 93 001002086), with the home network and ephemeral private keys.
var suciVectors = []struct {
	name                  string
	scheme                int
	hnPriv, hnPub, ephKey string
	output                string // the scheme output: ephemeral public key, ciphertext, MAC tag
}{
	{"Profile A", SchemeProfileA,
		"c53c22208b61860b06c62e5406a7b330c2b577aa5558981510d128247d38bd1d",
		"5a8d38864820197c3394b92613b20b91633cbd897119273bf8e4a6f4eec0a650",
		"c80949f13ebe61af4ebdbd293ea4f942696b9e815d7e8f0096bbf6ed7de62256",
		"b2e92f836055a255837debf850b528997ce0201cb82adfe4be1f587d07d8457d" + "cb02352410" + "cddd9e730ef3fa87"},
	{"Profile B", SchemeProfileB,
		"f1ab1074477ebcc7f554ea1c5fc368b1616730155e0041ac447d6301975fecda",
		"0272da71976234ce833a6907425867b82e074d44ef907dfb4b3e21c1c2256ebcd1",
		"99798858a1dc6a2c68637149a4b1dbfd1fdff5addd62a2142f06699ed7602529",
		"039aab8376597021e855679a9778ea0b67396e68c66df32c0f41e9acca2da9b9d1" + "46a33fc271" + "6ac7dae96aa30a4d"},
}

// curveOf returns the curve of a profile.
func curveOf(scheme int) ecdh.Curve {
	if scheme == SchemeProfileA {
		return ecdh.X25519()
	}
	return ecdh.P256()
}

// deconceal is the home network side: the MSIN BCD from a scheme output.
func deconceal(t *testing.T, scheme int, hnPriv, so []byte) []byte {
	t.Helper()
	priv, err := curveOf(scheme).NewPrivateKey(hnPriv)
	if err != nil {
		t.Fatal(err)
	}
	n := x25519KeyLen
	if scheme == SchemeProfileB {
		n = p256CompressedLen
	}
	pub, ct, tag := so[:n], so[n:len(so)-eciesMacLen], so[len(so)-eciesMacLen:]
	full := pub
	if scheme == SchemeProfileB {
		if full, err = uncompressP256(pub); err != nil {
			t.Fatal(err)
		}
	}
	ephPub, err := curveOf(scheme).NewPublicKey(full)
	if err != nil {
		t.Fatal(err)
	}
	z, err := priv.ECDH(ephPub)
	if err != nil {
		t.Fatal(err)
	}
	k := x963KDF(z, pub, eciesEncKeyLen+eciesICBLen+eciesMacKeyLen)
	mac := hmac.New(sha256.New, k[eciesEncKeyLen+eciesICBLen:])
	mac.Write(ct)
	if !bytes.Equal(mac.Sum(nil)[:eciesMacLen], tag) {
		t.Fatal("the MAC tag does not match")
	}
	block, _ := aes.NewCipher(k[:eciesEncKeyLen])
	plain := make([]byte, len(ct))
	cipher.NewCTR(block, k[eciesEncKeyLen:eciesEncKeyLen+eciesICBLen]).XORKeyStream(plain, ct)
	return plain
}

func TestSUCIProfiles(t *testing.T) {
	for _, v := range suciVectors {
		// The keys of the test data belong together.
		hn, err := curveOf(v.scheme).NewPrivateKey(h(v.hnPriv))
		if err != nil {
			t.Fatal(err)
		}
		pub := hn.PublicKey().Bytes()
		if v.scheme == SchemeProfileB {
			pub = compressP256(pub)
		}
		if !bytes.Equal(pub, h(v.hnPub)) {
			t.Fatalf("%s: the home network public key is %x", v.name, pub)
		}
		// With the ephemeral key of the test data, the output is the test data's.
		so, err := protect(v.scheme, h(v.hnPub), msinBCD("001002086"), func(c ecdh.Curve) (*ecdh.PrivateKey, error) { return c.NewPrivateKey(h(v.ephKey)) })
		if err != nil || !bytes.Equal(so, h(v.output)) {
			t.Fatalf("%s: scheme output %x, %v\nwant %s", v.name, so, err, v.output)
		}
		// A SUCI with a new ephemeral key: the home network gets the MSIN back.
		p := SUCIParams{MCC: "208", MNC: "93", RoutingIndicator: "0000", Scheme: v.scheme, KeyID: 7, HomeNetworkKey: h(v.hnPub)}
		a, err := SUCI(p, "0000000001")
		if err != nil {
			t.Fatal(err)
		}
		b, _ := SUCI(p, "0000000001")
		head := h("0102f8390000" + "0" + string(rune('0'+v.scheme)) + "07")
		if !bytes.HasPrefix(a, head) || bytes.Equal(a, b) {
			t.Fatalf("%s: SUCI %x (another %x)", v.name, a, b)
		}
		if got := deconceal(t, v.scheme, h(v.hnPriv), a[len(head):]); !bytes.Equal(got, msinBCD("0000000001")) {
			t.Fatalf("%s: deconcealed %x", v.name, got)
		}
		// The uncompressed form of a Profile B key works the same.
		if v.scheme == SchemeProfileB {
			if err := CheckHomeNetworkKey(v.scheme, hn.PublicKey().Bytes()); err != nil {
				t.Fatal(err)
			}
		}
	}
	for scheme, key := range map[int]string{SchemeProfileA: "0102", SchemeProfileB: "05" + "72da71976234ce833a6907425867b82e074d44ef907dfb4b3e21c1c2256ebcd1", 3: ""} {
		if err := CheckHomeNetworkKey(scheme, h(key)); err == nil {
			t.Errorf("scheme %d key %s: no error", scheme, key)
		}
	}
	if err := CheckHomeNetworkKey(SchemeNull, nil); err != nil {
		t.Fatal(err)
	}
}
