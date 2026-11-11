// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UPTransportLayerInformationItem represents ASN.1 UPTransportLayerInformationItem.
type UPTransportLayerInformationItem struct {
	NGUUPTNLInformation UPTransportLayerInformation
	HasIEExtensions     bool
	IEExtensions        []UPTransportLayerInformationItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformationItem", v)
}

// DecodeUPTransportLayerInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformationItem(data []byte) (UPTransportLayerInformationItem, int, error) {
	var v UPTransportLayerInformationItem
	err := ngapcodec.DecodeParameter("UPTransportLayerInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformationItem{}, 0, err
	}
	return v, len(data), err
}
