// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSSessionSetupRequestItemExtIEsEntryPresent uint32

const (
	MBSSessionSetupRequestItemExtIEsEntryPresentNothing MBSSessionSetupRequestItemExtIEsEntryPresent = iota
	MBSSessionSetupRequestItemExtIEsEntryPresentMBSAssistanceInformation
	MBSSessionSetupRequestItemExtIEsEntryPresentUnknown
)

type MBSSessionSetupRequestItemExtIEsEntry struct {
	ID                       ProtocolExtensionID
	Criticality              Criticality
	Present                  MBSSessionSetupRequestItemExtIEsEntryPresent
	MBSAssistanceInformation MBSAssistanceInformation `ngap:"id:416,criticality:ignore,presence:optional"`
	UnknownValue             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupRequestItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupRequestItem-ExtIEsEntry", v)
}

// DecodeMBSSessionSetupRequestItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupRequestItemExtIEsEntry(data []byte) (MBSSessionSetupRequestItemExtIEsEntry, int, error) {
	var v MBSSessionSetupRequestItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBSSessionSetupRequestItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupRequestItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
