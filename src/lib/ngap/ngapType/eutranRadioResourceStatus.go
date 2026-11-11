// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANRadioResourceStatus represents ASN.1 EUTRAN-RadioResourceStatus.
type EUTRANRadioResourceStatus struct {
	DLGBRPRBUsage                int64 `aper:"valueLB:0,valueUB:100"`
	ULGBRPRBUsage                int64 `aper:"valueLB:0,valueUB:100"`
	DLNonGBRPRBUsage             int64 `aper:"valueLB:0,valueUB:100"`
	ULNonGBRPRBUsage             int64 `aper:"valueLB:0,valueUB:100"`
	DLTotalPRBUsage              int64 `aper:"valueLB:0,valueUB:100"`
	ULTotalPRBUsage              int64 `aper:"valueLB:0,valueUB:100"`
	HasDLSchedulingPDCCHCCEUsage bool
	DLSchedulingPDCCHCCEUsage    int64 `aper:"valueLB:0,valueUB:100"`
	HasULSchedulingPDCCHCCEUsage bool
	ULSchedulingPDCCHCCEUsage    int64 `aper:"valueLB:0,valueUB:100"`
	HasIEExtensions              bool
	IEExtensions                 []EUTRANRadioResourceStatusExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANRadioResourceStatus) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-RadioResourceStatus", v)
}

// DecodeEUTRANRadioResourceStatus decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANRadioResourceStatus(data []byte) (EUTRANRadioResourceStatus, int, error) {
	var v EUTRANRadioResourceStatus
	err := ngapcodec.DecodeParameter("EUTRAN-RadioResourceStatus", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANRadioResourceStatus{}, 0, err
	}
	return v, len(data), err
}
