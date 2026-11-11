// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type WarningAreaListExtIEsEntryPresent uint32

const (
	WarningAreaListExtIEsEntryPresentNothing WarningAreaListExtIEsEntryPresent = iota
	WarningAreaListExtIEsEntryPresentUnknown
)

type WarningAreaListExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      WarningAreaListExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WarningAreaListExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WarningAreaList-ExtIEsEntry", v)
}

// DecodeWarningAreaListExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeWarningAreaListExtIEsEntry(data []byte) (WarningAreaListExtIEsEntry, int, error) {
	var v WarningAreaListExtIEsEntry
	err := ngapcodec.DecodeParameter("WarningAreaList-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return WarningAreaListExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
