// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SONInformationReplyExtIEsEntryPresent uint32

const (
	SONInformationReplyExtIEsEntryPresentNothing SONInformationReplyExtIEsEntryPresent = iota
	SONInformationReplyExtIEsEntryPresentUnknown
)

type SONInformationReplyExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SONInformationReplyExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformationReplyExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformationReply-ExtIEsEntry", v)
}

// DecodeSONInformationReplyExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformationReplyExtIEsEntry(data []byte) (SONInformationReplyExtIEsEntry, int, error) {
	var v SONInformationReplyExtIEsEntry
	err := ngapcodec.DecodeParameter("SONInformationReply-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformationReplyExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
