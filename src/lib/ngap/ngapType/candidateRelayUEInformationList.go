// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateRelayUEInformationList represents ASN.1 CandidateRelayUEInformationList.
type CandidateRelayUEInformationList struct {
	List []CandidateRelayUEInformationItem `aper:"sizeLB:1,sizeUB:32"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateRelayUEInformationList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateRelayUEInformationList", v)
}

// DecodeCandidateRelayUEInformationList decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateRelayUEInformationList(data []byte) (CandidateRelayUEInformationList, int, error) {
	var v CandidateRelayUEInformationList
	err := ngapcodec.DecodeParameter("CandidateRelayUEInformationList", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateRelayUEInformationList{}, 0, err
	}
	return v, len(data), err
}
