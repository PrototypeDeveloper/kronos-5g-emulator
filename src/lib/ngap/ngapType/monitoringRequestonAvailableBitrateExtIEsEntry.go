// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MonitoringRequestonAvailableBitrateExtIEsEntryPresent uint32

const (
	MonitoringRequestonAvailableBitrateExtIEsEntryPresentNothing MonitoringRequestonAvailableBitrateExtIEsEntryPresent = iota
	MonitoringRequestonAvailableBitrateExtIEsEntryPresentUnknown
)

type MonitoringRequestonAvailableBitrateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MonitoringRequestonAvailableBitrateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MonitoringRequestonAvailableBitrateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MonitoringRequestonAvailableBitrate-ExtIEsEntry", v)
}

// DecodeMonitoringRequestonAvailableBitrateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMonitoringRequestonAvailableBitrateExtIEsEntry(data []byte) (MonitoringRequestonAvailableBitrateExtIEsEntry, int, error) {
	var v MonitoringRequestonAvailableBitrateExtIEsEntry
	err := ngapcodec.DecodeParameter("MonitoringRequestonAvailableBitrate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MonitoringRequestonAvailableBitrateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
