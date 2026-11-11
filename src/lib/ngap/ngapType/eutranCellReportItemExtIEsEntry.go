// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRANCellReportItemExtIEsEntryPresent uint32

const (
	EUTRANCellReportItemExtIEsEntryPresentNothing EUTRANCellReportItemExtIEsEntryPresent = iota
	EUTRANCellReportItemExtIEsEntryPresentUnknown
)

type EUTRANCellReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRANCellReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANCellReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-CellReportItem-ExtIEsEntry", v)
}

// DecodeEUTRANCellReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANCellReportItemExtIEsEntry(data []byte) (EUTRANCellReportItemExtIEsEntry, int, error) {
	var v EUTRANCellReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRAN-CellReportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANCellReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
