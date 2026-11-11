// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EventTrigger represents ASN.1 EventTrigger.
type EventTriggerPresent uint32

const (
	EventTriggerPresentNothing EventTriggerPresent = iota
	EventTriggerPresentOutOfCoverage
	EventTriggerPresentEventL1LoggedMDTConfig
	EventTriggerPresentChoiceExtensions
)

type EventTrigger struct {
	Present                EventTriggerPresent
	OutOfCoverage          EventTriggerOutOfCoverage `aper:"valueExt,valueLB:0,valueUB:0"`
	EventL1LoggedMDTConfig EventL1LoggedMDTConfig
	ChoiceExtensions       EventTriggerExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventTrigger) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("EventTrigger", v) }

// DecodeEventTrigger decodes a standalone APER value and reports the consumed octets.
func DecodeEventTrigger(data []byte) (EventTrigger, int, error) {
	var v EventTrigger
	err := ngapcodec.DecodeParameter("EventTrigger", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventTrigger{}, 0, err
	}
	return v, len(data), err
}
