// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type OverloadResponseExtIEsEntryPresent uint32

const (
	OverloadResponseExtIEsEntryPresentNothing OverloadResponseExtIEsEntryPresent = iota
	OverloadResponseExtIEsEntryPresentUnknown
)

type OverloadResponseExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      OverloadResponseExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OverloadResponseExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OverloadResponse-ExtIEsEntry", v)
}

// DecodeOverloadResponseExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeOverloadResponseExtIEsEntry(data []byte) (OverloadResponseExtIEsEntry, int, error) {
	var v OverloadResponseExtIEsEntry
	err := ngapcodec.DecodeParameter("OverloadResponse-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return OverloadResponseExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
