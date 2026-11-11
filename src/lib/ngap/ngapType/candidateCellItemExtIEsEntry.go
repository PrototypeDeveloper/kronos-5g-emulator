// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CandidateCellItemExtIEsEntryPresent uint32

const (
	CandidateCellItemExtIEsEntryPresentNothing CandidateCellItemExtIEsEntryPresent = iota
	CandidateCellItemExtIEsEntryPresentUnknown
)

type CandidateCellItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CandidateCellItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCellItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateCellItem-ExtIEsEntry", v)
}

// DecodeCandidateCellItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCellItemExtIEsEntry(data []byte) (CandidateCellItemExtIEsEntry, int, error) {
	var v CandidateCellItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CandidateCellItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCellItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
