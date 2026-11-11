// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PDUSessionAggregateMaximumBitRate represents ASN.1 PDUSessionAggregateMaximumBitRate.
type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL BitRate
	PDUSessionAggregateMaximumBitRateUL BitRate
	HasIEExtensions                     bool
	IEExtensions                        []PDUSessionAggregateMaximumBitRateExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionAggregateMaximumBitRate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionAggregateMaximumBitRate", v)
}

// DecodePDUSessionAggregateMaximumBitRate decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionAggregateMaximumBitRate(data []byte) (PDUSessionAggregateMaximumBitRate, int, error) {
	var v PDUSessionAggregateMaximumBitRate
	err := ngapcodec.DecodeParameter("PDUSessionAggregateMaximumBitRate", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionAggregateMaximumBitRate{}, 0, err
	}
	return v, len(data), err
}
