// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LoggedMDTTrigger represents ASN.1 LoggedMDTTrigger.
type LoggedMDTTriggerPresent uint32

const (
	LoggedMDTTriggerPresentNothing LoggedMDTTriggerPresent = iota
	LoggedMDTTriggerPresentPeriodical
	LoggedMDTTriggerPresentEventTrigger
	LoggedMDTTriggerPresentChoiceExtensions
)

type LoggedMDTTrigger struct {
	Present          LoggedMDTTriggerPresent
	Periodical       struct{}
	EventTrigger     EventTrigger
	ChoiceExtensions LoggedMDTTriggerExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LoggedMDTTrigger) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LoggedMDTTrigger", v)
}

// DecodeLoggedMDTTrigger decodes a standalone APER value and reports the consumed octets.
func DecodeLoggedMDTTrigger(data []byte) (LoggedMDTTrigger, int, error) {
	var v LoggedMDTTrigger
	err := ngapcodec.DecodeParameter("LoggedMDTTrigger", data, &v)
	if ngapcodec.IsFatal(err) {
		return LoggedMDTTrigger{}, 0, err
	}
	return v, len(data), err
}
