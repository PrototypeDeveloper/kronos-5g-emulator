// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LastVisitedCellItemExtIEsEntryPresent uint32

const (
	LastVisitedCellItemExtIEsEntryPresentNothing LastVisitedCellItemExtIEsEntryPresent = iota
	LastVisitedCellItemExtIEsEntryPresentUnknown
)

type LastVisitedCellItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      LastVisitedCellItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LastVisitedCellItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LastVisitedCellItem-ExtIEsEntry", v)
}

// DecodeLastVisitedCellItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLastVisitedCellItemExtIEsEntry(data []byte) (LastVisitedCellItemExtIEsEntry, int, error) {
	var v LastVisitedCellItemExtIEsEntry
	err := ngapcodec.DecodeParameter("LastVisitedCellItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LastVisitedCellItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
