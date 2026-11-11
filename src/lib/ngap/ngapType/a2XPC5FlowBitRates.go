// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// A2XPC5FlowBitRates represents ASN.1 A2X-PC5-FlowBitRates.
type A2XPC5FlowBitRates struct {
	A2XGuaranteedFlowBitRate BitRate
	A2XMaximumFlowBitRate    BitRate
	HasIEExtensions          bool
	IEExtensions             []A2XPC5FlowBitRatesExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5FlowBitRates) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-FlowBitRates", v)
}

// DecodeA2XPC5FlowBitRates decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5FlowBitRates(data []byte) (A2XPC5FlowBitRates, int, error) {
	var v A2XPC5FlowBitRates
	err := ngapcodec.DecodeParameter("A2X-PC5-FlowBitRates", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5FlowBitRates{}, 0, err
	}
	return v, len(data), err
}
