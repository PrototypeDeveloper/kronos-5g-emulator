// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedMBSQosFlowSetuporModifyRequestList represents ASN.1 AssociatedMBSQosFlowSetuporModifyRequestList.
type AssociatedMBSQosFlowSetuporModifyRequestList struct {
	List []AssociatedMBSQosFlowSetuporModifyRequestItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedMBSQosFlowSetuporModifyRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedMBSQosFlowSetuporModifyRequestList", v)
}

// DecodeAssociatedMBSQosFlowSetuporModifyRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedMBSQosFlowSetuporModifyRequestList(data []byte) (AssociatedMBSQosFlowSetuporModifyRequestList, int, error) {
	var v AssociatedMBSQosFlowSetuporModifyRequestList
	err := ngapcodec.DecodeParameter("AssociatedMBSQosFlowSetuporModifyRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedMBSQosFlowSetuporModifyRequestList{}, 0, err
	}
	return v, len(data), err
}
