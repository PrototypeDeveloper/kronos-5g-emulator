// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockQualityAcceptanceCriteriaTraceabletoGNSS represents ASN.1 ClockQualityAcceptanceCriteriaTraceabletoGNSS.
const (
	ClockQualityAcceptanceCriteriaTraceabletoGNSSPresentTrue Enumerated = 0
)

type ClockQualityAcceptanceCriteriaTraceabletoGNSS struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v ClockQualityAcceptanceCriteriaTraceabletoGNSS) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:ClockQualityAcceptanceCriteriaTraceabletoGNSS", v)
}

// DecodeClockQualityAcceptanceCriteriaTraceabletoGNSS decodes a standalone APER value and reports consumed octets.
func DecodeClockQualityAcceptanceCriteriaTraceabletoGNSS(data []byte) (ClockQualityAcceptanceCriteriaTraceabletoGNSS, int, error) {
	var v ClockQualityAcceptanceCriteriaTraceabletoGNSS
	err := ngapcodec.DecodeParameter("inline:ClockQualityAcceptanceCriteriaTraceabletoGNSS", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityAcceptanceCriteriaTraceabletoGNSS{}, 0, err
	}
	return v, len(data), err
}
