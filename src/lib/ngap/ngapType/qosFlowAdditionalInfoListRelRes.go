// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAdditionalInfoListRelRes represents ASN.1 QosFlowAdditionalInfoListRelRes.
type QosFlowAdditionalInfoListRelRes struct {
	List []QosFlowAdditionalInfoItemRelRes `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAdditionalInfoListRelRes) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAdditionalInfoListRelRes", v)
}

// DecodeQosFlowAdditionalInfoListRelRes decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAdditionalInfoListRelRes(data []byte) (QosFlowAdditionalInfoListRelRes, int, error) {
	var v QosFlowAdditionalInfoListRelRes
	err := ngapcodec.DecodeParameter("QosFlowAdditionalInfoListRelRes", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAdditionalInfoListRelRes{}, 0, err
	}
	return v, len(data), err
}
