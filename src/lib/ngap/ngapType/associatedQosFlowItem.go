// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedQosFlowItem represents ASN.1 AssociatedQosFlowItem.
type AssociatedQosFlowItem struct {
	QosFlowIdentifier           QosFlowIdentifier
	HasQosFlowMappingIndication bool
	QosFlowMappingIndication    AssociatedQosFlowItemQosFlowMappingIndication `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions             bool
	IEExtensions                []AssociatedQosFlowItemExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedQosFlowItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedQosFlowItem", v)
}

// DecodeAssociatedQosFlowItem decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedQosFlowItem(data []byte) (AssociatedQosFlowItem, int, error) {
	var v AssociatedQosFlowItem
	err := ngapcodec.DecodeParameter("AssociatedQosFlowItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedQosFlowItem{}, 0, err
	}
	return v, len(data), err
}
