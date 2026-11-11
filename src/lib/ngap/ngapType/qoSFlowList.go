// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoSFlowList represents ASN.1 QoSFlowList.
type QoSFlowList struct {
	List []QosFlowIdentifier `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowList) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("QoSFlowList", v) }

// DecodeQoSFlowList decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowList(data []byte) (QoSFlowList, int, error) {
	var v QoSFlowList
	err := ngapcodec.DecodeParameter("QoSFlowList", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowList{}, 0, err
	}
	return v, len(data), err
}
