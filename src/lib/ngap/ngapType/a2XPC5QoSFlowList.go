// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// A2XPC5QoSFlowList represents ASN.1 A2X-PC5-QoS-Flow-List.
type A2XPC5QoSFlowList struct {
	List []A2XPC5QoSFlowItem `aper:"sizeLB:1,sizeUB:2048"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5QoSFlowList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-QoS-Flow-List", v)
}

// DecodeA2XPC5QoSFlowList decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5QoSFlowList(data []byte) (A2XPC5QoSFlowList, int, error) {
	var v A2XPC5QoSFlowList
	err := ngapcodec.DecodeParameter("A2X-PC5-QoS-Flow-List", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5QoSFlowList{}, 0, err
	}
	return v, len(data), err
}
