// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedMBSQosFlowSetuporModifyRequestItem represents ASN.1 AssociatedMBSQosFlowSetuporModifyRequestItem.
type AssociatedMBSQosFlowSetuporModifyRequestItem struct {
	MBSQosFlowIdentifier               QosFlowIdentifier
	AssociatedUnicastQosFlowIdentifier QosFlowIdentifier
	HasIEExtensions                    bool
	IEExtensions                       []AssociatedMBSQosFlowSetuporModifyRequestItemExtIEsEntry
	UnknownExtensions                  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedMBSQosFlowSetuporModifyRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedMBSQosFlowSetuporModifyRequestItem", v)
}

// DecodeAssociatedMBSQosFlowSetuporModifyRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedMBSQosFlowSetuporModifyRequestItem(data []byte) (AssociatedMBSQosFlowSetuporModifyRequestItem, int, error) {
	var v AssociatedMBSQosFlowSetuporModifyRequestItem
	err := ngapcodec.DecodeParameter("AssociatedMBSQosFlowSetuporModifyRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedMBSQosFlowSetuporModifyRequestItem{}, 0, err
	}
	return v, len(data), err
}
