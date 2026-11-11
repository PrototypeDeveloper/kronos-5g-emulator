// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TimeToTrigger represents ASN.1 TimeToTrigger.
const (
	TimeToTriggerPresentMs0    Enumerated = 0
	TimeToTriggerPresentMs40   Enumerated = 1
	TimeToTriggerPresentMs64   Enumerated = 2
	TimeToTriggerPresentMs80   Enumerated = 3
	TimeToTriggerPresentMs100  Enumerated = 4
	TimeToTriggerPresentMs128  Enumerated = 5
	TimeToTriggerPresentMs160  Enumerated = 6
	TimeToTriggerPresentMs256  Enumerated = 7
	TimeToTriggerPresentMs320  Enumerated = 8
	TimeToTriggerPresentMs480  Enumerated = 9
	TimeToTriggerPresentMs512  Enumerated = 10
	TimeToTriggerPresentMs640  Enumerated = 11
	TimeToTriggerPresentMs1024 Enumerated = 12
	TimeToTriggerPresentMs1280 Enumerated = 13
	TimeToTriggerPresentMs2560 Enumerated = 14
	TimeToTriggerPresentMs5120 Enumerated = 15
)

type TimeToTrigger struct {
	Value Enumerated `aper:"valueLB:0,valueUB:15"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TimeToTrigger) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TimeToTrigger", v) }

// DecodeTimeToTrigger decodes a standalone APER value and reports the consumed octets.
func DecodeTimeToTrigger(data []byte) (TimeToTrigger, int, error) {
	var v TimeToTrigger
	err := ngapcodec.DecodeParameter("TimeToTrigger", data, &v)
	if ngapcodec.IsFatal(err) {
		return TimeToTrigger{}, 0, err
	}
	return v, len(data), err
}
