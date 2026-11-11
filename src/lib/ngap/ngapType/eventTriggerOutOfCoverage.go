// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EventTriggerOutOfCoverage represents ASN.1 EventTriggerOutOfCoverage.
const (
	EventTriggerOutOfCoveragePresentTrue Enumerated = 0
)

type EventTriggerOutOfCoverage struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v EventTriggerOutOfCoverage) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:EventTriggerOutOfCoverage", v)
}

// DecodeEventTriggerOutOfCoverage decodes a standalone APER value and reports consumed octets.
func DecodeEventTriggerOutOfCoverage(data []byte) (EventTriggerOutOfCoverage, int, error) {
	var v EventTriggerOutOfCoverage
	err := ngapcodec.DecodeParameter("inline:EventTriggerOutOfCoverage", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventTriggerOutOfCoverage{}, 0, err
	}
	return v, len(data), err
}
