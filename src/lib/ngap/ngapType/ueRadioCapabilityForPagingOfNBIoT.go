// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERadioCapabilityForPagingOfNBIoT represents ASN.1 UERadioCapabilityForPagingOfNB-IoT.
type UERadioCapabilityForPagingOfNBIoT struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapabilityForPagingOfNBIoT) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapabilityForPagingOfNB-IoT", v)
}

// DecodeUERadioCapabilityForPagingOfNBIoT decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapabilityForPagingOfNBIoT(data []byte) (UERadioCapabilityForPagingOfNBIoT, int, error) {
	var v UERadioCapabilityForPagingOfNBIoT
	err := ngapcodec.DecodeParameter("UERadioCapabilityForPagingOfNB-IoT", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapabilityForPagingOfNBIoT{}, 0, err
	}
	return v, len(data), err
}
