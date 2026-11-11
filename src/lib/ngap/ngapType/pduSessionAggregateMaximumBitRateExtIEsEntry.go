// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionAggregateMaximumBitRateExtIEsEntryPresent uint32

const (
	PDUSessionAggregateMaximumBitRateExtIEsEntryPresentNothing PDUSessionAggregateMaximumBitRateExtIEsEntryPresent = iota
	PDUSessionAggregateMaximumBitRateExtIEsEntryPresentUnknown
)

type PDUSessionAggregateMaximumBitRateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionAggregateMaximumBitRateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionAggregateMaximumBitRateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionAggregateMaximumBitRate-ExtIEsEntry", v)
}

// DecodePDUSessionAggregateMaximumBitRateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionAggregateMaximumBitRateExtIEsEntry(data []byte) (PDUSessionAggregateMaximumBitRateExtIEsEntry, int, error) {
	var v PDUSessionAggregateMaximumBitRateExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionAggregateMaximumBitRate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionAggregateMaximumBitRateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
