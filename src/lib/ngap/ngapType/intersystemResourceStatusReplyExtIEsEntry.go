// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemResourceStatusReplyExtIEsEntryPresent uint32

const (
	IntersystemResourceStatusReplyExtIEsEntryPresentNothing IntersystemResourceStatusReplyExtIEsEntryPresent = iota
	IntersystemResourceStatusReplyExtIEsEntryPresentUnknown
)

type IntersystemResourceStatusReplyExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemResourceStatusReplyExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceStatusReplyExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceStatusReply-ExtIEsEntry", v)
}

// DecodeIntersystemResourceStatusReplyExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceStatusReplyExtIEsEntry(data []byte) (IntersystemResourceStatusReplyExtIEsEntry, int, error) {
	var v IntersystemResourceStatusReplyExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemResourceStatusReply-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceStatusReplyExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
