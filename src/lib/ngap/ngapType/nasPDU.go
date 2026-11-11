// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NASPDU represents ASN.1 NAS-PDU.
type NASPDU struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NASPDU) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NAS-PDU", v) }

// DecodeNASPDU decodes a standalone APER value and reports the consumed octets.
func DecodeNASPDU(data []byte) (NASPDU, int, error) {
	var v NASPDU
	err := ngapcodec.DecodeParameter("NAS-PDU", data, &v)
	if ngapcodec.IsFatal(err) {
		return NASPDU{}, 0, err
	}
	return v, len(data), err
}
