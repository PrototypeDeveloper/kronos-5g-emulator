// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExcessPacketDelayThresholdValue represents ASN.1 ExcessPacketDelayThresholdValue.
const (
	ExcessPacketDelayThresholdValuePresentMs0dot25 Enumerated = 0
	ExcessPacketDelayThresholdValuePresentMs0dot5  Enumerated = 1
	ExcessPacketDelayThresholdValuePresentMs1      Enumerated = 2
	ExcessPacketDelayThresholdValuePresentMs2      Enumerated = 3
	ExcessPacketDelayThresholdValuePresentMs4      Enumerated = 4
	ExcessPacketDelayThresholdValuePresentMs5      Enumerated = 5
	ExcessPacketDelayThresholdValuePresentMs10     Enumerated = 6
	ExcessPacketDelayThresholdValuePresentMs20     Enumerated = 7
	ExcessPacketDelayThresholdValuePresentMs30     Enumerated = 8
	ExcessPacketDelayThresholdValuePresentMs40     Enumerated = 9
	ExcessPacketDelayThresholdValuePresentMs50     Enumerated = 10
	ExcessPacketDelayThresholdValuePresentMs60     Enumerated = 11
	ExcessPacketDelayThresholdValuePresentMs70     Enumerated = 12
	ExcessPacketDelayThresholdValuePresentMs80     Enumerated = 13
	ExcessPacketDelayThresholdValuePresentMs90     Enumerated = 14
	ExcessPacketDelayThresholdValuePresentMs100    Enumerated = 15
	ExcessPacketDelayThresholdValuePresentMs150    Enumerated = 16
	ExcessPacketDelayThresholdValuePresentMs300    Enumerated = 17
	ExcessPacketDelayThresholdValuePresentMs500    Enumerated = 18
)

type ExcessPacketDelayThresholdValue struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:18"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExcessPacketDelayThresholdValue) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExcessPacketDelayThresholdValue", v)
}

// DecodeExcessPacketDelayThresholdValue decodes a standalone APER value and reports the consumed octets.
func DecodeExcessPacketDelayThresholdValue(data []byte) (ExcessPacketDelayThresholdValue, int, error) {
	var v ExcessPacketDelayThresholdValue
	err := ngapcodec.DecodeParameter("ExcessPacketDelayThresholdValue", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExcessPacketDelayThresholdValue{}, 0, err
	}
	return v, len(data), err
}
