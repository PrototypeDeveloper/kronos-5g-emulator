// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PC5FlowBitRates represents ASN.1 PC5FlowBitRates.
type PC5FlowBitRates struct {
	GuaranteedFlowBitRate BitRate
	MaximumFlowBitRate    BitRate
	HasIEExtensions       bool
	IEExtensions          []PC5FlowBitRatesExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5FlowBitRates) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5FlowBitRates", v)
}

// DecodePC5FlowBitRates decodes a standalone APER value and reports the consumed octets.
func DecodePC5FlowBitRates(data []byte) (PC5FlowBitRates, int, error) {
	var v PC5FlowBitRates
	err := ngapcodec.DecodeParameter("PC5FlowBitRates", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5FlowBitRates{}, 0, err
	}
	return v, len(data), err
}
