// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaScopeOfNeighCellsItemExtIEsEntryPresent uint32

const (
	AreaScopeOfNeighCellsItemExtIEsEntryPresentNothing AreaScopeOfNeighCellsItemExtIEsEntryPresent = iota
	AreaScopeOfNeighCellsItemExtIEsEntryPresentUnknown
)

type AreaScopeOfNeighCellsItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AreaScopeOfNeighCellsItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfNeighCellsItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfNeighCellsItem-ExtIEsEntry", v)
}

// DecodeAreaScopeOfNeighCellsItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfNeighCellsItemExtIEsEntry(data []byte) (AreaScopeOfNeighCellsItemExtIEsEntry, int, error) {
	var v AreaScopeOfNeighCellsItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AreaScopeOfNeighCellsItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfNeighCellsItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
