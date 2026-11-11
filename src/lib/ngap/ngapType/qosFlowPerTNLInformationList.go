// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowPerTNLInformationList represents ASN.1 QosFlowPerTNLInformationList.
type QosFlowPerTNLInformationList struct {
	List []QosFlowPerTNLInformationItem `aper:"sizeLB:1,sizeUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowPerTNLInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowPerTNLInformationList", v)
}

// DecodeQosFlowPerTNLInformationList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowPerTNLInformationList(data []byte) (QosFlowPerTNLInformationList, int, error) {
	var v QosFlowPerTNLInformationList
	err := ngapcodec.DecodeParameter("QosFlowPerTNLInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowPerTNLInformationList{}, 0, err
	}
	return v, len(data), err
}
