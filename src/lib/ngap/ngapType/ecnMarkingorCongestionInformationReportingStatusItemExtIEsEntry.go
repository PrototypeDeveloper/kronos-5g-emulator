// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntryPresent uint32

const (
	ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntryPresentNothing ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntryPresent = iota
	ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntryPresentUnknown
)

type ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ECNMarkingorCongestionInformationReportingStatus-Item-ExtIEsEntry", v)
}

// DecodeECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry(data []byte) (ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry, int, error) {
	var v ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry
	err := ngapcodec.DecodeParameter("ECNMarkingorCongestionInformationReportingStatus-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ECNMarkingorCongestionInformationReportingStatusItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
