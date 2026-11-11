// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ResetTypeExtIEsEntryPresent uint32

const (
	ResetTypeExtIEsEntryPresentNothing ResetTypeExtIEsEntryPresent = iota
	ResetTypeExtIEsEntryPresentUnknown
)

type ResetTypeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ResetTypeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ResetTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ResetType-ExtIEsEntry", v)
}

// DecodeResetTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeResetTypeExtIEsEntry(data []byte) (ResetTypeExtIEsEntry, int, error) {
	var v ResetTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("ResetType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ResetTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
