// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockQualityAcceptanceCriteriaTraceabletoUTC represents ASN.1 ClockQualityAcceptanceCriteriaTraceabletoUTC.
const (
	ClockQualityAcceptanceCriteriaTraceabletoUTCPresentTrue Enumerated = 0
)

type ClockQualityAcceptanceCriteriaTraceabletoUTC struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v ClockQualityAcceptanceCriteriaTraceabletoUTC) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:ClockQualityAcceptanceCriteriaTraceabletoUTC", v)
}

// DecodeClockQualityAcceptanceCriteriaTraceabletoUTC decodes a standalone APER value and reports consumed octets.
func DecodeClockQualityAcceptanceCriteriaTraceabletoUTC(data []byte) (ClockQualityAcceptanceCriteriaTraceabletoUTC, int, error) {
	var v ClockQualityAcceptanceCriteriaTraceabletoUTC
	err := ngapcodec.DecodeParameter("inline:ClockQualityAcceptanceCriteriaTraceabletoUTC", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityAcceptanceCriteriaTraceabletoUTC{}, 0, err
	}
	return v, len(data), err
}
