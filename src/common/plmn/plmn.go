// Package plmn encodes PLMN identities (MCC/MNC) in the 3-octet BCD form
// used by NGAP and NAS (TS 24.008 §10.5.1.13, TS 38.413 §9.3.3.5).
package plmn

import "strings"

// Encode encodes MCC (3 digits) and MNC (2 or 3 digits).
func Encode(mcc, mnc string) ([]byte, bool) {
	if len(mcc) != mccDigits || (len(mnc) != mncMinDigits && len(mnc) != mncMaxDigits) {
		return nil, false
	}
	var n [mccDigits + mncMaxDigits]byte
	for i, c := range []byte(mcc + mnc) {
		if c < '0' || c > '9' {
			return nil, false
		}
		n[i] = c - '0'
	}
	mnc3 := byte(fillerNibble)
	if len(mnc) == mncMaxDigits {
		mnc3 = n[5]
	}
	return []byte{n[1]<<4 | n[0], mnc3<<4 | n[2], n[4]<<4 | n[3]}, true
}

// Decode decodes a PLMN identity. ok is false when a digit is not BCD.
func Decode(b []byte) (mcc, mnc string, ok bool) {
	if len(b) != encodedLen {
		return "", "", false
	}
	digits := []byte{b[0] & 0xf, b[0] >> 4, b[1] & 0xf, b[2] & 0xf, b[2] >> 4, b[1] >> 4}
	for i, v := range digits {
		if v > 9 && !(i == 5 && v == fillerNibble) {
			return "", "", false
		}
	}
	str := func(ds []byte) string {
		var sb strings.Builder
		for _, v := range ds {
			sb.WriteByte('0' + v)
		}
		return sb.String()
	}
	mcc = str(digits[:3])
	if digits[5] == fillerNibble {
		return mcc, str(digits[3:5]), true
	}
	return mcc, str(digits[3:6]), true
}

// String formats a PLMN identity as "MCC/MNC", or as hex when it is not valid BCD.
func String(b []byte) string {
	if mcc, mnc, ok := Decode(b); ok {
		return mcc + "/" + mnc
	}
	const hexdigits = "0123456789abcdef"
	var sb strings.Builder
	for _, c := range b {
		sb.WriteByte(hexdigits[c>>4])
		sb.WriteByte(hexdigits[c&0xf])
	}
	return sb.String()
}
