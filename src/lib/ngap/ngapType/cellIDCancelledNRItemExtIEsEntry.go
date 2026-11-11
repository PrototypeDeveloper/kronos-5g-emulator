// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CellIDCancelledNRItemExtIEsEntryPresent uint32

const (
	CellIDCancelledNRItemExtIEsEntryPresentNothing CellIDCancelledNRItemExtIEsEntryPresent = iota
	CellIDCancelledNRItemExtIEsEntryPresentUnknown
)

type CellIDCancelledNRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CellIDCancelledNRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CellIDCancelledNRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CellIDCancelledNR-Item-ExtIEsEntry", v)
}

// DecodeCellIDCancelledNRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCellIDCancelledNRItemExtIEsEntry(data []byte) (CellIDCancelledNRItemExtIEsEntry, int, error) {
	var v CellIDCancelledNRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CellIDCancelledNR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CellIDCancelledNRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
