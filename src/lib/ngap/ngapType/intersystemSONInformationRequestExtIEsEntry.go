// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONInformationRequestExtIEsEntryPresent uint32

const (
	IntersystemSONInformationRequestExtIEsEntryPresentNothing IntersystemSONInformationRequestExtIEsEntryPresent = iota
	IntersystemSONInformationRequestExtIEsEntryPresentUnknown
)

type IntersystemSONInformationRequestExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      IntersystemSONInformationRequestExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationRequestExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformationRequest-ExtIEsEntry", v)
}

// DecodeIntersystemSONInformationRequestExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationRequestExtIEsEntry(data []byte) (IntersystemSONInformationRequestExtIEsEntry, int, error) {
	var v IntersystemSONInformationRequestExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONInformationRequest-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationRequestExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
