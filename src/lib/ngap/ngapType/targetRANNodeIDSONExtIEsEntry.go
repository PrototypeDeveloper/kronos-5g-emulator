// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetRANNodeIDSONExtIEsEntryPresent uint32

const (
	TargetRANNodeIDSONExtIEsEntryPresentNothing TargetRANNodeIDSONExtIEsEntryPresent = iota
	TargetRANNodeIDSONExtIEsEntryPresentNRCGI
	TargetRANNodeIDSONExtIEsEntryPresentUnknown
)

type TargetRANNodeIDSONExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargetRANNodeIDSONExtIEsEntryPresent
	NRCGI        NRCGI `ngap:"id:45,criticality:ignore,presence:optional"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRANNodeIDSONExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRANNodeID-SON-ExtIEsEntry", v)
}

// DecodeTargetRANNodeIDSONExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRANNodeIDSONExtIEsEntry(data []byte) (TargetRANNodeIDSONExtIEsEntry, int, error) {
	var v TargetRANNodeIDSONExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetRANNodeID-SON-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRANNodeIDSONExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
