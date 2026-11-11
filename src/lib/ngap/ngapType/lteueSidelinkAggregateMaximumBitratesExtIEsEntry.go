// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LTEUESidelinkAggregateMaximumBitratesExtIEsEntryPresent uint32

const (
	LTEUESidelinkAggregateMaximumBitratesExtIEsEntryPresentNothing LTEUESidelinkAggregateMaximumBitratesExtIEsEntryPresent = iota
	LTEUESidelinkAggregateMaximumBitratesExtIEsEntryPresentUnknown
)

type LTEUESidelinkAggregateMaximumBitratesExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      LTEUESidelinkAggregateMaximumBitratesExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEUESidelinkAggregateMaximumBitratesExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTEUE-Sidelink-Aggregate-MaximumBitrates-ExtIEsEntry", v)
}

// DecodeLTEUESidelinkAggregateMaximumBitratesExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLTEUESidelinkAggregateMaximumBitratesExtIEsEntry(data []byte) (LTEUESidelinkAggregateMaximumBitratesExtIEsEntry, int, error) {
	var v LTEUESidelinkAggregateMaximumBitratesExtIEsEntry
	err := ngapcodec.DecodeParameter("LTEUE-Sidelink-Aggregate-MaximumBitrates-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEUESidelinkAggregateMaximumBitratesExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
