// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CauseExtIEsEntryPresent uint32

const (
	CauseExtIEsEntryPresentNothing CauseExtIEsEntryPresent = iota
	CauseExtIEsEntryPresentUnknown
)

type CauseExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      CauseExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CauseExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Cause-ExtIEsEntry", v)
}

// DecodeCauseExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCauseExtIEsEntry(data []byte) (CauseExtIEsEntry, int, error) {
	var v CauseExtIEsEntry
	err := ngapcodec.DecodeParameter("Cause-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CauseExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
