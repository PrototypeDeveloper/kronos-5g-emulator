// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BurstArrivalTimeWindowExtIEsEntryPresent uint32

const (
	BurstArrivalTimeWindowExtIEsEntryPresentNothing BurstArrivalTimeWindowExtIEsEntryPresent = iota
	BurstArrivalTimeWindowExtIEsEntryPresentUnknown
)

type BurstArrivalTimeWindowExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      BurstArrivalTimeWindowExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BurstArrivalTimeWindowExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BurstArrivalTimeWindow-ExtIEsEntry", v)
}

// DecodeBurstArrivalTimeWindowExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBurstArrivalTimeWindowExtIEsEntry(data []byte) (BurstArrivalTimeWindowExtIEsEntry, int, error) {
	var v BurstArrivalTimeWindowExtIEsEntry
	err := ngapcodec.DecodeParameter("BurstArrivalTimeWindow-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BurstArrivalTimeWindowExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
