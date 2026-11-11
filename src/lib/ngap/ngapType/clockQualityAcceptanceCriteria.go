// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockQualityAcceptanceCriteria represents ASN.1 ClockQualityAcceptanceCriteria.
type ClockQualityAcceptanceCriteria struct {
	HasSynchronisationState    bool
	SynchronisationState       BitString `aper:"sizeExt,sizeLB:8,sizeUB:8"`
	HasTraceabletoUTC          bool
	TraceabletoUTC             ClockQualityAcceptanceCriteriaTraceabletoUTC `aper:"valueExt,valueLB:0,valueUB:0"`
	HasTraceabletoGNSS         bool
	TraceabletoGNSS            ClockQualityAcceptanceCriteriaTraceabletoGNSS `aper:"valueExt,valueLB:0,valueUB:0"`
	HasClockFrequencyStability bool
	ClockFrequencyStability    BitString `aper:"sizeLB:16,sizeUB:16"`
	HasClockAccuracy           bool
	ClockAccuracy              int64 `aper:"valueExt,valueLB:1,valueUB:40000000"`
	HasParentTImeSource        bool
	ParentTImeSource           BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
	HasIEExtensions            bool
	IEExtensions               []ClockQualityAcceptanceCriteriaExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityAcceptanceCriteria) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityAcceptanceCriteria", v)
}

// DecodeClockQualityAcceptanceCriteria decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityAcceptanceCriteria(data []byte) (ClockQualityAcceptanceCriteria, int, error) {
	var v ClockQualityAcceptanceCriteria
	err := ngapcodec.DecodeParameter("ClockQualityAcceptanceCriteria", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityAcceptanceCriteria{}, 0, err
	}
	return v, len(data), err
}
