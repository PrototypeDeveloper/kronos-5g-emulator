// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PC5QoSFlowList represents ASN.1 PC5QoSFlowList.
type PC5QoSFlowList struct {
	List []PC5QoSFlowItem `aper:"sizeLB:1,sizeUB:2048"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5QoSFlowList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5QoSFlowList", v)
}

// DecodePC5QoSFlowList decodes a standalone APER value and reports the consumed octets.
func DecodePC5QoSFlowList(data []byte) (PC5QoSFlowList, int, error) {
	var v PC5QoSFlowList
	err := ngapcodec.DecodeParameter("PC5QoSFlowList", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5QoSFlowList{}, 0, err
	}
	return v, len(data), err
}
