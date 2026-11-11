package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// protect is the ECIES of the SUCI profiles (TS 33.501 C.3.4) on plain for
// the home network public key hn. It returns the scheme output: the
// ephemeral public key (compressed for Profile B), the ciphertext and the
// MAC tag. eph makes the ephemeral key pair.
func protect(scheme int, hn, plain []byte, eph func(ecdh.Curve) (*ecdh.PrivateKey, error)) ([]byte, error) {
	curve, hnKey, err := homeNetworkKey(scheme, hn)
	if err != nil {
		return nil, err
	}
	priv, err := eph(curve)
	if err != nil {
		return nil, err
	}
	z, err := priv.ECDH(hnKey) // X25519: the shared secret; secp256r1: its x-coordinate
	if err != nil {
		return nil, err
	}
	pub := priv.PublicKey().Bytes()
	if scheme == SchemeProfileB {
		pub = compressP256(pub)
	}
	k := x963KDF(z, pub, eciesEncKeyLen+eciesICBLen+eciesMacKeyLen)
	encKey, icb, macKey := k[:eciesEncKeyLen], k[eciesEncKeyLen:eciesEncKeyLen+eciesICBLen], k[eciesEncKeyLen+eciesICBLen:]
	block, err := aes.NewCipher(encKey)
	if err != nil {
		return nil, err
	}
	ct := make([]byte, len(plain))
	cipher.NewCTR(block, icb).XORKeyStream(ct, plain)
	mac := hmac.New(sha256.New, macKey)
	mac.Write(ct)
	out := append(append([]byte(nil), pub...), ct...)
	return append(out, mac.Sum(nil)[:eciesMacLen]...), nil
}

// homeNetworkKey reads the home network public key of a profile.
func homeNetworkKey(scheme int, hn []byte) (ecdh.Curve, *ecdh.PublicKey, error) {
	switch scheme {
	case SchemeProfileA:
		if len(hn) != x25519KeyLen {
			return nil, nil, fmt.Errorf("suci: Profile A takes a %d-octet X25519 public key (%d octets)", x25519KeyLen, len(hn))
		}
		k, err := ecdh.X25519().NewPublicKey(hn)
		return ecdh.X25519(), k, err
	case SchemeProfileB:
		full, err := uncompressP256(hn)
		if err != nil {
			return nil, nil, err
		}
		k, err := ecdh.P256().NewPublicKey(full)
		if err != nil {
			return nil, nil, fmt.Errorf("suci: the Profile B public key is not a secp256r1 point: %v", err)
		}
		return ecdh.P256(), k, nil
	}
	return nil, nil, fmt.Errorf("suci: protection scheme %d has no public key (1: Profile A, 2: Profile B)", scheme)
}

// CheckHomeNetworkKey reports what is wrong with the home network public
// key of a protection scheme (nothing for the null scheme).
func CheckHomeNetworkKey(scheme int, hn []byte) error {
	if scheme == SchemeNull {
		return nil
	}
	_, _, err := homeNetworkKey(scheme, hn)
	return err
}

// uncompressP256 returns a secp256r1 point in the uncompressed form; it
// takes either form.
func uncompressP256(b []byte) ([]byte, error) {
	switch len(b) {
	case p256FullLen:
		return b, nil
	case p256CompressedLen:
		x, y := elliptic.UnmarshalCompressed(elliptic.P256(), b)
		if x == nil {
			return nil, fmt.Errorf("suci: the Profile B public key is not a compressed secp256r1 point")
		}
		out := make([]byte, p256FullLen)
		out[0] = p256Uncompressed
		x.FillBytes(out[1 : 1+p256CoordLen])
		y.FillBytes(out[1+p256CoordLen:])
		return out, nil
	}
	return nil, fmt.Errorf("suci: Profile B takes a secp256r1 public key of %d (compressed) or %d octets (%d octets)", p256CompressedLen, p256FullLen, len(b))
}

// compressP256 returns an uncompressed secp256r1 point in the compressed form.
func compressP256(full []byte) []byte {
	out := make([]byte, p256CompressedLen)
	out[0] = p256EvenY | full[p256FullLen-1]&1
	copy(out[1:], full[1:1+p256CoordLen])
	return out
}

// x963KDF is the key derivation function of ANSI X9.63 with SHA-256: the
// hashes of z, a 32-bit counter from 1 and the shared info, up to n octets.
func x963KDF(z, sharedInfo []byte, n int) []byte {
	var out []byte
	for counter := uint32(1); len(out) < n; counter++ {
		h := sha256.New()
		h.Write(z)
		h.Write(binary.BigEndian.AppendUint32(make([]byte, 0, kdfCounterLen), counter))
		h.Write(sharedInfo)
		out = h.Sum(out)
	}
	return out[:n]
}
