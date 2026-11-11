// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANMeasurementConfigurationWlanRssi represents ASN.1 WLANMeasurementConfigurationWlanRssi.
const (
	WLANMeasurementConfigurationWlanRssiPresentTrue Enumerated = 0
)

type WLANMeasurementConfigurationWlanRssi struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v WLANMeasurementConfigurationWlanRssi) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:WLANMeasurementConfigurationWlanRssi", v)
}

// DecodeWLANMeasurementConfigurationWlanRssi decodes a standalone APER value and reports consumed octets.
func DecodeWLANMeasurementConfigurationWlanRssi(data []byte) (WLANMeasurementConfigurationWlanRssi, int, error) {
	var v WLANMeasurementConfigurationWlanRssi
	err := ngapcodec.DecodeParameter("inline:WLANMeasurementConfigurationWlanRssi", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasurementConfigurationWlanRssi{}, 0, err
	}
	return v, len(data), err
}
