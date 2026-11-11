// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LoggedMDTNr represents ASN.1 LoggedMDTNr.
type LoggedMDTNr struct {
	LoggingInterval                      LoggingInterval
	LoggingDuration                      LoggingDuration
	LoggedMDTTrigger                     LoggedMDTTrigger
	HasBluetoothMeasurementConfiguration bool
	BluetoothMeasurementConfiguration    BluetoothMeasurementConfiguration
	HasWLANMeasurementConfiguration      bool
	WLANMeasurementConfiguration         WLANMeasurementConfiguration
	HasSensorMeasurementConfiguration    bool
	SensorMeasurementConfiguration       SensorMeasurementConfiguration
	HasAreaScopeOfNeighCellsList         bool
	AreaScopeOfNeighCellsList            AreaScopeOfNeighCellsList
	HasIEExtensions                      bool
	IEExtensions                         []LoggedMDTNrExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LoggedMDTNr) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("LoggedMDTNr", v) }

// DecodeLoggedMDTNr decodes a standalone APER value and reports the consumed octets.
func DecodeLoggedMDTNr(data []byte) (LoggedMDTNr, int, error) {
	var v LoggedMDTNr
	err := ngapcodec.DecodeParameter("LoggedMDTNr", data, &v)
	if ngapcodec.IsFatal(err) {
		return LoggedMDTNr{}, 0, err
	}
	return v, len(data), err
}
