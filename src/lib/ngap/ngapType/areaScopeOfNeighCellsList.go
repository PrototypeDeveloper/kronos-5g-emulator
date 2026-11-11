// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaScopeOfNeighCellsList represents ASN.1 AreaScopeOfNeighCellsList.
type AreaScopeOfNeighCellsList struct {
	List []AreaScopeOfNeighCellsItem `aper:"sizeLB:1,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaScopeOfNeighCellsList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaScopeOfNeighCellsList", v)
}

// DecodeAreaScopeOfNeighCellsList decodes a standalone APER value and reports the consumed octets.
func DecodeAreaScopeOfNeighCellsList(data []byte) (AreaScopeOfNeighCellsList, int, error) {
	var v AreaScopeOfNeighCellsList
	err := ngapcodec.DecodeParameter("AreaScopeOfNeighCellsList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaScopeOfNeighCellsList{}, 0, err
	}
	return v, len(data), err
}
