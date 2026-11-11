// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEAggregateMaximumBitRate represents ASN.1 UEAggregateMaximumBitRate.
type UEAggregateMaximumBitRate struct {
	UEAggregateMaximumBitRateDL BitRate
	UEAggregateMaximumBitRateUL BitRate
	HasIEExtensions             bool
	IEExtensions                []UEAggregateMaximumBitRateExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAggregateMaximumBitRate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEAggregateMaximumBitRate", v)
}

// DecodeUEAggregateMaximumBitRate decodes a standalone APER value and reports the consumed octets.
func DecodeUEAggregateMaximumBitRate(data []byte) (UEAggregateMaximumBitRate, int, error) {
	var v UEAggregateMaximumBitRate
	err := ngapcodec.DecodeParameter("UEAggregateMaximumBitRate", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAggregateMaximumBitRate{}, 0, err
	}
	return v, len(data), err
}
