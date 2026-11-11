// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIListForInactiveItemExtIEsEntryPresent uint32

const (
	TAIListForInactiveItemExtIEsEntryPresentNothing TAIListForInactiveItemExtIEsEntryPresent = iota
	TAIListForInactiveItemExtIEsEntryPresentUnknown
)

type TAIListForInactiveItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIListForInactiveItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForInactiveItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForInactiveItem-ExtIEsEntry", v)
}

// DecodeTAIListForInactiveItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForInactiveItemExtIEsEntry(data []byte) (TAIListForInactiveItemExtIEsEntry, int, error) {
	var v TAIListForInactiveItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIListForInactiveItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForInactiveItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
