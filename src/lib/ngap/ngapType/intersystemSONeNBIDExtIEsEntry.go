// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONeNBIDExtIEsEntryPresent uint32

const (
	IntersystemSONeNBIDExtIEsEntryPresentNothing IntersystemSONeNBIDExtIEsEntryPresent = iota
	IntersystemSONeNBIDExtIEsEntryPresentUnknown
)

type IntersystemSONeNBIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemSONeNBIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONeNBIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONeNBID-ExtIEsEntry", v)
}

// DecodeIntersystemSONeNBIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONeNBIDExtIEsEntry(data []byte) (IntersystemSONeNBIDExtIEsEntry, int, error) {
	var v IntersystemSONeNBIDExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONeNBID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONeNBIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
