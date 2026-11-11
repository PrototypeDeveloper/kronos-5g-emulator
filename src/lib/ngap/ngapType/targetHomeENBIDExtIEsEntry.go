// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetHomeENBIDExtIEsEntryPresent uint32

const (
	TargetHomeENBIDExtIEsEntryPresentNothing TargetHomeENBIDExtIEsEntryPresent = iota
	TargetHomeENBIDExtIEsEntryPresentUnknown
)

type TargetHomeENBIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargetHomeENBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetHomeENBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetHomeENB-ID-ExtIEsEntry", v)
}

// DecodeTargetHomeENBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetHomeENBIDExtIEsEntry(data []byte) (TargetHomeENBIDExtIEsEntry, int, error) {
	var v TargetHomeENBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetHomeENB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetHomeENBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
