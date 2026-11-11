// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ENBIDExtIEsEntryPresent uint32

const (
	ENBIDExtIEsEntryPresentNothing ENBIDExtIEsEntryPresent = iota
	ENBIDExtIEsEntryPresentUnknown
)

type ENBIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ENBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ENBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ENB-ID-ExtIEsEntry", v)
}

// DecodeENBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeENBIDExtIEsEntry(data []byte) (ENBIDExtIEsEntry, int, error) {
	var v ENBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("ENB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ENBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
