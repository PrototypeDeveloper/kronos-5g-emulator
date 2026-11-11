// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaOfInterestCellItemExtIEsEntryPresent uint32

const (
	AreaOfInterestCellItemExtIEsEntryPresentNothing AreaOfInterestCellItemExtIEsEntryPresent = iota
	AreaOfInterestCellItemExtIEsEntryPresentUnknown
)

type AreaOfInterestCellItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AreaOfInterestCellItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestCellItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestCellItem-ExtIEsEntry", v)
}

// DecodeAreaOfInterestCellItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestCellItemExtIEsEntry(data []byte) (AreaOfInterestCellItemExtIEsEntry, int, error) {
	var v AreaOfInterestCellItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaOfInterestCellItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestCellItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
