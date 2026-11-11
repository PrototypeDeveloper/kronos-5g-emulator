// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DLCountChoiceExtIEsEntryPresent uint32

const (
	DLCountChoiceExtIEsEntryPresentNothing DLCountChoiceExtIEsEntryPresent = iota
	DLCountChoiceExtIEsEntryPresentUnknown
)

type DLCountChoiceExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      DLCountChoiceExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLCountChoiceExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DLCountChoice-ExtIEsEntry", v)
}

// DecodeDLCountChoiceExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDLCountChoiceExtIEsEntry(data []byte) (DLCountChoiceExtIEsEntry, int, error) {
	var v DLCountChoiceExtIEsEntry
	err := ngapcodec.DecodeParameter("DLCountChoice-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLCountChoiceExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
