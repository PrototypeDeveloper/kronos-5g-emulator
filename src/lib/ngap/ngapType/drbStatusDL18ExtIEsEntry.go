// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBStatusDL18ExtIEsEntryPresent uint32

const (
	DRBStatusDL18ExtIEsEntryPresentNothing DRBStatusDL18ExtIEsEntryPresent = iota
	DRBStatusDL18ExtIEsEntryPresentUnknown
)

type DRBStatusDL18ExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      DRBStatusDL18ExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusDL18ExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBStatusDL18-ExtIEsEntry", v)
}

// DecodeDRBStatusDL18ExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusDL18ExtIEsEntry(data []byte) (DRBStatusDL18ExtIEsEntry, int, error) {
	var v DRBStatusDL18ExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBStatusDL18-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusDL18ExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
