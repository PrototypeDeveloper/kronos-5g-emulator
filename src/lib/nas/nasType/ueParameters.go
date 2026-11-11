package nasType

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// EncodeNullSUCI constructs the TS 24.501 IMSI-based, null-scheme identity.
func EncodeNullSUCI(plmn []byte, imsi, routing string) ([]byte, error) {
	mcc, mnc, err := DecodePLMN(plmn)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(imsi, mcc+mnc) || len(imsi) > 15 || len(imsi) <= len(mcc+mnc) || len(routing) < 1 || len(routing) > 4 {
		return nil, fmt.Errorf("invalid IMSI/routing indicator")
	}
	bcd := func(s string) ([]byte, error) {
		var out []byte
		for i := 0; i < len(s); i += 2 {
			lo := s[i] - '0'
			if lo > 9 {
				return nil, fmt.Errorf("invalid decimal identity")
			}
			hi := byte(15)
			if i+1 < len(s) {
				hi = s[i+1] - '0'
				if hi > 9 {
					return nil, fmt.Errorf("invalid decimal identity")
				}
			}
			out = append(out, lo|hi<<4)
		}
		return out, nil
	}
	ri, err := bcd(routing)
	if err != nil {
		return nil, err
	}
	if len(ri) == 1 {
		ri = append(ri, 255)
	}
	msin, err := bcd(imsi[len(mcc+mnc):])
	if err != nil {
		return nil, err
	}
	out := append([]byte{1}, plmn...)
	out = append(out, ri...)
	out = append(out, 0, 0)
	out = append(out, msin...)
	return out, nil
}
func EncodeDNN(name string) ([]byte, error) {
	var out []byte
	for _, label := range strings.Split(name, ".") {
		if len(label) < 1 || len(label) > 63 {
			return nil, fmt.Errorf("invalid DNN label")
		}
		for _, c := range label {
			if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-') {
				return nil, fmt.Errorf("invalid DNN character")
			}
		}
		out = append(out, byte(len(label)))
		out = append(out, label...)
	}
	if len(out) > 100 {
		return nil, fmt.Errorf("DNN too long")
	}
	return out, nil
}
func EncodeSNSSAI(sst uint8, sd string) ([]byte, error) {
	out := []byte{sst}
	if sd != "" {
		v, err := hex.DecodeString(sd)
		if err != nil || len(v) != 3 {
			return nil, fmt.Errorf("invalid SD")
		}
		out = append(out, v...)
	}
	return out, nil
}

// DecodePDUAddress separates session type, IPv6 interface ID, and IPv4 address.
func DecodePDUAddress(b []byte) (kind uint8, ipv4, interfaceID []byte, err error) {
	if len(b) < 1 {
		return 0, nil, nil, fmt.Errorf("empty PDU address")
	}
	if b[0]&0xf8 != 0 {
		return 0, nil, nil, fmt.Errorf("extended PDU address is not supported")
	}
	kind = b[0] & 7
	switch kind {
	case 1:
		if len(b) != 5 {
			err = fmt.Errorf("invalid IPv4 PDU address")
		} else {
			ipv4 = append([]byte(nil), b[1:]...)
		}
	case 2:
		if len(b) != 9 {
			err = fmt.Errorf("invalid IPv6 PDU address")
		} else {
			interfaceID = append([]byte(nil), b[1:9]...)
		}
	case 3:
		if len(b) != 13 {
			err = fmt.Errorf("invalid dual-stack PDU address")
		} else {
			interfaceID = append([]byte(nil), b[1:9]...)
			ipv4 = append([]byte(nil), b[9:13]...)
		}
	default:
		err = fmt.Errorf("unsupported PDU address type")
	}
	return
}
