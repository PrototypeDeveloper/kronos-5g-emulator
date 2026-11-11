// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationToAddItem represents ASN.1 AMF-TNLAssociationToAddItem.
type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	HasTNLAssociationUsage   bool
	TNLAssociationUsage      TNLAssociationUsage
	TNLAddressWeightFactor   TNLAddressWeightFactor
	HasIEExtensions          bool
	IEExtensions             []AMFTNLAssociationToAddItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToAddItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToAddItem", v)
}

// DecodeAMFTNLAssociationToAddItem decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToAddItem(data []byte) (AMFTNLAssociationToAddItem, int, error) {
	var v AMFTNLAssociationToAddItem
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToAddItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToAddItem{}, 0, err
	}
	return v, len(data), err
}
