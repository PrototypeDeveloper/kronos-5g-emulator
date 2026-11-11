// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidatePCI represents ASN.1 CandidatePCI.
type CandidatePCI struct {
	CandidatePCI      int64 `aper:"valueExt,valueLB:0,valueUB:1007"`
	CandidateNRARFCN  int64 `aper:"valueLB:0,valueUB:3279165"`
	HasIEExtensions   bool
	IEExtensions      []CandidatePCIExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidatePCI) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CandidatePCI", v) }

// DecodeCandidatePCI decodes a standalone APER value and reports the consumed octets.
func DecodeCandidatePCI(data []byte) (CandidatePCI, int, error) {
	var v CandidatePCI
	err := ngapcodec.DecodeParameter("CandidatePCI", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidatePCI{}, 0, err
	}
	return v, len(data), err
}
