package security

import (
	"encoding/binary"
	"libnas/security/internal/snow3g"
)

// Byte-aligned 128-EEA1/EIA1 (NEA1/NIA1), using the SNOW 3G primitive.
// Key/IV layouts and polynomial evaluation follow TS 35.215 / TS 33.401.
func snowKey(key [16]byte) (words [4]uint32) {
	for i := range words {
		words[i] = binary.BigEndian.Uint32(key[(3-i)*4:])
	}
	return
}
func cipherSnow(key [16]byte, count uint32, bearer, direction uint8, data []byte) []byte {
	v := uint32(bearer)<<27 | uint32(direction)<<26
	stream := snow3g.GetKeyStream(snowKey(key), [4]uint32{v, count, v, count}, (len(data)+3)/4)
	out := make([]byte, len(data))
	for i, b := range data {
		out[i] = b ^ byte(stream[i/4]>>uint(24-8*(i%4)))
	}
	return out
}

// Multiplication in GF(2^64), with reduction polynomial x^64+x^4+x^3+x+1.
func snowMultiply(v, p uint64) (out uint64) {
	for range 64 {
		out ^= v & (0 - (p & 1))
		v = (v << 1) ^ (0x1b & (0 - (v >> 63)))
		p >>= 1
	}
	return
}
func integritySnow(key [16]byte, count uint32, bearer, direction uint8, data []byte) (out [4]byte) {
	fresh, d := uint32(bearer)<<27, uint32(direction)
	z := snow3g.GetKeyStream(snowKey(key), [4]uint32{fresh ^ (d << 15), count ^ (d << 31), fresh, count}, 5)
	p, q := uint64(z[0])<<32|uint64(z[1]), uint64(z[2])<<32|uint64(z[3])
	var eval uint64
	for off := 0; off < len(data); off += 8 {
		var block [8]byte
		copy(block[:], data[off:])
		eval = snowMultiply(eval^binary.BigEndian.Uint64(block[:]), p)
	}
	eval = snowMultiply(eval^(uint64(len(data))*8), q)
	binary.BigEndian.PutUint32(out[:], uint32(eval>>32)^z[4])
	return
}
