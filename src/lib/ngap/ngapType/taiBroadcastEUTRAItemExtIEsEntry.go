// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIBroadcastEUTRAItemExtIEsEntryPresent uint32

const (
	TAIBroadcastEUTRAItemExtIEsEntryPresentNothing TAIBroadcastEUTRAItemExtIEsEntryPresent = iota
	TAIBroadcastEUTRAItemExtIEsEntryPresentUnknown
)

type TAIBroadcastEUTRAItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIBroadcastEUTRAItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIBroadcastEUTRAItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIBroadcastEUTRA-Item-ExtIEsEntry", v)
}

// DecodeTAIBroadcastEUTRAItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIBroadcastEUTRAItemExtIEsEntry(data []byte) (TAIBroadcastEUTRAItemExtIEsEntry, int, error) {
	var v TAIBroadcastEUTRAItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIBroadcastEUTRA-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIBroadcastEUTRAItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
