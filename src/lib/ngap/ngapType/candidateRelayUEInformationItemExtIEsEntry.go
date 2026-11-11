// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CandidateRelayUEInformationItemExtIEsEntryPresent uint32

const (
	CandidateRelayUEInformationItemExtIEsEntryPresentNothing CandidateRelayUEInformationItemExtIEsEntryPresent = iota
	CandidateRelayUEInformationItemExtIEsEntryPresentUnknown
)

type CandidateRelayUEInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CandidateRelayUEInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidateRelayUEInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidateRelayUEInformationItem-ExtIEsEntry", v)
}

// DecodeCandidateRelayUEInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCandidateRelayUEInformationItemExtIEsEntry(data []byte) (CandidateRelayUEInformationItemExtIEsEntry, int, error) {
	var v CandidateRelayUEInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("CandidateRelayUEInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidateRelayUEInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
