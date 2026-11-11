// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowToBeForwardedList represents ASN.1 QosFlowToBeForwardedList.
type QosFlowToBeForwardedList struct {
	List []QosFlowToBeForwardedItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowToBeForwardedList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowToBeForwardedList", v)
}

// DecodeQosFlowToBeForwardedList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowToBeForwardedList(data []byte) (QosFlowToBeForwardedList, int, error) {
	var v QosFlowToBeForwardedList
	err := ngapcodec.DecodeParameter("QosFlowToBeForwardedList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowToBeForwardedList{}, 0, err
	}
	return v, len(data), err
}
