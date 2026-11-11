// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSActiveSessionInformationSourcetoTargetItemExtIEsEntryPresent uint32

const (
	MBSActiveSessionInformationSourcetoTargetItemExtIEsEntryPresentNothing MBSActiveSessionInformationSourcetoTargetItemExtIEsEntryPresent = iota
	MBSActiveSessionInformationSourcetoTargetItemExtIEsEntryPresentUnknown
)

type MBSActiveSessionInformationSourcetoTargetItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSActiveSessionInformationSourcetoTargetItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSActiveSessionInformationSourcetoTargetItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ActiveSessionInformation-SourcetoTargetItem-ExtIEsEntry", v)
}

// DecodeMBSActiveSessionInformationSourcetoTargetItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSActiveSessionInformationSourcetoTargetItemExtIEsEntry(data []byte) (MBSActiveSessionInformationSourcetoTargetItemExtIEsEntry, int, error) {
	var v MBSActiveSessionInformationSourcetoTargetItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-ActiveSessionInformation-SourcetoTargetItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSActiveSessionInformationSourcetoTargetItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
