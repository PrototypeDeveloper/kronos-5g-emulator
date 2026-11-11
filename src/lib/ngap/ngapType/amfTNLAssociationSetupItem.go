// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AMFTNLAssociationSetupItem represents ASN.1 AMF-TNLAssociationSetupItem.
type AMFTNLAssociationSetupItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation
	HasIEExtensions          bool
	IEExtensions             []AMFTNLAssociationSetupItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AMFTNLAssociationSetupItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AMF-TNLAssociationSetupItem", v)
}

// DecodeAMFTNLAssociationSetupItem decodes a standalone APER value and reports the consumed octets.
func DecodeAMFTNLAssociationSetupItem(data []byte) (AMFTNLAssociationSetupItem, int, error) {
	var v AMFTNLAssociationSetupItem
	err := ngapcodec.DecodeParameter("AMF-TNLAssociationSetupItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AMFTNLAssociationSetupItem{}, 0, err
	}
	return v, len(data), err
}
