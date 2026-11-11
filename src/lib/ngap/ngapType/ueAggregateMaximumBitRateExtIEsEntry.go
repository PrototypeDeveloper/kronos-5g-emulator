// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UEAggregateMaximumBitRateExtIEsEntryPresent uint32

const (
	UEAggregateMaximumBitRateExtIEsEntryPresentNothing UEAggregateMaximumBitRateExtIEsEntryPresent = iota
	UEAggregateMaximumBitRateExtIEsEntryPresentUnknown
)

type UEAggregateMaximumBitRateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UEAggregateMaximumBitRateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAggregateMaximumBitRateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEAggregateMaximumBitRate-ExtIEsEntry", v)
}

// DecodeUEAggregateMaximumBitRateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUEAggregateMaximumBitRateExtIEsEntry(data []byte) (UEAggregateMaximumBitRateExtIEsEntry, int, error) {
	var v UEAggregateMaximumBitRateExtIEsEntry
	err := ngapcodec.DecodeParameter("UEAggregateMaximumBitRate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAggregateMaximumBitRateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
