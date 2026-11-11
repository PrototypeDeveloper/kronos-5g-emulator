// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MeasurementThresholdL1LoggedMDTExtIEsEntryPresent uint32

const (
	MeasurementThresholdL1LoggedMDTExtIEsEntryPresentNothing MeasurementThresholdL1LoggedMDTExtIEsEntryPresent = iota
	MeasurementThresholdL1LoggedMDTExtIEsEntryPresentUnknown
)

type MeasurementThresholdL1LoggedMDTExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      MeasurementThresholdL1LoggedMDTExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MeasurementThresholdL1LoggedMDTExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MeasurementThresholdL1LoggedMDT-ExtIEsEntry", v)
}

// DecodeMeasurementThresholdL1LoggedMDTExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMeasurementThresholdL1LoggedMDTExtIEsEntry(data []byte) (MeasurementThresholdL1LoggedMDTExtIEsEntry, int, error) {
	var v MeasurementThresholdL1LoggedMDTExtIEsEntry
	err := ngapcodec.DecodeParameter("MeasurementThresholdL1LoggedMDT-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MeasurementThresholdL1LoggedMDTExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
