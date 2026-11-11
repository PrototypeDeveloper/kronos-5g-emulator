// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TAIListForPagingItemExtIEsEntryPresent uint32

const (
	TAIListForPagingItemExtIEsEntryPresentNothing TAIListForPagingItemExtIEsEntryPresent = iota
	TAIListForPagingItemExtIEsEntryPresentUnknown
)

type TAIListForPagingItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TAIListForPagingItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TAIListForPagingItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TAIListForPagingItem-ExtIEsEntry", v)
}

// DecodeTAIListForPagingItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTAIListForPagingItemExtIEsEntry(data []byte) (TAIListForPagingItemExtIEsEntry, int, error) {
	var v TAIListForPagingItemExtIEsEntry
	err := ngapcodec.DecodeParameter("TAIListForPagingItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TAIListForPagingItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
