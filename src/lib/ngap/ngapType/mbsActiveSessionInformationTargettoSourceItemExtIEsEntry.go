// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type MBSActiveSessionInformationTargettoSourceItemExtIEsEntryPresent uint32

const (
	MBSActiveSessionInformationTargettoSourceItemExtIEsEntryPresentNothing MBSActiveSessionInformationTargettoSourceItemExtIEsEntryPresent = iota
	MBSActiveSessionInformationTargettoSourceItemExtIEsEntryPresentUnknown
)

type MBSActiveSessionInformationTargettoSourceItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      MBSActiveSessionInformationTargettoSourceItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSActiveSessionInformationTargettoSourceItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-ActiveSessionInformation-TargettoSourceItem-ExtIEsEntry", v)
}

// DecodeMBSActiveSessionInformationTargettoSourceItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeMBSActiveSessionInformationTargettoSourceItemExtIEsEntry(data []byte) (MBSActiveSessionInformationTargettoSourceItemExtIEsEntry, int, error) {
	var v MBSActiveSessionInformationTargettoSourceItemExtIEsEntry
	err := ngapcodec.DecodeParameter("MBS-ActiveSessionInformation-TargettoSourceItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSActiveSessionInformationTargettoSourceItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
