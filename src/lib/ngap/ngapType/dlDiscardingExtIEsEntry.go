// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DLDiscardingExtIEsEntryPresent uint32

const (
	DLDiscardingExtIEsEntryPresentNothing DLDiscardingExtIEsEntryPresent = iota
	DLDiscardingExtIEsEntryPresentUnknown
)

type DLDiscardingExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DLDiscardingExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DLDiscardingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DLDiscarding-ExtIEsEntry", v)
}

// DecodeDLDiscardingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDLDiscardingExtIEsEntry(data []byte) (DLDiscardingExtIEsEntry, int, error) {
	var v DLDiscardingExtIEsEntry
	err := ngapcodec.DecodeParameter("DLDiscarding-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DLDiscardingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
