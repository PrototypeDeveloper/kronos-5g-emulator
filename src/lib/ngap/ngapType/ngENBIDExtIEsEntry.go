// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NgENBIDExtIEsEntryPresent uint32

const (
	NgENBIDExtIEsEntryPresentNothing NgENBIDExtIEsEntryPresent = iota
	NgENBIDExtIEsEntryPresentUnknown
)

type NgENBIDExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NgENBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NgENBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NgENB-ID-ExtIEsEntry", v)
}

// DecodeNgENBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNgENBIDExtIEsEntry(data []byte) (NgENBIDExtIEsEntry, int, error) {
	var v NgENBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("NgENB-ID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NgENBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
