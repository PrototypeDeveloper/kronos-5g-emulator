// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRUESidelinkAggregateMaximumBitrateExtIEsEntryPresent uint32

const (
	NRUESidelinkAggregateMaximumBitrateExtIEsEntryPresentNothing NRUESidelinkAggregateMaximumBitrateExtIEsEntryPresent = iota
	NRUESidelinkAggregateMaximumBitrateExtIEsEntryPresentUnknown
)

type NRUESidelinkAggregateMaximumBitrateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRUESidelinkAggregateMaximumBitrateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRUESidelinkAggregateMaximumBitrateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRUESidelinkAggregateMaximumBitrate-ExtIEsEntry", v)
}

// DecodeNRUESidelinkAggregateMaximumBitrateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRUESidelinkAggregateMaximumBitrateExtIEsEntry(data []byte) (NRUESidelinkAggregateMaximumBitrateExtIEsEntry, int, error) {
	var v NRUESidelinkAggregateMaximumBitrateExtIEsEntry
	err := ngapcodec.DecodeParameter("NRUESidelinkAggregateMaximumBitrate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRUESidelinkAggregateMaximumBitrateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
