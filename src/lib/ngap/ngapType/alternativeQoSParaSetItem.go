// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AlternativeQoSParaSetItem represents ASN.1 AlternativeQoSParaSetItem.
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParaSetIndex AlternativeQoSParaSetIndex
	HasGuaranteedFlowBitRateDL bool
	GuaranteedFlowBitRateDL    BitRate
	HasGuaranteedFlowBitRateUL bool
	GuaranteedFlowBitRateUL    BitRate
	HasPacketDelayBudget       bool
	PacketDelayBudget          PacketDelayBudget
	HasPacketErrorRate         bool
	PacketErrorRate            PacketErrorRate
	HasIEExtensions            bool
	IEExtensions               []AlternativeQoSParaSetItemExtIEsEntry
	UnknownExtensions          []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AlternativeQoSParaSetItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AlternativeQoSParaSetItem", v)
}

// DecodeAlternativeQoSParaSetItem decodes a standalone APER value and reports the consumed octets.
func DecodeAlternativeQoSParaSetItem(data []byte) (AlternativeQoSParaSetItem, int, error) {
	var v AlternativeQoSParaSetItem
	err := ngapcodec.DecodeParameter("AlternativeQoSParaSetItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AlternativeQoSParaSetItem{}, 0, err
	}
	return v, len(data), err
}
