// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetRANNodeIDExtIEsEntryPresent uint32

const (
	TargetRANNodeIDExtIEsEntryPresentNothing TargetRANNodeIDExtIEsEntryPresent = iota
	TargetRANNodeIDExtIEsEntryPresentSelectedTargetSNPNIdentity
	TargetRANNodeIDExtIEsEntryPresentUnknown
)

type TargetRANNodeIDExtIEsEntry struct {
	ID                         ProtocolExtensionID
	Criticality                Criticality
	Present                    TargetRANNodeIDExtIEsEntryPresent
	SelectedTargetSNPNIdentity NID `ngap:"id:369,criticality:reject,presence:optional"`
	UnknownValue               []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRANNodeIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRANNodeID-ExtIEsEntry", v)
}

// DecodeTargetRANNodeIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRANNodeIDExtIEsEntry(data []byte) (TargetRANNodeIDExtIEsEntry, int, error) {
	var v TargetRANNodeIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetRANNodeID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRANNodeIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
