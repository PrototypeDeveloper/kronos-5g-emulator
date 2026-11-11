// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANMeasurementConfigurationWlanRtt represents ASN.1 WLANMeasurementConfigurationWlanRtt.
const (
	WLANMeasurementConfigurationWlanRttPresentTrue Enumerated = 0
)

type WLANMeasurementConfigurationWlanRtt struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v WLANMeasurementConfigurationWlanRtt) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:WLANMeasurementConfigurationWlanRtt", v)
}

// DecodeWLANMeasurementConfigurationWlanRtt decodes a standalone APER value and reports consumed octets.
func DecodeWLANMeasurementConfigurationWlanRtt(data []byte) (WLANMeasurementConfigurationWlanRtt, int, error) {
	var v WLANMeasurementConfigurationWlanRtt
	err := ngapcodec.DecodeParameter("inline:WLANMeasurementConfigurationWlanRtt", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasurementConfigurationWlanRtt{}, 0, err
	}
	return v, len(data), err
}
