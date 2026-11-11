// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBStatusULExtIEsEntryPresent uint32

const (
	DRBStatusULExtIEsEntryPresentNothing DRBStatusULExtIEsEntryPresent = iota
	DRBStatusULExtIEsEntryPresentUnknown
)

type DRBStatusULExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      DRBStatusULExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusULExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBStatusUL-ExtIEsEntry", v)
}

// DecodeDRBStatusULExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusULExtIEsEntry(data []byte) (DRBStatusULExtIEsEntry, int, error) {
	var v DRBStatusULExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBStatusUL-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusULExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
