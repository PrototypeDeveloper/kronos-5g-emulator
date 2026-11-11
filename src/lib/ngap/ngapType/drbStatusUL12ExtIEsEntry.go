// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBStatusUL12ExtIEsEntryPresent uint32

const (
	DRBStatusUL12ExtIEsEntryPresentNothing DRBStatusUL12ExtIEsEntryPresent = iota
	DRBStatusUL12ExtIEsEntryPresentUnknown
)

type DRBStatusUL12ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DRBStatusUL12ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusUL12ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBStatusUL12-ExtIEsEntry", v)
}

// DecodeDRBStatusUL12ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusUL12ExtIEsEntry(data []byte) (DRBStatusUL12ExtIEsEntry, int, error) {
	var v DRBStatusUL12ExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBStatusUL12-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusUL12ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
