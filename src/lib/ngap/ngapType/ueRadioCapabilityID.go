// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERadioCapabilityID represents ASN.1 UERadioCapabilityID.
type UERadioCapabilityID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapabilityID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapabilityID", v)
}

// DecodeUERadioCapabilityID decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapabilityID(data []byte) (UERadioCapabilityID, int, error) {
	var v UERadioCapabilityID
	err := ngapcodec.DecodeParameter("UERadioCapabilityID", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapabilityID{}, 0, err
	}
	return v, len(data), err
}
