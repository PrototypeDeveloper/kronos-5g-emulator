// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type DRBStatusDLExtIEsEntryPresent uint32

const (
	DRBStatusDLExtIEsEntryPresentNothing DRBStatusDLExtIEsEntryPresent = iota
	DRBStatusDLExtIEsEntryPresentUnknown
)

type DRBStatusDLExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      DRBStatusDLExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusDLExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DRBStatusDL-ExtIEsEntry", v)
}

// DecodeDRBStatusDLExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusDLExtIEsEntry(data []byte) (DRBStatusDLExtIEsEntry, int, error) {
	var v DRBStatusDLExtIEsEntry
	err := ngapcodec.DecodeParameter("DRBStatusDL-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusDLExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
