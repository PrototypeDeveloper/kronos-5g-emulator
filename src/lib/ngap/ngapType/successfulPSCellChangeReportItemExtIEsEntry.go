// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SuccessfulPSCellChangeReportItemExtIEsEntryPresent uint32

const (
	SuccessfulPSCellChangeReportItemExtIEsEntryPresentNothing SuccessfulPSCellChangeReportItemExtIEsEntryPresent = iota
	SuccessfulPSCellChangeReportItemExtIEsEntryPresentUnknown
)

type SuccessfulPSCellChangeReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SuccessfulPSCellChangeReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuccessfulPSCellChangeReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuccessfulPSCellChangeReport-Item-ExtIEsEntry", v)
}

// DecodeSuccessfulPSCellChangeReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSuccessfulPSCellChangeReportItemExtIEsEntry(data []byte) (SuccessfulPSCellChangeReportItemExtIEsEntry, int, error) {
	var v SuccessfulPSCellChangeReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SuccessfulPSCellChangeReport-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuccessfulPSCellChangeReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
