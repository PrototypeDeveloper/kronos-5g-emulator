// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowListWithDataForwarding represents ASN.1 QosFlowListWithDataForwarding.
type QosFlowListWithDataForwarding struct {
	List []QosFlowItemWithDataForwarding `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowListWithDataForwarding) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowListWithDataForwarding", v)
}

// DecodeQosFlowListWithDataForwarding decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowListWithDataForwarding(data []byte) (QosFlowListWithDataForwarding, int, error) {
	var v QosFlowListWithDataForwarding
	err := ngapcodec.DecodeParameter("QosFlowListWithDataForwarding", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowListWithDataForwarding{}, 0, err
	}
	return v, len(data), err
}
