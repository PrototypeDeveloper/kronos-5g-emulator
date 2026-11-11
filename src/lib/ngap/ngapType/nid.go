// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NID represents ASN.1 NID.
type NID struct {
	Value BitString `aper:"sizeLB:44,sizeUB:44"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("NID", v) }

// DecodeNID decodes a standalone APER value and reports the consumed octets.
func DecodeNID(data []byte) (NID, int, error) {
	var v NID
	err := ngapcodec.DecodeParameter("NID", data, &v)
	if ngapcodec.IsFatal(err) {
		return NID{}, 0, err
	}
	return v, len(data), err
}
