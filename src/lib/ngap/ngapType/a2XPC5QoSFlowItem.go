// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// A2XPC5QoSFlowItem represents ASN.1 A2X-PC5-QoS-Flow-Item.
type A2XPC5QoSFlowItem struct {
	A2XPQI                int64 `aper:"valueExt,valueLB:0,valueUB:255"`
	HasA2XPC5FlowBitRates bool
	A2XPC5FlowBitRates    A2XPC5FlowBitRates
	HasA2XRange           bool
	A2XRange              A2XPC5QoSFlowItemA2XRange `aper:"valueExt,valueLB:0,valueUB:8"`
	HasIEExtensions       bool
	IEExtensions          []A2XPC5QoSFlowItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5QoSFlowItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-QoS-Flow-Item", v)
}

// DecodeA2XPC5QoSFlowItem decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5QoSFlowItem(data []byte) (A2XPC5QoSFlowItem, int, error) {
	var v A2XPC5QoSFlowItem
	err := ngapcodec.DecodeParameter("A2X-PC5-QoS-Flow-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5QoSFlowItem{}, 0, err
	}
	return v, len(data), err
}
