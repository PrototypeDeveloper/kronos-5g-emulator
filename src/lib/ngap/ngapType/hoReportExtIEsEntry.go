// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type HOReportExtIEsEntryPresent uint32

const (
	HOReportExtIEsEntryPresentNothing HOReportExtIEsEntryPresent = iota
	HOReportExtIEsEntryPresentExtendedMobilityInformation
	HOReportExtIEsEntryPresentTargetCellCRNTI
	HOReportExtIEsEntryPresentTimeSinceFailure
	HOReportExtIEsEntryPresentUnknown
)

type HOReportExtIEsEntry struct {
	ID                          ProtocolExtensionID
	Criticality                 Criticality
	Present                     HOReportExtIEsEntryPresent
	ExtendedMobilityInformation ExtendedMobilityInformation `ngap:"id:366,criticality:ignore,presence:optional"`
	TargetCellCRNTI             CRNTI                       `ngap:"id:385,criticality:ignore,presence:optional"`
	TimeSinceFailure            TimeSinceFailure            `ngap:"id:386,criticality:ignore,presence:optional"`
	UnknownValue                []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HOReportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("HOReport-ExtIEsEntry", v)
}

// DecodeHOReportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeHOReportExtIEsEntry(data []byte) (HOReportExtIEsEntry, int, error) {
	var v HOReportExtIEsEntry
	err := ngapcodec.DecodeParameter("HOReport-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return HOReportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
