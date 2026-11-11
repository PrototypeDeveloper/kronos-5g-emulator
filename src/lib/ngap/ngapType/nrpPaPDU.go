// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPPaPDU represents ASN.1 NRPPa-PDU.
type NRPPaPDU struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPPaPDU) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NRPPa-PDU", v) }

// DecodeNRPPaPDU decodes a standalone APER value and reports the consumed octets.
func DecodeNRPPaPDU(data []byte) (NRPPaPDU, int, error) {
	var v NRPPaPDU
	err := ngapcodec.DecodeParameter("NRPPa-PDU", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPPaPDU{}, 0, err
	}
	return v, len(data), err
}
