// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIBroadcastNRItemExtIEsEntryPresent uint32

const (
	TAIBroadcastNRItemExtIEsEntryPresentNothing TAIBroadcastNRItemExtIEsEntryPresent = iota
	TAIBroadcastNRItemExtIEsEntryPresentUnknown
)

type TAIBroadcastNRItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIBroadcastNRItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBroadcastNRItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBroadcastNR-Item-ExtIEsEntry", v)
}

// DecodeTAIBroadcastNRItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBroadcastNRItemExtIEsEntry(data []byte) (TAIBroadcastNRItemExtIEsEntry, int, error) {
	var v TAIBroadcastNRItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIBroadcastNR-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBroadcastNRItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
