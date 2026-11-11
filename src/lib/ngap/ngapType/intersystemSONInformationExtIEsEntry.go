// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONInformationExtIEsEntryPresent uint32

const (
	IntersystemSONInformationExtIEsEntryPresentNothing IntersystemSONInformationExtIEsEntryPresent = iota
	IntersystemSONInformationExtIEsEntryPresentIntersystemSONInformationRequest
	IntersystemSONInformationExtIEsEntryPresentIntersystemSONInformationReply
	IntersystemSONInformationExtIEsEntryPresentUnknown
)

type IntersystemSONInformationExtIEsEntry struct {
	ID                               ProtocolIEID
	Criticality                      Criticality
	Present                          IntersystemSONInformationExtIEsEntryPresent
	IntersystemSONInformationRequest IntersystemSONInformationRequest `ngap:"id:290,criticality:ignore,presence:mandatory"`
	IntersystemSONInformationReply   IntersystemSONInformationReply   `ngap:"id:291,criticality:ignore,presence:mandatory"`
	UnknownValue                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformation-ExtIEsEntry", v)
}

// DecodeIntersystemSONInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationExtIEsEntry(data []byte) (IntersystemSONInformationExtIEsEntry, int, error) {
	var v IntersystemSONInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
