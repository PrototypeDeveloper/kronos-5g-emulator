// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CNTypeRestrictionsForEquivalentItem represents ASN.1 CNTypeRestrictionsForEquivalentItem.
type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity      PLMNIdentity
	CnType            CNTypeRestrictionsForEquivalentItemCnType `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions   bool
	IEExtensions      []CNTypeRestrictionsForEquivalentItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNTypeRestrictionsForEquivalentItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CNTypeRestrictionsForEquivalentItem", v)
}

// DecodeCNTypeRestrictionsForEquivalentItem decodes a standalone APER value and reports the consumed octets.
func DecodeCNTypeRestrictionsForEquivalentItem(data []byte) (CNTypeRestrictionsForEquivalentItem, int, error) {
	var v CNTypeRestrictionsForEquivalentItem
	err := ngapcodec.DecodeParameter("CNTypeRestrictionsForEquivalentItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNTypeRestrictionsForEquivalentItem{}, 0, err
	}
	return v, len(data), err
}
