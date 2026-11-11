// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAdditionalInfoListRelCom represents ASN.1 QosFlowAdditionalInfoListRelCom.
type QosFlowAdditionalInfoListRelCom struct {
	List []QosFlowAdditionalInfoItemRelCom `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAdditionalInfoListRelCom) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAdditionalInfoListRelCom", v)
}

// DecodeQosFlowAdditionalInfoListRelCom decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAdditionalInfoListRelCom(data []byte) (QosFlowAdditionalInfoListRelCom, int, error) {
	var v QosFlowAdditionalInfoListRelCom
	err := ngapcodec.DecodeParameter("QosFlowAdditionalInfoListRelCom", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAdditionalInfoListRelCom{}, 0, err
	}
	return v, len(data), err
}
