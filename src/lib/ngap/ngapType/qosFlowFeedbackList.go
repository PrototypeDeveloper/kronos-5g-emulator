// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowFeedbackList represents ASN.1 QosFlowFeedbackList.
type QosFlowFeedbackList struct {
	List []QosFlowFeedbackItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowFeedbackList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowFeedbackList", v)
}

// DecodeQosFlowFeedbackList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowFeedbackList(data []byte) (QosFlowFeedbackList, int, error) {
	var v QosFlowFeedbackList
	err := ngapcodec.DecodeParameter("QosFlowFeedbackList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowFeedbackList{}, 0, err
	}
	return v, len(data), err
}
