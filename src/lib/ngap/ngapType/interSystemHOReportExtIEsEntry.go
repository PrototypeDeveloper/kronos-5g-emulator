// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InterSystemHOReportExtIEsEntryPresent uint32

const (
	InterSystemHOReportExtIEsEntryPresentNothing InterSystemHOReportExtIEsEntryPresent = iota
	InterSystemHOReportExtIEsEntryPresentUnknown
)

type InterSystemHOReportExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      InterSystemHOReportExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterSystemHOReportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterSystemHOReport-ExtIEsEntry", v)
}

// DecodeInterSystemHOReportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInterSystemHOReportExtIEsEntry(data []byte) (InterSystemHOReportExtIEsEntry, int, error) {
	var v InterSystemHOReportExtIEsEntry
	err := ngapcodec.DecodeParameter("InterSystemHOReport-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterSystemHOReportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
