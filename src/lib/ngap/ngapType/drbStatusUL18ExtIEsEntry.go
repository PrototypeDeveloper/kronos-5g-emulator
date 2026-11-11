// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBStatusUL18ExtIEsEntryPresent uint32

const (
	DRBStatusUL18ExtIEsEntryPresentNothing DRBStatusUL18ExtIEsEntryPresent = iota
	DRBStatusUL18ExtIEsEntryPresentUnknown
)

type DRBStatusUL18ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DRBStatusUL18ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusUL18ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBStatusUL18-ExtIEsEntry", v)
}

// DecodeDRBStatusUL18ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusUL18ExtIEsEntry(data []byte) (DRBStatusUL18ExtIEsEntry, int, error) {
	var v DRBStatusUL18ExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBStatusUL18-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusUL18ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
