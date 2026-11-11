// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TNLAssociationItem represents ASN.1 TNLAssociationItem.
type TNLAssociationItem struct {
	TNLAssociationAddress CPTransportLayerInformation
	Cause                 Cause
	HasIEExtensions       bool
	IEExtensions          []TNLAssociationItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TNLAssociationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TNLAssociationItem", v)
}

// DecodeTNLAssociationItem decodes a standalone APER value and reports the consumed octets.
func DecodeTNLAssociationItem(data []byte) (TNLAssociationItem, int, error) {
	var v TNLAssociationItem
	err := ngapcodec.DecodeParameter("TNLAssociationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return TNLAssociationItem{}, 0, err
	}
	return v, len(data), err
}
