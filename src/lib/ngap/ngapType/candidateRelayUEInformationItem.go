// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CandidateRelayUEInformationItem represents ASN.1 CandidateRelayUEInformationItem.
type CandidateRelayUEInformationItem struct {
	CandidateRelayUEId CandidateRelayUEID
	HasIEExtensions    bool
	IEExtensions       []CandidateRelayUEInformationItemExtIEsEntry
	UnknownExtensions  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateRelayUEInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateRelayUEInformationItem", v)
}

// DecodeCandidateRelayUEInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateRelayUEInformationItem(data []byte) (CandidateRelayUEInformationItem, int, error) {
	var v CandidateRelayUEInformationItem
	err := ngapcodec.DecodeParameter("CandidateRelayUEInformationItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateRelayUEInformationItem{}, 0, err
	}
	return v, len(data), err
}
