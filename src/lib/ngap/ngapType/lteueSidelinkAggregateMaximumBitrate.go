// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LTEUESidelinkAggregateMaximumBitrate represents ASN.1 LTEUESidelinkAggregateMaximumBitrate.
type LTEUESidelinkAggregateMaximumBitrate struct {
	UESidelinkAggregateMaximumBitRate BitRate
	HasIEExtensions                   bool
	IEExtensions                      []LTEUESidelinkAggregateMaximumBitratesExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEUESidelinkAggregateMaximumBitrate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTEUESidelinkAggregateMaximumBitrate", v)
}

// DecodeLTEUESidelinkAggregateMaximumBitrate decodes a standalone APER value and reports the consumed octets.
func DecodeLTEUESidelinkAggregateMaximumBitrate(data []byte) (LTEUESidelinkAggregateMaximumBitrate, int, error) {
	var v LTEUESidelinkAggregateMaximumBitrate
	err := ngapcodec.DecodeParameter("LTEUESidelinkAggregateMaximumBitrate", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEUESidelinkAggregateMaximumBitrate{}, 0, err
	}
	return v, len(data), err
}
