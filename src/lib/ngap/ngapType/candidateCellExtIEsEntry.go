// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CandidateCellExtIEsEntryPresent uint32

const (
	CandidateCellExtIEsEntryPresentNothing CandidateCellExtIEsEntryPresent = iota
	CandidateCellExtIEsEntryPresentUnknown
)

type CandidateCellExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      CandidateCellExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCellExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateCell-ExtIEsEntry", v)
}

// DecodeCandidateCellExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCellExtIEsEntry(data []byte) (CandidateCellExtIEsEntry, int, error) {
	var v CandidateCellExtIEsEntry
	err := ngapcodec.DecodeParameter("CandidateCell-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCellExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
