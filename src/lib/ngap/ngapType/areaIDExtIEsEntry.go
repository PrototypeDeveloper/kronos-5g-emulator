// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AreaIDExtIEsEntryPresent uint32

const (
	AreaIDExtIEsEntryPresentNothing AreaIDExtIEsEntryPresent = iota
	AreaIDExtIEsEntryPresentUnknown
)

type AreaIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      AreaIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Area-ID-ExtIEsEntry", v)
}

// DecodeAreaIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAreaIDExtIEsEntry(data []byte) (AreaIDExtIEsEntry, int, error) {
	var v AreaIDExtIEsEntry
	err := ngapcodec.DecodeParameter("Area-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
