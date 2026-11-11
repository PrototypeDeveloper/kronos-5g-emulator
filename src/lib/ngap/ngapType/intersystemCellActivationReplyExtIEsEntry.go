// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemCellActivationReplyExtIEsEntryPresent uint32

const (
	IntersystemCellActivationReplyExtIEsEntryPresentNothing IntersystemCellActivationReplyExtIEsEntryPresent = iota
	IntersystemCellActivationReplyExtIEsEntryPresentUnknown
)

type IntersystemCellActivationReplyExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemCellActivationReplyExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemCellActivationReplyExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemCellActivationReply-ExtIEsEntry", v)
}

// DecodeIntersystemCellActivationReplyExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemCellActivationReplyExtIEsEntry(data []byte) (IntersystemCellActivationReplyExtIEsEntry, int, error) {
	var v IntersystemCellActivationReplyExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemCellActivationReply-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemCellActivationReplyExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
