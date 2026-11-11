// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationToRemoveItem represents ASN.1 AMF-TNLAssociationToRemoveItem.
type AMFTNLAssociationToRemoveItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	HasIEExtensions          bool
	IEExtensions             []AMFTNLAssociationToRemoveItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToRemoveItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToRemoveItem", v)
}

// DecodeAMFTNLAssociationToRemoveItem decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToRemoveItem(data []byte) (AMFTNLAssociationToRemoveItem, int, error) {
	var v AMFTNLAssociationToRemoveItem
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToRemoveItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToRemoveItem{}, 0, err
	}
	return v, len(data), err
}
