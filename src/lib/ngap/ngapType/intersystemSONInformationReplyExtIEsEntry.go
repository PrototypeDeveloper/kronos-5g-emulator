// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONInformationReplyExtIEsEntryPresent uint32

const (
	IntersystemSONInformationReplyExtIEsEntryPresentNothing IntersystemSONInformationReplyExtIEsEntryPresent = iota
	IntersystemSONInformationReplyExtIEsEntryPresentUnknown
)

type IntersystemSONInformationReplyExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      IntersystemSONInformationReplyExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationReplyExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformationReply-ExtIEsEntry", v)
}

// DecodeIntersystemSONInformationReplyExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationReplyExtIEsEntry(data []byte) (IntersystemSONInformationReplyExtIEsEntry, int, error) {
	var v IntersystemSONInformationReplyExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONInformationReply-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationReplyExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
