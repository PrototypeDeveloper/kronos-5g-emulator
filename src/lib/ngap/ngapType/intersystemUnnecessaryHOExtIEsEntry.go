// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemUnnecessaryHOExtIEsEntryPresent uint32

const (
	IntersystemUnnecessaryHOExtIEsEntryPresentNothing IntersystemUnnecessaryHOExtIEsEntryPresent = iota
	IntersystemUnnecessaryHOExtIEsEntryPresentUnknown
)

type IntersystemUnnecessaryHOExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemUnnecessaryHOExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemUnnecessaryHOExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemUnnecessaryHO-ExtIEsEntry", v)
}

// DecodeIntersystemUnnecessaryHOExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemUnnecessaryHOExtIEsEntry(data []byte) (IntersystemUnnecessaryHOExtIEsEntry, int, error) {
	var v IntersystemUnnecessaryHOExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemUnnecessaryHO-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemUnnecessaryHOExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
