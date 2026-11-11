// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CandidateCellIDExtIEsEntryPresent uint32

const (
	CandidateCellIDExtIEsEntryPresentNothing CandidateCellIDExtIEsEntryPresent = iota
	CandidateCellIDExtIEsEntryPresentUnknown
)

type CandidateCellIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CandidateCellIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCellIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateCellID-ExtIEsEntry", v)
}

// DecodeCandidateCellIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCellIDExtIEsEntry(data []byte) (CandidateCellIDExtIEsEntry, int, error) {
	var v CandidateCellIDExtIEsEntry
	err := ngapcodec.DecodeParameter("CandidateCellID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCellIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
