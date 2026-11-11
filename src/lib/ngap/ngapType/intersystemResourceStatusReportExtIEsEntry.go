// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemResourceStatusReportExtIEsEntryPresent uint32

const (
	IntersystemResourceStatusReportExtIEsEntryPresentNothing IntersystemResourceStatusReportExtIEsEntryPresent = iota
	IntersystemResourceStatusReportExtIEsEntryPresentUnknown
)

type IntersystemResourceStatusReportExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      IntersystemResourceStatusReportExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemResourceStatusReportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemResourceStatusReport-ExtIEsEntry", v)
}

// DecodeIntersystemResourceStatusReportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemResourceStatusReportExtIEsEntry(data []byte) (IntersystemResourceStatusReportExtIEsEntry, int, error) {
	var v IntersystemResourceStatusReportExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemResourceStatusReport-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemResourceStatusReportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
