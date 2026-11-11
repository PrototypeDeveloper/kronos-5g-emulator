// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UPTransportLayerInformationPairItem represents ASN.1 UPTransportLayerInformationPairItem.
type UPTransportLayerInformationPairItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation
	DLNGUUPTNLInformation UPTransportLayerInformation
	HasIEExtensions       bool
	IEExtensions          []UPTransportLayerInformationPairItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationPairItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformationPairItem", v)
}

// DecodeUPTransportLayerInformationPairItem decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationPairItem(data []byte) (UPTransportLayerInformationPairItem, int, error) {
	var v UPTransportLayerInformationPairItem
	err := ngapcodec.DecodeParameter("UPTransportLayerInformationPairItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationPairItem{}, 0, err
	}
	return v, len(data), err
}
