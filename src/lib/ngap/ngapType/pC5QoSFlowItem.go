// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PC5QoSFlowItem represents ASN.1 PC5QoSFlowItem.
type PC5QoSFlowItem struct {
	PQI                FiveQI
	HasPc5FlowBitRates bool
	Pc5FlowBitRates    PC5FlowBitRates
	HasRange           bool
	Range              Range
	HasIEExtensions    bool
	IEExtensions       []PC5QoSFlowItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5QoSFlowItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5QoSFlowItem", v)
}

// DecodePC5QoSFlowItem decodes a standalone APER value and reports the consumed octets.
func DecodePC5QoSFlowItem(data []byte) (PC5QoSFlowItem, int, error) {
	var v PC5QoSFlowItem
	err := ngapcodec.DecodeParameter("PC5QoSFlowItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5QoSFlowItem{}, 0, err
	}
	return v, len(data), err
}
