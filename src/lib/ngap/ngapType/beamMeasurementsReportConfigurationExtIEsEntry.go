// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BeamMeasurementsReportConfigurationExtIEsEntryPresent uint32

const (
	BeamMeasurementsReportConfigurationExtIEsEntryPresentNothing BeamMeasurementsReportConfigurationExtIEsEntryPresent = iota
	BeamMeasurementsReportConfigurationExtIEsEntryPresentUnknown
)

type BeamMeasurementsReportConfigurationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BeamMeasurementsReportConfigurationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BeamMeasurementsReportConfigurationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BeamMeasurementsReportConfiguration-ExtIEsEntry", v)
}

// DecodeBeamMeasurementsReportConfigurationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBeamMeasurementsReportConfigurationExtIEsEntry(data []byte) (BeamMeasurementsReportConfigurationExtIEsEntry, int, error) {
	var v BeamMeasurementsReportConfigurationExtIEsEntry
	err := ngapcodec.DecodeParameter("BeamMeasurementsReportConfiguration-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportConfigurationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
