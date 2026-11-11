// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBStatusDL12ExtIEsEntryPresent uint32

const (
	DRBStatusDL12ExtIEsEntryPresentNothing DRBStatusDL12ExtIEsEntryPresent = iota
	DRBStatusDL12ExtIEsEntryPresentUnknown
)

type DRBStatusDL12ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DRBStatusDL12ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusDL12ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBStatusDL12-ExtIEsEntry", v)
}

// DecodeDRBStatusDL12ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusDL12ExtIEsEntry(data []byte) (DRBStatusDL12ExtIEsEntry, int, error) {
	var v DRBStatusDL12ExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBStatusDL12-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusDL12ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
