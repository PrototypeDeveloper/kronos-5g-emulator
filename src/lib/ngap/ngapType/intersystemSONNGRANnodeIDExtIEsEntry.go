// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONNGRANnodeIDExtIEsEntryPresent uint32

const (
	IntersystemSONNGRANnodeIDExtIEsEntryPresentNothing IntersystemSONNGRANnodeIDExtIEsEntryPresent = iota
	IntersystemSONNGRANnodeIDExtIEsEntryPresentUnknown
)

type IntersystemSONNGRANnodeIDExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemSONNGRANnodeIDExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONNGRANnodeIDExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONNGRANnodeID-ExtIEsEntry", v)
}

// DecodeIntersystemSONNGRANnodeIDExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONNGRANnodeIDExtIEsEntry(data []byte) (IntersystemSONNGRANnodeIDExtIEsEntry, int, error) {
	var v IntersystemSONNGRANnodeIDExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONNGRANnodeID-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONNGRANnodeIDExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
