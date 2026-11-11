// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetIDExtIEsEntryPresent uint32

const (
	TargetIDExtIEsEntryPresentNothing TargetIDExtIEsEntryPresent = iota
	TargetIDExtIEsEntryPresentTargetRNCID
	TargetIDExtIEsEntryPresentTargetHomeENBID
	TargetIDExtIEsEntryPresentUnknown
)

type TargetIDExtIEsEntry struct {
	ID              ProtocolIEID
	Criticality     Criticality
	Present         TargetIDExtIEsEntryPresent
	TargetRNCID     TargetRNCID     `ngap:"id:178,criticality:reject,presence:mandatory"`
	TargetHomeENBID TargetHomeENBID `ngap:"id:364,criticality:reject,presence:mandatory"`
	UnknownValue    []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetID-ExtIEsEntry", v)
}

// DecodeTargetIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetIDExtIEsEntry(data []byte) (TargetIDExtIEsEntry, int, error) {
	var v TargetIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
