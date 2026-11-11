// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NGRANCellReportItemExtIEsEntryPresent uint32

const (
	NGRANCellReportItemExtIEsEntryPresentNothing NGRANCellReportItemExtIEsEntryPresent = iota
	NGRANCellReportItemExtIEsEntryPresentUnknown
)

type NGRANCellReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NGRANCellReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANCellReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-CellReportItem-ExtIEsEntry", v)
}

// DecodeNGRANCellReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANCellReportItemExtIEsEntry(data []byte) (NGRANCellReportItemExtIEsEntry, int, error) {
	var v NGRANCellReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("NGRAN-CellReportItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANCellReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
