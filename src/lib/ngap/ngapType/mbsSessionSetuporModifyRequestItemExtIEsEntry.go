// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionSetuporModifyRequestItemExtIEsEntryPresent uint32

const (
	MBSSessionSetuporModifyRequestItemExtIEsEntryPresentNothing MBSSessionSetuporModifyRequestItemExtIEsEntryPresent = iota
	MBSSessionSetuporModifyRequestItemExtIEsEntryPresentMBSAssistanceInformation
	MBSSessionSetuporModifyRequestItemExtIEsEntryPresentUnknown
)

type MBSSessionSetuporModifyRequestItemExtIEsEntry struct {
	ID                       ProtocolExtensionID
	Criticality              Criticality
	Present                  MBSSessionSetuporModifyRequestItemExtIEsEntryPresent
	MBSAssistanceInformation MBSAssistanceInformation `ngap:"id:416,criticality:ignore,presence:optional"`
	UnknownValue             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetuporModifyRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetuporModifyRequestItem-ExtIEsEntry", v)
}

// DecodeMBSSessionSetuporModifyRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetuporModifyRequestItemExtIEsEntry(data []byte) (MBSSessionSetuporModifyRequestItemExtIEsEntry, int, error) {
	var v MBSSessionSetuporModifyRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionSetuporModifyRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetuporModifyRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
