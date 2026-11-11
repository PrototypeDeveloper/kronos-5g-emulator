// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAcceptedList represents ASN.1 QosFlowAcceptedList.
type QosFlowAcceptedList struct {
	List []QosFlowAcceptedItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAcceptedList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAcceptedList", v)
}

// DecodeQosFlowAcceptedList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAcceptedList(data []byte) (QosFlowAcceptedList, int, error) {
	var v QosFlowAcceptedList
	err := ngapcodec.DecodeParameter("QosFlowAcceptedList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAcceptedList{}, 0, err
	}
	return v, len(data), err
}
