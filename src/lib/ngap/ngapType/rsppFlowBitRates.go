// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RSPPFlowBitRates represents ASN.1 RSPPFlowBitRates.
type RSPPFlowBitRates struct {
	GuaranteedFlowBitRate BitRate
	MaximumFlowBitRate    BitRate
	HasIEExtensions       bool
	IEExtensions          []RSPPFlowBitRatesExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RSPPFlowBitRates) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RSPPFlowBitRates", v)
}

// DecodeRSPPFlowBitRates decodes a standalone APER value and reports the consumed octets.
func DecodeRSPPFlowBitRates(data []byte) (RSPPFlowBitRates, int, error) {
	var v RSPPFlowBitRates
	err := ngapcodec.DecodeParameter("RSPPFlowBitRates", data, &v)
	if ngapcodec.IsFatal(err) {
		return RSPPFlowBitRates{}, 0, err
	}
	return v, len(data), err
}
