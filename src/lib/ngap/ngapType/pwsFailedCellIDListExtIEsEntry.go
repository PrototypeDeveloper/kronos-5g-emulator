// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PWSFailedCellIDListExtIEsEntryPresent uint32

const (
	PWSFailedCellIDListExtIEsEntryPresentNothing PWSFailedCellIDListExtIEsEntryPresent = iota
	PWSFailedCellIDListExtIEsEntryPresentUnknown
)

type PWSFailedCellIDListExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      PWSFailedCellIDListExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PWSFailedCellIDListExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PWSFailedCellIDList-ExtIEsEntry", v)
}

// DecodePWSFailedCellIDListExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePWSFailedCellIDListExtIEsEntry(data []byte) (PWSFailedCellIDListExtIEsEntry, int, error) {
	var v PWSFailedCellIDListExtIEsEntry
	err := ngapcodec.DecodeParameter("PWSFailedCellIDList-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PWSFailedCellIDListExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
