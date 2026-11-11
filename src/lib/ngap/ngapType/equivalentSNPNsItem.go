// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EquivalentSNPNsItem represents ASN.1 EquivalentSNPNsItem.
type EquivalentSNPNsItem struct {
	PlmnIdentity      PLMNIdentity
	NID               NID
	HasIEExtensions   bool
	IEExtensions      []EquivalentSNPNsItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EquivalentSNPNsItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EquivalentSNPNsItem", v)
}

// DecodeEquivalentSNPNsItem decodes a standalone APER value and reports the consumed octets.
func DecodeEquivalentSNPNsItem(data []byte) (EquivalentSNPNsItem, int, error) {
	var v EquivalentSNPNsItem
	err := ngapcodec.DecodeParameter("EquivalentSNPNsItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return EquivalentSNPNsItem{}, 0, err
	}
	return v, len(data), err
}
