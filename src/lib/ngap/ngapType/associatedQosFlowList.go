// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedQosFlowList represents ASN.1 AssociatedQosFlowList.
type AssociatedQosFlowList struct {
	List []AssociatedQosFlowItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedQosFlowList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedQosFlowList", v)
}

// DecodeAssociatedQosFlowList decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedQosFlowList(data []byte) (AssociatedQosFlowList, int, error) {
	var v AssociatedQosFlowList
	err := ngapcodec.DecodeParameter("AssociatedQosFlowList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedQosFlowList{}, 0, err
	}
	return v, len(data), err
}
