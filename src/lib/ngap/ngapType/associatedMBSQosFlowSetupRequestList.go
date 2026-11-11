// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedMBSQosFlowSetupRequestList represents ASN.1 AssociatedMBSQosFlowSetupRequestList.
type AssociatedMBSQosFlowSetupRequestList struct {
	List []AssociatedMBSQosFlowSetupRequestItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedMBSQosFlowSetupRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedMBSQosFlowSetupRequestList", v)
}

// DecodeAssociatedMBSQosFlowSetupRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedMBSQosFlowSetupRequestList(data []byte) (AssociatedMBSQosFlowSetupRequestList, int, error) {
	var v AssociatedMBSQosFlowSetupRequestList
	err := ngapcodec.DecodeParameter("AssociatedMBSQosFlowSetupRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedMBSQosFlowSetupRequestList{}, 0, err
	}
	return v, len(data), err
}
