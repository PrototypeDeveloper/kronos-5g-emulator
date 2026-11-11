// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PrivateMessageIEsEntryPresent uint32

const (
	PrivateMessageIEsEntryPresentNothing PrivateMessageIEsEntryPresent = iota
	PrivateMessageIEsEntryPresentUnknown
)

type PrivateMessageIEsEntry struct {
	ID           PrivateIEID
	Criticality  Criticality
	Present      PrivateMessageIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PrivateMessageIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PrivateMessageIEsEntry", v)
}

// DecodePrivateMessageIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePrivateMessageIEsEntry(data []byte) (PrivateMessageIEsEntry, int, error) {
	var v PrivateMessageIEsEntry
	err := ngapcodec.DecodeParameter("PrivateMessageIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PrivateMessageIEsEntry{}, 0, err
	}
	return v, len(data), err
}
