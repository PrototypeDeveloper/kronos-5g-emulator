// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationToUpdateItem represents ASN.1 AMF-TNLAssociationToUpdateItem.
type AMFTNLAssociationToUpdateItem struct {
	AMFTNLAssociationAddress  CPTransportLayerInformation
	HasTNLAssociationUsage    bool
	TNLAssociationUsage       TNLAssociationUsage
	HasTNLAddressWeightFactor bool
	TNLAddressWeightFactor    TNLAddressWeightFactor
	HasIEExtensions           bool
	IEExtensions              []AMFTNLAssociationToUpdateItemExtIEsEntry
	UnknownExtensions         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationToUpdateItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationToUpdateItem", v)
}

// DecodeAMFTNLAssociationToUpdateItem decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationToUpdateItem(data []byte) (AMFTNLAssociationToUpdateItem, int, error) {
	var v AMFTNLAssociationToUpdateItem
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationToUpdateItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationToUpdateItem{}, 0, err
	}
	return v, len(data), err
}
