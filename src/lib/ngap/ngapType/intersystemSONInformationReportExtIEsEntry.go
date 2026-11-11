// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type IntersystemSONInformationReportExtIEsEntryPresent uint32

const (
	IntersystemSONInformationReportExtIEsEntryPresentNothing IntersystemSONInformationReportExtIEsEntryPresent = iota
	IntersystemSONInformationReportExtIEsEntryPresentEnergySavingIndication
	IntersystemSONInformationReportExtIEsEntryPresentIntersystemResourceStatusUpdate
	IntersystemSONInformationReportExtIEsEntryPresentUnknown
)

type IntersystemSONInformationReportExtIEsEntry struct {
	ID                              ProtocolIEID
	Criticality                     Criticality
	Present                         IntersystemSONInformationReportExtIEsEntryPresent
	EnergySavingIndication          IntersystemCellStateIndication  `ngap:"id:292,criticality:ignore,presence:mandatory"`
	IntersystemResourceStatusUpdate IntersystemResourceStatusReport `ngap:"id:293,criticality:ignore,presence:mandatory"`
	UnknownValue                    []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemSONInformationReportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemSONInformationReport-ExtIEsEntry", v)
}

// DecodeIntersystemSONInformationReportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemSONInformationReportExtIEsEntry(data []byte) (IntersystemSONInformationReportExtIEsEntry, int, error) {
	var v IntersystemSONInformationReportExtIEsEntry
	err := ngapcodec.DecodeParameter("IntersystemSONInformationReport-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemSONInformationReportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
