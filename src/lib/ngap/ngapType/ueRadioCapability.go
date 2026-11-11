// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERadioCapability represents ASN.1 UERadioCapability.
type UERadioCapability struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapability) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapability", v)
}

// DecodeUERadioCapability decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapability(data []byte) (UERadioCapability, int, error) {
	var v UERadioCapability
	err := ngapcodec.DecodeParameter("UERadioCapability", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapability{}, 0, err
	}
	return v, len(data), err
}
