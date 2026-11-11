// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GNBIDExtIEsEntryPresent uint32

const (
	GNBIDExtIEsEntryPresentNothing GNBIDExtIEsEntryPresent = iota
	GNBIDExtIEsEntryPresentUnknown
)

type GNBIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      GNBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GNBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GNB-ID-ExtIEsEntry", v)
}

// DecodeGNBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGNBIDExtIEsEntry(data []byte) (GNBIDExtIEsEntry, int, error) {
	var v GNBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("GNB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GNBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
