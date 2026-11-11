// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERadioCapabilityForPaging represents ASN.1 UERadioCapabilityForPaging.
type UERadioCapabilityForPaging struct {
	HasUERadioCapabilityForPagingOfNR    bool
	UERadioCapabilityForPagingOfNR       UERadioCapabilityForPagingOfNR
	HasUERadioCapabilityForPagingOfEUTRA bool
	UERadioCapabilityForPagingOfEUTRA    UERadioCapabilityForPagingOfEUTRA
	HasIEExtensions                      bool
	IEExtensions                         []UERadioCapabilityForPagingExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERadioCapabilityForPaging) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERadioCapabilityForPaging", v)
}

// DecodeUERadioCapabilityForPaging decodes a standalone APER value and reports the consumed octets.
func DecodeUERadioCapabilityForPaging(data []byte) (UERadioCapabilityForPaging, int, error) {
	var v UERadioCapabilityForPaging
	err := ngapcodec.DecodeParameter("UERadioCapabilityForPaging", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERadioCapabilityForPaging{}, 0, err
	}
	return v, len(data), err
}
