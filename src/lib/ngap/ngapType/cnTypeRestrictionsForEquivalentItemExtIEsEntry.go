// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CNTypeRestrictionsForEquivalentItemExtIEsEntryPresent uint32

const (
	CNTypeRestrictionsForEquivalentItemExtIEsEntryPresentNothing CNTypeRestrictionsForEquivalentItemExtIEsEntryPresent = iota
	CNTypeRestrictionsForEquivalentItemExtIEsEntryPresentUnknown
)

type CNTypeRestrictionsForEquivalentItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CNTypeRestrictionsForEquivalentItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CNTypeRestrictionsForEquivalentItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CNTypeRestrictionsForEquivalentItem-ExtIEsEntry", v)
}

// DecodeCNTypeRestrictionsForEquivalentItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCNTypeRestrictionsForEquivalentItemExtIEsEntry(data []byte) (CNTypeRestrictionsForEquivalentItemExtIEsEntry, int, error) {
	var v CNTypeRestrictionsForEquivalentItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CNTypeRestrictionsForEquivalentItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CNTypeRestrictionsForEquivalentItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
