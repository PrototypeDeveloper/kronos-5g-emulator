// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANMeasurementConfiguration represents ASN.1 WLANMeasurementConfiguration.
type WLANMeasurementConfiguration struct {
	WlanMeasConfig            WLANMeasConfig
	HasWlanMeasConfigNameList bool
	WlanMeasConfigNameList    WLANMeasConfigNameList
	HasWlanRssi               bool
	WlanRssi                  WLANMeasurementConfigurationWlanRssi `aper:"valueExt,valueLB:0,valueUB:0"`
	HasWlanRtt                bool
	WlanRtt                   WLANMeasurementConfigurationWlanRtt `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions           bool
	IEExtensions              []WLANMeasurementConfigurationExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANMeasurementConfiguration) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WLANMeasurementConfiguration", v)
}

// DecodeWLANMeasurementConfiguration decodes a standalone APER value and reports the consumed octets.
func DecodeWLANMeasurementConfiguration(data []byte) (WLANMeasurementConfiguration, int, error) {
	var v WLANMeasurementConfiguration
	err := ngapcodec.DecodeParameter("WLANMeasurementConfiguration", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasurementConfiguration{}, 0, err
	}
	return v, len(data), err
}
