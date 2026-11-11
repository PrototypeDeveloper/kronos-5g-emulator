// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowListWithCause represents ASN.1 QosFlowListWithCause.
type QosFlowListWithCause struct {
	List []QosFlowWithCauseItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowListWithCause) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowListWithCause", v)
}

// DecodeQosFlowListWithCause decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowListWithCause(data []byte) (QosFlowListWithCause, int, error) {
	var v QosFlowListWithCause
	err := ngapcodec.DecodeParameter("QosFlowListWithCause", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowListWithCause{}, 0, err
	}
	return v, len(data), err
}
