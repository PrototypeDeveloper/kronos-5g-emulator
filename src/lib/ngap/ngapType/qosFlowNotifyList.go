// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowNotifyList represents ASN.1 QosFlowNotifyList.
type QosFlowNotifyList struct {
	List []QosFlowNotifyItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowNotifyList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowNotifyList", v)
}

// DecodeQosFlowNotifyList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowNotifyList(data []byte) (QosFlowNotifyList, int, error) {
	var v QosFlowNotifyList
	err := ngapcodec.DecodeParameter("QosFlowNotifyList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowNotifyList{}, 0, err
	}
	return v, len(data), err
}
