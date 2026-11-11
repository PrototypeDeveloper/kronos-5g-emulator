// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BroadcastCancelledAreaListExtIEsEntryPresent uint32

const (
	BroadcastCancelledAreaListExtIEsEntryPresentNothing BroadcastCancelledAreaListExtIEsEntryPresent = iota
	BroadcastCancelledAreaListExtIEsEntryPresentUnknown
)

type BroadcastCancelledAreaListExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      BroadcastCancelledAreaListExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastCancelledAreaListExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastCancelledAreaList-ExtIEsEntry", v)
}

// DecodeBroadcastCancelledAreaListExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastCancelledAreaListExtIEsEntry(data []byte) (BroadcastCancelledAreaListExtIEsEntry, int, error) {
	var v BroadcastCancelledAreaListExtIEsEntry
	err := ngapcodec.DecodeParameter("BroadcastCancelledAreaList-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastCancelledAreaListExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
