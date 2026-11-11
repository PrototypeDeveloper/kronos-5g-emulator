// Package security holds the cryptographic primitives that both roles need:
// Milenage (TS 35.206), the 5G key derivation functions
// (TS 33.501 Annex A over the KDF of TS 33.220) and the null-scheme SUCI.
// NAS ciphering and integrity themselves are in libnas.
package security

import (
	"crypto/aes"
	"errors"
)

// NewMilenage takes the subscriber key K and the operator variant OPc.
func NewMilenage(k, opc []byte) (*Milenage, error) {
	if len(k) != milenageKeyLen || len(opc) != milenageKeyLen {
		return nil, errors.New("milenage: K and OPc must be 16 octets")
	}
	b, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	m := &Milenage{block: b}
	copy(m.opc[:], opc)
	return m, nil
}

// OPc derives OPc = E_K(OP) XOR OP.
func OPc(k, op []byte) ([]byte, error) {
	if len(k) != milenageKeyLen || len(op) != milenageKeyLen {
		return nil, errors.New("milenage: K and OP must be 16 octets")
	}
	b, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	out := make([]byte, milenageKeyLen)
	b.Encrypt(out, op)
	for i := range out {
		out[i] ^= op[i]
	}
	return out, nil
}

// temp is TEMP = E_K(RAND XOR OPc).
func (m *Milenage) temp(rand []byte) [16]byte {
	var in, t [16]byte
	for i := range in {
		in[i] = rand[i] ^ m.opc[i]
	}
	m.block.Encrypt(t[:], in[:])
	return t
}

// out computes E_K(rot(x XOR OPc, r) XOR c) XOR OPc, with r in bytes and c in the last octet.
func (m *Milenage) out(x [16]byte, r int, c byte) [16]byte {
	var in, o [16]byte
	for i := range in {
		in[i] = x[(i+r)%16] ^ m.opc[(i+r)%16]
	}
	in[15] ^= c
	m.block.Encrypt(o[:], in[:])
	for i := range o {
		o[i] ^= m.opc[i]
	}
	return o
}

// F1 returns MAC-A (f1) and MAC-S (f1*) for SQN (6 octets) and AMF (2 octets).
func (m *Milenage) F1(rand, sqn, amf []byte) (macA, macS []byte) {
	t := m.temp(rand)
	var in1, x [16]byte
	copy(in1[0:], sqn[:6])
	copy(in1[6:], amf[:2])
	copy(in1[8:], sqn[:6])
	copy(in1[14:], amf[:2])
	// OUT1 = E_K(TEMP XOR rot(IN1 XOR OPc, r1) XOR c1) XOR OPc with r1 = 64 bits, c1 = 0.
	for i := range x {
		x[i] = t[i] ^ in1[(i+8)%16] ^ m.opc[(i+8)%16]
	}
	var o [16]byte
	m.block.Encrypt(o[:], x[:])
	for i := range o {
		o[i] ^= m.opc[i]
	}
	return append([]byte(nil), o[:8]...), append([]byte(nil), o[8:]...)
}

// F2345 returns RES (f2, 8 octets), CK (f3), IK (f4) and AK (f5, 6 octets).
func (m *Milenage) F2345(rand []byte) (res, ck, ik, ak []byte) {
	t := m.temp(rand)
	o2 := m.out(t, 0, 1) // r2 = 0, c2 = 1
	o3 := m.out(t, 4, 2) // r3 = 32 bits, c3 = 2
	o4 := m.out(t, 8, 4) // r4 = 64 bits, c4 = 4
	return append([]byte(nil), o2[8:]...), o3[:], o4[:], append([]byte(nil), o2[:6]...)
}

// F5Star returns AK for resynchronisation (f5*, 6 octets).
func (m *Milenage) F5Star(rand []byte) []byte {
	o5 := m.out(m.temp(rand), 12, 8) // r5 = 96 bits, c5 = 8
	return append([]byte(nil), o5[:6]...)
}
