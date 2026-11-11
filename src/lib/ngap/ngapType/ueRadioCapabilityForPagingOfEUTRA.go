// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERadioCapabilityForPagingOfEUTRA represents ASN.1 UERadioCapabilityForPagingOfEUTRA.
type UERadioCapabilityForPagingOfEUTRA struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapabilityForPagingOfEUTRA) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapabilityForPagingOfEUTRA", v)
}

// DecodeUERadioCapabilityForPagingOfEUTRA decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapabilityForPagingOfEUTRA(data []byte) (UERadioCapabilityForPagingOfEUTRA, int, error) {
	var v UERadioCapabilityForPagingOfEUTRA
	err := ngapcodec.DecodeParameter("UERadioCapabilityForPagingOfEUTRA", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapabilityForPagingOfEUTRA{}, 0, err
	}
	return v, len(data), err
}
