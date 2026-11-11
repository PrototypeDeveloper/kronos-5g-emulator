// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TSCFeedbackInformation represents ASN.1 TSCFeedbackInformation.
type TSCFeedbackInformation struct {
	BurstArrivalTimeOffset int64 `aper:"valueExt,valueLB:-640000,valueUB:640000"`
	HasAdjustedPeriodicity bool
	AdjustedPeriodicity    Periodicity
	HasIEExtensions        bool
	IEExtensions           []TSCFeedbackInformationExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TSCFeedbackInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TSCFeedbackInformation", v)
}

// DecodeTSCFeedbackInformation decodes a standalone APER value and reports the consumed octets.
func DecodeTSCFeedbackInformation(data []byte) (TSCFeedbackInformation, int, error) {
	var v TSCFeedbackInformation
	err := ngapcodec.DecodeParameter("TSCFeedbackInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return TSCFeedbackInformation{}, 0, err
	}
	return v, len(data), err
}
