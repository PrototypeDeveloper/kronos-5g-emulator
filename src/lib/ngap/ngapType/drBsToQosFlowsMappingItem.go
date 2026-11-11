// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBsToQosFlowsMappingItem represents ASN.1 DRBsToQosFlowsMappingItem.
type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID
	AssociatedQosFlowList AssociatedQosFlowList
	HasIEExtensions       bool
	IEExtensions          []DRBsToQosFlowsMappingItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsToQosFlowsMappingItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsToQosFlowsMappingItem", v)
}

// DecodeDRBsToQosFlowsMappingItem decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsToQosFlowsMappingItem(data []byte) (DRBsToQosFlowsMappingItem, int, error) {
	var v DRBsToQosFlowsMappingItem
	err := ngapcodec.DecodeParameter("DRBsToQosFlowsMappingItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsToQosFlowsMappingItem{}, 0, err
	}
	return v, len(data), err
}
