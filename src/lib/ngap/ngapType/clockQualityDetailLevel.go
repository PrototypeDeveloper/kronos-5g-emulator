// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ClockQualityDetailLevel represents ASN.1 ClockQualityDetailLevel.
type ClockQualityDetailLevelPresent uint32

const (
	ClockQualityDetailLevelPresentNothing ClockQualityDetailLevelPresent = iota
	ClockQualityDetailLevelPresentClockQualityMetrics
	ClockQualityDetailLevelPresentClockQualityAcceptanceIndication
	ClockQualityDetailLevelPresentChoiceExtensions
)

type ClockQualityDetailLevel struct {
	Present                          ClockQualityDetailLevelPresent
	ClockQualityMetrics              struct{}
	ClockQualityAcceptanceIndication ClockQualityAcceptanceIndication
	ChoiceExtensions                 ClockQualityDetailLevelExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ClockQualityDetailLevel) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ClockQualityDetailLevel", v)
}

// DecodeClockQualityDetailLevel decodes a standalone APER value and reports the consumed octets.
func DecodeClockQualityDetailLevel(data []byte) (ClockQualityDetailLevel, int, error) {
	var v ClockQualityDetailLevel
	err := ngapcodec.DecodeParameter("ClockQualityDetailLevel", data, &v)
	if ngapcodec.IsFatal(err) {
		return ClockQualityDetailLevel{}, 0, err
	}
	return v, len(data), err
}
