// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SourceRANNodeIDExtIEsEntryPresent uint32

const (
	SourceRANNodeIDExtIEsEntryPresentNothing SourceRANNodeIDExtIEsEntryPresent = iota
	SourceRANNodeIDExtIEsEntryPresentUnknown
)

type SourceRANNodeIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SourceRANNodeIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SourceRANNodeIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SourceRANNodeID-ExtIEsEntry", v)
}

// DecodeSourceRANNodeIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSourceRANNodeIDExtIEsEntry(data []byte) (SourceRANNodeIDExtIEsEntry, int, error) {
	var v SourceRANNodeIDExtIEsEntry
	err := ngapcodec.DecodeParameter("SourceRANNodeID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SourceRANNodeIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
