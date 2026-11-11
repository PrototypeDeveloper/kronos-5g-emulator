// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExcessPacketDelayThresholdConfiguration represents ASN.1 ExcessPacketDelayThresholdConfiguration.
type ExcessPacketDelayThresholdConfiguration struct {
	List []ExcessPacketDelayThresholdItem `aper:"sizeLB:1,sizeUB:255"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExcessPacketDelayThresholdConfiguration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExcessPacketDelayThresholdConfiguration", v)
}

// DecodeExcessPacketDelayThresholdConfiguration decodes a standalone APER value and reports the consumed octets.
func DecodeExcessPacketDelayThresholdConfiguration(data []byte) (ExcessPacketDelayThresholdConfiguration, int, error) {
	var v ExcessPacketDelayThresholdConfiguration
	err := ngapcodec.DecodeParameter("ExcessPacketDelayThresholdConfiguration", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExcessPacketDelayThresholdConfiguration{}, 0, err
	}
	return v, len(data), err
}
