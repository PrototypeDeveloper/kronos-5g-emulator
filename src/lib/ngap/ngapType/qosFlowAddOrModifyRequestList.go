// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAddOrModifyRequestList represents ASN.1 QosFlowAddOrModifyRequestList.
type QosFlowAddOrModifyRequestList struct {
	List []QosFlowAddOrModifyRequestItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAddOrModifyRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAddOrModifyRequestList", v)
}

// DecodeQosFlowAddOrModifyRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAddOrModifyRequestList(data []byte) (QosFlowAddOrModifyRequestList, int, error) {
	var v QosFlowAddOrModifyRequestList
	err := ngapcodec.DecodeParameter("QosFlowAddOrModifyRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAddOrModifyRequestList{}, 0, err
	}
	return v, len(data), err
}
