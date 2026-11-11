// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NTNGeographicalAreaItem represents ASN.1 NTN-Geographical-Area-Item.
type NTNGeographicalAreaItemPresent uint32

const (
	NTNGeographicalAreaItemPresentNothing NTNGeographicalAreaItemPresent = iota
	NTNGeographicalAreaItemPresentCircle
	NTNGeographicalAreaItemPresentPolygon
	NTNGeographicalAreaItemPresentChoiceExtensions
)

type NTNGeographicalAreaItem struct {
	Present          NTNGeographicalAreaItemPresent
	Circle           MDTCircle
	Polygon          MDTPolygon
	ChoiceExtensions NTNGeographiclaAreaItemExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NTNGeographicalAreaItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NTN-Geographical-Area-Item", v)
}

// DecodeNTNGeographicalAreaItem decodes a standalone APER value and reports the consumed octets.
func DecodeNTNGeographicalAreaItem(data []byte) (NTNGeographicalAreaItem, int, error) {
	var v NTNGeographicalAreaItem
	err := ngapcodec.DecodeParameter("NTN-Geographical-Area-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return NTNGeographicalAreaItem{}, 0, err
	}
	return v, len(data), err
}
