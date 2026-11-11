// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONTransferTypeExtIEsEntryPresent uint32

const (
	IntersystemSONTransferTypeExtIEsEntryPresentNothing IntersystemSONTransferTypeExtIEsEntryPresent = iota
	IntersystemSONTransferTypeExtIEsEntryPresentUnknown
)

type IntersystemSONTransferTypeExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      IntersystemSONTransferTypeExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONTransferTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONTransferType-ExtIEsEntry", v)
}

// DecodeIntersystemSONTransferTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONTransferTypeExtIEsEntry(data []byte) (IntersystemSONTransferTypeExtIEsEntry, int, error) {
	var v IntersystemSONTransferTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONTransferType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONTransferTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
