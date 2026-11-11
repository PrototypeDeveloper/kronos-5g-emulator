// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaScopeOfNeighCellsItem represents ASN.1 AreaScopeOfNeighCellsItem.
type AreaScopeOfNeighCellsItem struct {
	NrFrequencyInfo   NRFrequencyInfo
	HasPciListForMDT  bool
	PciListForMDT     PCIListForMDT
	HasIEExtensions   bool
	IEExtensions      []AreaScopeOfNeighCellsItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfNeighCellsItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfNeighCellsItem", v)
}

// DecodeAreaScopeOfNeighCellsItem decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfNeighCellsItem(data []byte) (AreaScopeOfNeighCellsItem, int, error) {
	var v AreaScopeOfNeighCellsItem
	err := ngapcodec.DecodeParameter("AreaScopeOfNeighCellsItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfNeighCellsItem{}, 0, err
	}
	return v, len(data), err
}
