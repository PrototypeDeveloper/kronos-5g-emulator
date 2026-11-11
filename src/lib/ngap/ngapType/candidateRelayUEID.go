// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateRelayUEID represents ASN.1 CandidateRelayUE-ID.
type CandidateRelayUEID struct {
	Value BitString `aper:"sizeLB:24,sizeUB:24"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateRelayUEID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateRelayUE-ID", v)
}

// DecodeCandidateRelayUEID decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateRelayUEID(data []byte) (CandidateRelayUEID, int, error) {
	var v CandidateRelayUEID
	err := ngapcodec.DecodeParameter("CandidateRelayUE-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateRelayUEID{}, 0, err
	}
	return v, len(data), err
}
