// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowSetupRequestList represents ASN.1 QosFlowSetupRequestList.
type QosFlowSetupRequestList struct {
	List []QosFlowSetupRequestItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowSetupRequestList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowSetupRequestList", v)
}

// DecodeQosFlowSetupRequestList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowSetupRequestList(data []byte) (QosFlowSetupRequestList, int, error) {
	var v QosFlowSetupRequestList
	err := ngapcodec.DecodeParameter("QosFlowSetupRequestList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowSetupRequestList{}, 0, err
	}
	return v, len(data), err
}
