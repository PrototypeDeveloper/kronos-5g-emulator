// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANTNLAssociationToRemoveItem represents ASN.1 NGRAN-TNLAssociationToRemoveItem.
type NGRANTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress       CPTransportLayerInformation
	HasTNLAssociationTransportLayerAddressAMF bool
	TNLAssociationTransportLayerAddressAMF    CPTransportLayerInformation
	HasIEExtensions                           bool
	IEExtensions                              []NGRANTNLAssociationToRemoveItemExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANTNLAssociationToRemoveItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-TNLAssociationToRemoveItem", v)
}

// DecodeNGRANTNLAssociationToRemoveItem decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANTNLAssociationToRemoveItem(data []byte) (NGRANTNLAssociationToRemoveItem, int, error) {
	var v NGRANTNLAssociationToRemoveItem
	err := ngapcodec.DecodeParameter("NGRAN-TNLAssociationToRemoveItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANTNLAssociationToRemoveItem{}, 0, err
	}
	return v, len(data), err
}
