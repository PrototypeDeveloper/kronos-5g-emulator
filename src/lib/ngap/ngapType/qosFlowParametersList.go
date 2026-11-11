// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowParametersList represents ASN.1 QosFlowParametersList.
type QosFlowParametersList struct {
	List []QosFlowParametersItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowParametersList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowParametersList", v)
}

// DecodeQosFlowParametersList decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowParametersList(data []byte) (QosFlowParametersList, int, error) {
	var v QosFlowParametersList
	err := ngapcodec.DecodeParameter("QosFlowParametersList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowParametersList{}, 0, err
	}
	return v, len(data), err
}
