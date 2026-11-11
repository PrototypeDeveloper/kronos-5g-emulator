// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBsSubjectToDLDiscardingItemExtIEsEntryPresent uint32

const (
	DRBsSubjectToDLDiscardingItemExtIEsEntryPresentNothing DRBsSubjectToDLDiscardingItemExtIEsEntryPresent = iota
	DRBsSubjectToDLDiscardingItemExtIEsEntryPresentUnknown
)

type DRBsSubjectToDLDiscardingItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DRBsSubjectToDLDiscardingItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBsSubjectToDLDiscardingItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBsSubjectToDLDiscarding-Item-ExtIEsEntry", v)
}

// DecodeDRBsSubjectToDLDiscardingItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBsSubjectToDLDiscardingItemExtIEsEntry(data []byte) (DRBsSubjectToDLDiscardingItemExtIEsEntry, int, error) {
	var v DRBsSubjectToDLDiscardingItemExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBsSubjectToDLDiscarding-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBsSubjectToDLDiscardingItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
