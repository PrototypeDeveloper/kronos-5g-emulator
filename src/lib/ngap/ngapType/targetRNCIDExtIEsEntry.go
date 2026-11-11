// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type TargetRNCIDExtIEsEntryPresent uint32

const (
	TargetRNCIDExtIEsEntryPresentNothing TargetRNCIDExtIEsEntryPresent = iota
	TargetRNCIDExtIEsEntryPresentUnknown
)

type TargetRNCIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      TargetRNCIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TargetRNCIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("TargetRNC-ID-ExtIEsEntry", v)
}

// DecodeTargetRNCIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeTargetRNCIDExtIEsEntry(data []byte) (TargetRNCIDExtIEsEntry, int, error) {
	var v TargetRNCIDExtIEsEntry
	err := ngapcodec.DecodeParameter("TargetRNC-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return TargetRNCIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
