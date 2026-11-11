// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRCellIdentity represents ASN.1 NRCellIdentity.
type NRCellIdentity struct {
	Value BitString `aper:"sizeLB:36,sizeUB:36"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRCellIdentity) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRCellIdentity", v)
}

// DecodeNRCellIdentity decodes a standalone APER value and reports the consumed octets.
func DecodeNRCellIdentity(data []byte) (NRCellIdentity, int, error) {
	var v NRCellIdentity
	err := ngapcodec.DecodeParameter("NRCellIdentity", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRCellIdentity{}, 0, err
	}
	return v, len(data), err
}
