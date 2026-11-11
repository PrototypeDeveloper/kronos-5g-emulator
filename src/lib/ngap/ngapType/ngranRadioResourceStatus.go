// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANRadioResourceStatus represents ASN.1 NGRAN-RadioResourceStatus.
type NGRANRadioResourceStatus struct {
	DLGBRPRBUsageForMIMO    int64 `aper:"valueLB:0,valueUB:100"`
	ULGBRPRBUsageForMIMO    int64 `aper:"valueLB:0,valueUB:100"`
	DLNonGBRPRBUsageForMIMO int64 `aper:"valueLB:0,valueUB:100"`
	ULNonGBRPRBUsageForMIMO int64 `aper:"valueLB:0,valueUB:100"`
	DLTotalPRBUsageForMIMO  int64 `aper:"valueLB:0,valueUB:100"`
	ULTotalPRBUsageForMIMO  int64 `aper:"valueLB:0,valueUB:100"`
	HasIEExtensions         bool
	IEExtensions            []NGRANRadioResourceStatusExtIEsEntry
	UnknownExtensions       []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANRadioResourceStatus) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-RadioResourceStatus", v)
}

// DecodeNGRANRadioResourceStatus decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANRadioResourceStatus(data []byte) (NGRANRadioResourceStatus, int, error) {
	var v NGRANRadioResourceStatus
	err := ngapcodec.DecodeParameter("NGRAN-RadioResourceStatus", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANRadioResourceStatus{}, 0, err
	}
	return v, len(data), err
}
