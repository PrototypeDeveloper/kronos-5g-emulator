package security

import (
	"fmt"

	"kronos/common/plmn"
)

// NullSUCI encodes the 5GS mobile identity value of an IMSI-based SUCI with
// the null protection scheme (TS 24.501 §9.11.3.4, TS 33.501 Annex C): the
// MSIN goes out in clear BCD and the home network public key ID is 0.
func NullSUCI(mcc, mnc, routingIndicator, msin string) ([]byte, error) {
	p, ok := plmn.Encode(mcc, mnc)
	if !ok {
		return nil, fmt.Errorf("suci: invalid PLMN %s/%s", mcc, mnc)
	}
	if routingIndicator == "" {
		routingIndicator = "0"
	}
	if len(routingIndicator) > maxRoutingIndicatorDigits || !digits(routingIndicator) || msin == "" || !digits(msin) {
		return nil, fmt.Errorf("suci: invalid routing indicator %q or MSIN %q", routingIndicator, msin)
	}
	out := []byte{identityTypeSUCI} // SUPI format IMSI, type of identity SUCI
	out = append(out, p...)
	out = append(out, bcd(routingIndicator, maxRoutingIndicatorDigits)...)
	out = append(out, 0x00, 0x00) // protection scheme null, home network public key ID 0
	return append(out, bcd(msin, len(msin)+len(msin)%2)...), nil
}

// digits reports whether every character of s is a decimal digit.
func digits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// bcd packs digits two per octet, low nibble first, padding to n nibbles with F.
func bcd(s string, n int) []byte {
	nib := make([]byte, n)
	for i := range nib {
		nib[i] = bcdFiller
		if i < len(s) {
			nib[i] = s[i] - '0'
		}
	}
	out := make([]byte, n/2)
	for i := range out {
		out[i] = nib[2*i+1]<<4 | nib[2*i]
	}
	return out
}

// IMEISV encodes a 16-digit IMEISV as a 5GS mobile identity value (type 5).
func IMEISV(s string) ([]byte, error) {
	if len(s) != imeisvDigits || !digits(s) {
		return nil, fmt.Errorf("imeisv: %q is not 16 digits", s)
	}
	out := []byte{(s[0]-'0')<<4 | identityTypeIMEISV} // even number of digits, type IMEISV
	return append(out, bcd(s[1:], 16)...), nil
}

// IMEI encodes the IMEI of a 16-digit IMEISV as a 5GS mobile identity value
// (type 3): its first 14 digits and the check digit (TS 23.003 annex B).
func IMEI(imeisv string) ([]byte, error) {
	if len(imeisv) != imeisvDigits || !digits(imeisv) {
		return nil, fmt.Errorf("imeisv: %q is not 16 digits", imeisv)
	}
	body := imeisv[:imeiBodyDigits]
	sum := 0
	for i := range body {
		d := int(body[i] - '0')
		if i%luhnWeight == 1 { // every second digit from the left is doubled
			d *= luhnWeight
			d = d/decimalBase + d%decimalBase
		}
		sum += d
	}
	s := body + string(rune('0'+(decimalBase-sum%decimalBase)%decimalBase))
	out := []byte{(s[0]-'0')<<4 | identityTypeIMEI}
	return append(out, bcd(s[1:], imeiBodyDigits)...), nil
}
