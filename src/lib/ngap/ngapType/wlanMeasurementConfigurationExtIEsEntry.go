// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type WLANMeasurementConfigurationExtIEsEntryPresent uint32

const (
	WLANMeasurementConfigurationExtIEsEntryPresentNothing WLANMeasurementConfigurationExtIEsEntryPresent = iota
	WLANMeasurementConfigurationExtIEsEntryPresentUnknown
)

type WLANMeasurementConfigurationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      WLANMeasurementConfigurationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANMeasurementConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WLANMeasurementConfiguration-ExtIEsEntry", v)
}

// DecodeWLANMeasurementConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeWLANMeasurementConfigurationExtIEsEntry(data []byte) (WLANMeasurementConfigurationExtIEsEntry, int, error) {
	var v WLANMeasurementConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("WLANMeasurementConfiguration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasurementConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
