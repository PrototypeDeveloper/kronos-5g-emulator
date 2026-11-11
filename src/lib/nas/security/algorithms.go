// Package security implements byte-aligned NAS confidentiality and integrity.
package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
)

const (
	NEA0 uint8 = 0
	NEA1 uint8 = 1
	NEA2 uint8 = 2
	NIA0 uint8 = 0
	NIA1 uint8 = 1
	NIA2 uint8 = 2
)

// Cipher returns a new buffer. NEA2 is AES-CTR per TS 33.401 Annex B.1.3.
func Cipher(algorithm uint8, key [16]byte, count uint32, bearer, direction uint8, data []byte) ([]byte, error) {
	if bearer > 31 || direction > 1 {
		return nil, fmt.Errorf("invalid NAS security input")
	}
	out := make([]byte, len(data))
	switch algorithm {
	case NEA0:
		copy(out, data)
	case NEA1:
		return cipherSnow(key, count, bearer, direction, data), nil
	case NEA2:
		block, err := aes.NewCipher(key[:])
		if err != nil {
			return nil, err
		}
		var iv [16]byte
		binary.BigEndian.PutUint32(iv[:4], count)
		iv[4] = bearer<<3 | direction<<2
		cipher.NewCTR(block, iv[:]).XORKeyStream(out, data)
	default:
		return nil, fmt.Errorf("unsupported NAS ciphering algorithm %d", algorithm)
	}
	return out, nil
}

// Integrity returns the most significant 32 bits of AES-CMAC for NIA2.
// data includes the NAS sequence number followed by the transmitted payload.
func Integrity(algorithm uint8, key [16]byte, count uint32, bearer, direction uint8, data []byte) ([4]byte, error) {
	var mac [4]byte
	if bearer > 31 || direction > 1 {
		return mac, fmt.Errorf("invalid NAS security input")
	}
	switch algorithm {
	case NIA0:
		return mac, nil
	case NIA1:
		return integritySnow(key, count, bearer, direction, data), nil
	case NIA2:
		input := make([]byte, 8+len(data))
		binary.BigEndian.PutUint32(input, count)
		input[4] = bearer<<3 | direction<<2
		copy(input[8:], data)
		full := cmac(key, input)
		copy(mac[:], full[:4])
		return mac, nil
	default:
		return mac, fmt.Errorf("unsupported NAS integrity algorithm %d", algorithm)
	}
}

// cmac follows RFC 4493; AES is provided by the Go standard library.
func cmac(key [16]byte, data []byte) [16]byte {
	block, _ := aes.NewCipher(key[:])
	var zero, k1, k2, state, last [16]byte
	block.Encrypt(k1[:], zero[:])
	double := func(v *[16]byte) {
		carry := v[0] >> 7
		for i := 0; i < 15; i++ {
			v[i] = v[i]<<1 | v[i+1]>>7
		}
		v[15] <<= 1
		v[15] ^= 0x87 & (0 - carry)
	}
	double(&k1)
	k2 = k1
	double(&k2)
	n := (len(data) + 15) / 16
	if n == 0 {
		n = 1
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < 16; j++ {
			state[j] ^= data[i*16+j]
		}
		block.Encrypt(state[:], state[:])
	}
	tail := data[(n-1)*16:]
	copy(last[:], tail)
	if len(tail) == 16 {
		for i := range last {
			last[i] ^= k1[i]
		}
	} else {
		last[len(tail)] = 0x80
		for i := range last {
			last[i] ^= k2[i]
		}
	}
	for i := range last {
		last[i] ^= state[i]
	}
	block.Encrypt(last[:], last[:])
	return last
}
