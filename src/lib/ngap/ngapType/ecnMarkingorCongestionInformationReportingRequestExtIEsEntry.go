// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ECNMarkingorCongestionInformationReportingRequestExtIEsEntryPresent uint32

const (
	ECNMarkingorCongestionInformationReportingRequestExtIEsEntryPresentNothing ECNMarkingorCongestionInformationReportingRequestExtIEsEntryPresent = iota
	ECNMarkingorCongestionInformationReportingRequestExtIEsEntryPresentUnknown
)

type ECNMarkingorCongestionInformationReportingRequestExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      ECNMarkingorCongestionInformationReportingRequestExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingorCongestionInformationReportingRequestExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingorCongestionInformationReportingRequest-ExtIEsEntry", v)
}

// DecodeECNMarkingorCongestionInformationReportingRequestExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingorCongestionInformationReportingRequestExtIEsEntry(data []byte) (ECNMarkingorCongestionInformationReportingRequestExtIEsEntry, int, error) {
	var v ECNMarkingorCongestionInformationReportingRequestExtIEsEntry
	err := ngapcodec.DecodeParameter("ECNMarkingorCongestionInformationReportingRequest-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingorCongestionInformationReportingRequestExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
