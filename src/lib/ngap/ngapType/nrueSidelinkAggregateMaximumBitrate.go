// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRUESidelinkAggregateMaximumBitrate represents ASN.1 NRUESidelinkAggregateMaximumBitrate.
type NRUESidelinkAggregateMaximumBitrate struct {
	UESidelinkAggregateMaximumBitRate BitRate
	HasIEExtensions                   bool
	IEExtensions                      []NRUESidelinkAggregateMaximumBitrateExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRUESidelinkAggregateMaximumBitrate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRUESidelinkAggregateMaximumBitrate", v)
}

// DecodeNRUESidelinkAggregateMaximumBitrate decodes a standalone APER value and reports the consumed octets.
func DecodeNRUESidelinkAggregateMaximumBitrate(data []byte) (NRUESidelinkAggregateMaximumBitrate, int, error) {
	var v NRUESidelinkAggregateMaximumBitrate
	err := ngapcodec.DecodeParameter("NRUESidelinkAggregateMaximumBitrate", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRUESidelinkAggregateMaximumBitrate{}, 0, err
	}
	return v, len(data), err
}
