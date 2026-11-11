// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SONInformationReportExtIEsEntryPresent uint32

const (
	SONInformationReportExtIEsEntryPresentNothing SONInformationReportExtIEsEntryPresent = iota
	SONInformationReportExtIEsEntryPresentSuccessfulHandoverReportList
	SONInformationReportExtIEsEntryPresentSuccessfulPSCellChangeReportList
	SONInformationReportExtIEsEntryPresentUnknown
)

type SONInformationReportExtIEsEntry struct {
	ID                               ProtocolIEID
	Criticality                      Criticality
	Present                          SONInformationReportExtIEsEntryPresent
	SuccessfulHandoverReportList     SuccessfulHandoverReportList     `ngap:"id:294,criticality:ignore,presence:mandatory"`
	SuccessfulPSCellChangeReportList SuccessfulPSCellChangeReportList `ngap:"id:383,criticality:ignore,presence:mandatory"`
	UnknownValue                     []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SONInformationReportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SONInformationReport-ExtIEsEntry", v)
}

// DecodeSONInformationReportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSONInformationReportExtIEsEntry(data []byte) (SONInformationReportExtIEsEntry, int, error) {
	var v SONInformationReportExtIEsEntry
	err := ngapcodec.DecodeParameter("SONInformationReport-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SONInformationReportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
