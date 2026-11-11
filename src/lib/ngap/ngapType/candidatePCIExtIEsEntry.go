// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CandidatePCIExtIEsEntryPresent uint32

const (
	CandidatePCIExtIEsEntryPresentNothing CandidatePCIExtIEsEntryPresent = iota
	CandidatePCIExtIEsEntryPresentUnknown
)

type CandidatePCIExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      CandidatePCIExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CandidatePCIExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CandidatePCI-ExtIEsEntry", v)
}

// DecodeCandidatePCIExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCandidatePCIExtIEsEntry(data []byte) (CandidatePCIExtIEsEntry, int, error) {
	var v CandidatePCIExtIEsEntry
	err := ngapcodec.DecodeParameter("CandidatePCI-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CandidatePCIExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
