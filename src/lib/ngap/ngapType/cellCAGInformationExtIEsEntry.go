// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellCAGInformationExtIEsEntryPresent uint32

const (
	CellCAGInformationExtIEsEntryPresentNothing CellCAGInformationExtIEsEntryPresent = iota
	CellCAGInformationExtIEsEntryPresentUnknown
)

type CellCAGInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellCAGInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellCAGInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Cell-CAGInformation-ExtIEsEntry", v)
}

// DecodeCellCAGInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellCAGInformationExtIEsEntry(data []byte) (CellCAGInformationExtIEsEntry, int, error) {
	var v CellCAGInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("Cell-CAGInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellCAGInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
