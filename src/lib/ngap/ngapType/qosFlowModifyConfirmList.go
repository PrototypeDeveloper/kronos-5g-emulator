// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowModifyConfirmList represents ASN.1 QosFlowModifyConfirmList.
type QosFlowModifyConfirmList struct {
	List []QosFlowModifyConfirmItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowModifyConfirmList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowModifyConfirmList", v)
}

// DecodeQosFlowModifyConfirmList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowModifyConfirmList(data []byte) (QosFlowModifyConfirmList, int, error) {
	var v QosFlowModifyConfirmList
	err := ngapcodec.DecodeParameter("QosFlowModifyConfirmList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowModifyConfirmList{}, 0, err
	}
	return v, len(data), err
}
