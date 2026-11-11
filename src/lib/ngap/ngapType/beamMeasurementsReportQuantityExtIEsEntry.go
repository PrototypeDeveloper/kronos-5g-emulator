// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BeamMeasurementsReportQuantityExtIEsEntryPresent uint32

const (
	BeamMeasurementsReportQuantityExtIEsEntryPresentNothing BeamMeasurementsReportQuantityExtIEsEntryPresent = iota
	BeamMeasurementsReportQuantityExtIEsEntryPresentUnknown
)

type BeamMeasurementsReportQuantityExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BeamMeasurementsReportQuantityExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BeamMeasurementsReportQuantityExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BeamMeasurementsReportQuantity-ExtIEsEntry", v)
}

// DecodeBeamMeasurementsReportQuantityExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBeamMeasurementsReportQuantityExtIEsEntry(data []byte) (BeamMeasurementsReportQuantityExtIEsEntry, int, error) {
	var v BeamMeasurementsReportQuantityExtIEsEntry
	err := ngapcodec.DecodeParameter("BeamMeasurementsReportQuantity-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BeamMeasurementsReportQuantityExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
