// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ImmediateMDTNr represents ASN.1 ImmediateMDTNr.
type ImmediateMDTNr struct {
	MeasurementsToActivate               MeasurementsToActivate
	HasM1Configuration                   bool
	M1Configuration                      M1Configuration
	HasM4Configuration                   bool
	M4Configuration                      M4Configuration
	HasM5Configuration                   bool
	M5Configuration                      M5Configuration
	HasM6Configuration                   bool
	M6Configuration                      M6Configuration
	HasM7Configuration                   bool
	M7Configuration                      M7Configuration
	HasBluetoothMeasurementConfiguration bool
	BluetoothMeasurementConfiguration    BluetoothMeasurementConfiguration
	HasWLANMeasurementConfiguration      bool
	WLANMeasurementConfiguration         WLANMeasurementConfiguration
	HasMDTLocationInfo                   bool
	MDTLocationInfo                      MDTLocationInfo
	HasSensorMeasurementConfiguration    bool
	SensorMeasurementConfiguration       SensorMeasurementConfiguration
	HasIEExtensions                      bool
	IEExtensions                         []ImmediateMDTNrExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ImmediateMDTNr) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ImmediateMDTNr", v)
}

// DecodeImmediateMDTNr decodes a standalone APER value and reports the consumed octets.
func DecodeImmediateMDTNr(data []byte) (ImmediateMDTNr, int, error) {
	var v ImmediateMDTNr
	err := ngapcodec.DecodeParameter("ImmediateMDTNr", data, &v)
	if ngapcodec.IsFatal(err) {
		return ImmediateMDTNr{}, 0, err
	}
	return v, len(data), err
}
