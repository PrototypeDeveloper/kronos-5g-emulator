// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EventL1LoggedMDTConfig represents ASN.1 EventL1LoggedMDTConfig.
type EventL1LoggedMDTConfig struct {
	L1Threshold       MeasurementThresholdL1LoggedMDT
	Hysteresis        Hysteresis
	TimeToTrigger     TimeToTrigger
	HasIEExtensions   bool
	IEExtensions      []EventL1LoggedMDTConfigExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventL1LoggedMDTConfig) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EventL1LoggedMDTConfig", v)
}

// DecodeEventL1LoggedMDTConfig decodes a standalone APER value and reports the consumed octets.
func DecodeEventL1LoggedMDTConfig(data []byte) (EventL1LoggedMDTConfig, int, error) {
	var v EventL1LoggedMDTConfig
	err := ngapcodec.DecodeParameter("EventL1LoggedMDTConfig", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventL1LoggedMDTConfig{}, 0, err
	}
	return v, len(data), err
}
