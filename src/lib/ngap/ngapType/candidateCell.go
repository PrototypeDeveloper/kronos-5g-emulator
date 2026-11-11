// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateCell represents ASN.1 CandidateCell.
type CandidateCellPresent uint32

const (
	CandidateCellPresentNothing CandidateCellPresent = iota
	CandidateCellPresentCandidateCGI
	CandidateCellPresentCandidatePCI
	CandidateCellPresentChoiceExtensions
)

type CandidateCell struct {
	Present          CandidateCellPresent
	CandidateCGI     CandidateCellID
	CandidatePCI     CandidatePCI
	ChoiceExtensions CandidateCellExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCell) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CandidateCell", v) }

// DecodeCandidateCell decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCell(data []byte) (CandidateCell, int, error) {
	var v CandidateCell
	err := ngapcodec.DecodeParameter("CandidateCell", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCell{}, 0, err
	}
	return v, len(data), err
}
