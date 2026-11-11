// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PLMNSupportItem represents ASN.1 PLMNSupportItem.
type PLMNSupportItem struct {
	PLMNIdentity      PLMNIdentity
	SliceSupportList  SliceSupportList
	HasIEExtensions   bool
	IEExtensions      []PLMNSupportItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PLMNSupportItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PLMNSupportItem", v)
}

// DecodePLMNSupportItem decodes a standalone APER value and reports the consumed octets.
func DecodePLMNSupportItem(data []byte) (PLMNSupportItem, int, error) {
	var v PLMNSupportItem
	err := ngapcodec.DecodeParameter("PLMNSupportItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PLMNSupportItem{}, 0, err
	}
	return v, len(data), err
}
