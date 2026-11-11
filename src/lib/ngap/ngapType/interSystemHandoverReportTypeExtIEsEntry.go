// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InterSystemHandoverReportTypeExtIEsEntryPresent uint32

const (
	InterSystemHandoverReportTypeExtIEsEntryPresentNothing InterSystemHandoverReportTypeExtIEsEntryPresent = iota
	InterSystemHandoverReportTypeExtIEsEntryPresentIntersystemMobilityFailureforVoiceFallback
	InterSystemHandoverReportTypeExtIEsEntryPresentUnknown
)

type InterSystemHandoverReportTypeExtIEsEntry struct {
	ID                                         ProtocolIEID
	Criticality                                Criticality
	Present                                    InterSystemHandoverReportTypeExtIEsEntryPresent
	IntersystemMobilityFailureforVoiceFallback IntersystemMobilityFailureforVoiceFallback `ngap:"id:384,criticality:ignore,presence:mandatory"`
	UnknownValue                               []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InterSystemHandoverReportTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InterSystemHandoverReportType-ExtIEsEntry", v)
}

// DecodeInterSystemHandoverReportTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInterSystemHandoverReportTypeExtIEsEntry(data []byte) (InterSystemHandoverReportTypeExtIEsEntry, int, error) {
	var v InterSystemHandoverReportTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("InterSystemHandoverReportType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InterSystemHandoverReportTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
