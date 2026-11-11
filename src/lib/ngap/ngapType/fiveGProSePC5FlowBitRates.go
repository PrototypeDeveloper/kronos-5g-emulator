// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// FiveGProSePC5FlowBitRates represents ASN.1 FiveGProSePC5FlowBitRates.
type FiveGProSePC5FlowBitRates struct {
	FiveGproSeguaranteedFlowBitRate BitRate
	FiveGproSemaximumFlowBitRate    BitRate
	HasIEExtensions                 bool
	IEExtensions                    []FiveGProSePC5FlowBitRatesExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5FlowBitRates) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSePC5FlowBitRates", v)
}

// DecodeFiveGProSePC5FlowBitRates decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5FlowBitRates(data []byte) (FiveGProSePC5FlowBitRates, int, error) {
	var v FiveGProSePC5FlowBitRates
	err := ngapcodec.DecodeParameter("FiveGProSePC5FlowBitRates", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5FlowBitRates{}, 0, err
	}
	return v, len(data), err
}
