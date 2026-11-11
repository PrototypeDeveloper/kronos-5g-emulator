// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERadioCapabilityForPagingOfNR represents ASN.1 UERadioCapabilityForPagingOfNR.
type UERadioCapabilityForPagingOfNR struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapabilityForPagingOfNR) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapabilityForPagingOfNR", v)
}

// DecodeUERadioCapabilityForPagingOfNR decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapabilityForPagingOfNR(data []byte) (UERadioCapabilityForPagingOfNR, int, error) {
	var v UERadioCapabilityForPagingOfNR
	err := ngapcodec.DecodeParameter("UERadioCapabilityForPagingOfNR", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapabilityForPagingOfNR{}, 0, err
	}
	return v, len(data), err
}
