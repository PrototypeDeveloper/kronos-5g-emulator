// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAddOrModifyResponseList represents ASN.1 QosFlowAddOrModifyResponseList.
type QosFlowAddOrModifyResponseList struct {
	List []QosFlowAddOrModifyResponseItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAddOrModifyResponseList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAddOrModifyResponseList", v)
}

// DecodeQosFlowAddOrModifyResponseList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAddOrModifyResponseList(data []byte) (QosFlowAddOrModifyResponseList, int, error) {
	var v QosFlowAddOrModifyResponseList
	err := ngapcodec.DecodeParameter("QosFlowAddOrModifyResponseList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAddOrModifyResponseList{}, 0, err
	}
	return v, len(data), err
}
