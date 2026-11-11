// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemResourceStatusRequestExtIEsEntryPresent uint32

const (
	IntersystemResourceStatusRequestExtIEsEntryPresentNothing IntersystemResourceStatusRequestExtIEsEntryPresent = iota
	IntersystemResourceStatusRequestExtIEsEntryPresentUnknown
)

type IntersystemResourceStatusRequestExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemResourceStatusRequestExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceStatusRequestExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceStatusRequest-ExtIEsEntry", v)
}

// DecodeIntersystemResourceStatusRequestExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceStatusRequestExtIEsEntry(data []byte) (IntersystemResourceStatusRequestExtIEsEntry, int, error) {
	var v IntersystemResourceStatusRequestExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemResourceStatusRequest-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceStatusRequestExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
