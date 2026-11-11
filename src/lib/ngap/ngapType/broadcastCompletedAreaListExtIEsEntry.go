// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type BroadcastCompletedAreaListExtIEsEntryPresent uint32

const (
	BroadcastCompletedAreaListExtIEsEntryPresentNothing BroadcastCompletedAreaListExtIEsEntryPresent = iota
	BroadcastCompletedAreaListExtIEsEntryPresentUnknown
)

type BroadcastCompletedAreaListExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      BroadcastCompletedAreaListExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BroadcastCompletedAreaListExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BroadcastCompletedAreaList-ExtIEsEntry", v)
}

// DecodeBroadcastCompletedAreaListExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeBroadcastCompletedAreaListExtIEsEntry(data []byte) (BroadcastCompletedAreaListExtIEsEntry, int, error) {
	var v BroadcastCompletedAreaListExtIEsEntry
	err := ngapcodec.DecodeParameter("BroadcastCompletedAreaList-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return BroadcastCompletedAreaListExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
