// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedMBSQosFlowSetupRequestItem represents ASN.1 AssociatedMBSQosFlowSetupRequestItem.
type AssociatedMBSQosFlowSetupRequestItem struct {
	MBSQosFlowIdentifier               QosFlowIdentifier
	AssociatedUnicastQosFlowIdentifier QosFlowIdentifier
	HasIEExtensions                    bool
	IEExtensions                       []AssociatedMBSQosFlowSetupRequestItemExtIEsEntry
	UnknownExtensions                  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedMBSQosFlowSetupRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedMBSQosFlowSetupRequestItem", v)
}

// DecodeAssociatedMBSQosFlowSetupRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedMBSQosFlowSetupRequestItem(data []byte) (AssociatedMBSQosFlowSetupRequestItem, int, error) {
	var v AssociatedMBSQosFlowSetupRequestItem
	err := ngapcodec.DecodeParameter("AssociatedMBSQosFlowSetupRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedMBSQosFlowSetupRequestItem{}, 0, err
	}
	return v, len(data), err
}
