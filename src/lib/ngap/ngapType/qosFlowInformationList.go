// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowInformationList represents ASN.1 QosFlowInformationList.
type QosFlowInformationList struct {
	List []QosFlowInformationItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowInformationList", v)
}

// DecodeQosFlowInformationList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowInformationList(data []byte) (QosFlowInformationList, int, error) {
	var v QosFlowInformationList
	err := ngapcodec.DecodeParameter("QosFlowInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowInformationList{}, 0, err
	}
	return v, len(data), err
}
