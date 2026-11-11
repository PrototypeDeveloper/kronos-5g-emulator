// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateCellID represents ASN.1 CandidateCellID.
type CandidateCellID struct {
	CandidateCellID   NRCGI
	HasIEExtensions   bool
	IEExtensions      []CandidateCellIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateCellID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateCellID", v)
}

// DecodeCandidateCellID decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateCellID(data []byte) (CandidateCellID, int, error) {
	var v CandidateCellID
	err := ngapcodec.DecodeParameter("CandidateCellID", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateCellID{}, 0, err
	}
	return v, len(data), err
}
