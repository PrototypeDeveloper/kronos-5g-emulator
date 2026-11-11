// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SourceNodeIDExtIEsEntryPresent uint32

const (
	SourceNodeIDExtIEsEntryPresentNothing SourceNodeIDExtIEsEntryPresent = iota
	SourceNodeIDExtIEsEntryPresentUnknown
)

type SourceNodeIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      SourceNodeIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceNodeIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceNodeID-ExtIEsEntry", v)
}

// DecodeSourceNodeIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSourceNodeIDExtIEsEntry(data []byte) (SourceNodeIDExtIEsEntry, int, error) {
	var v SourceNodeIDExtIEsEntry
	err := ngapcodec.DecodeParameter("SourceNodeID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceNodeIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
