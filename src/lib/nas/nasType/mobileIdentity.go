package nasType

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	MobileIdentitySUCI   uint8 = 1
	MobileIdentity5GGUTI uint8 = 2
)

// MobileIdentity is a decoded IMSI-based SUCI or 5G-GUTI. SUCI deconcealment
// belongs to UDM; protected scheme output is forwarded unchanged to AUSF.
type MobileIdentity struct {
	Type uint8
	SUCI string
	GUTI [11]byte
}

func DecodePLMN(p []byte) (mcc, mnc string, err error) {
	if len(p) != 3 {
		return "", "", fmt.Errorf("PLMN must contain three octets")
	}
	a := []byte{p[0] & 15, p[0] >> 4, p[1] & 15}
	b := []byte{p[2] & 15, p[2] >> 4}
	if p[1]>>4 != 15 {
		b = append(b, p[1]>>4)
	}
	for _, v := range append(a, b...) {
		if v > 9 {
			return "", "", fmt.Errorf("invalid PLMN digit")
		}
	}
	for _, v := range a {
		mcc += string('0' + v)
	}
	for _, v := range b {
		mnc += string('0' + v)
	}
	return mcc, mnc, nil
}
func bcdDigits(p []byte) (string, error) {
	var result strings.Builder
	for i, b := range p {
		for j, v := range []byte{b & 15, b >> 4} {
			if v == 15 && i == len(p)-1 && j == 1 {
				return result.String(), nil
			}
			if v > 9 {
				return "", fmt.Errorf("invalid BCD digit")
			}
			result.WriteByte('0' + v)
		}
	}
	return result.String(), nil
}
func DecodeMobileIdentity(p []byte) (MobileIdentity, error) {
	var id MobileIdentity
	if len(p) == 0 {
		return id, fmt.Errorf("empty mobile identity")
	}
	id.Type = p[0] & 7
	switch id.Type {
	case MobileIdentity5GGUTI:
		if len(p) != 11 || p[0] != 0xf2 {
			return id, fmt.Errorf("invalid 5G-GUTI")
		}
		if _, _, err := DecodePLMN(p[1:4]); err != nil {
			return id, err
		}
		copy(id.GUTI[:], p)
	case MobileIdentitySUCI:
		if len(p) < 9 || p[0] != 1 || p[6]&0xf0 != 0 {
			return id, fmt.Errorf("only IMSI-based SUCI is supported")
		}
		mcc, mnc, err := DecodePLMN(p[1:4])
		if err != nil {
			return id, err
		}
		// Routing indicator is 1..4 digits padded with trailing F nibbles.
		ri := ""
		filler := false
		for _, b := range p[4:6] {
			for _, v := range []byte{b & 15, b >> 4} {
				if v == 15 {
					filler = true
					continue
				}
				if v > 9 || filler {
					return id, fmt.Errorf("invalid routing indicator")
				}
				ri += string('0' + v)
			}
		}
		if ri == "" {
			return id, fmt.Errorf("empty routing indicator")
		}
		scheme := p[6] & 15
		output := hex.EncodeToString(p[8:])
		if scheme == 0 {
			if p[7] != 0 {
				return id, fmt.Errorf("null SUCI requires key ID zero")
			}
			output, err = bcdDigits(p[8:])
			if err != nil {
				return id, err
			}
			if len(mcc)+len(mnc)+len(output) > 15 || len(output) == 0 {
				return id, fmt.Errorf("invalid IMSI length")
			}
		}
		id.SUCI = fmt.Sprintf("suci-0-%s-%s-%s-%x-%d-%s", mcc, mnc, ri, scheme, p[7], output)
	default:
		return id, fmt.Errorf("unsupported mobile identity type %d", id.Type)
	}
	return id, nil
}
