// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SuccessfulHandoverReportItemExtIEsEntryPresent uint32

const (
	SuccessfulHandoverReportItemExtIEsEntryPresentNothing SuccessfulHandoverReportItemExtIEsEntryPresent = iota
	SuccessfulHandoverReportItemExtIEsEntryPresentUnknown
)

type SuccessfulHandoverReportItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SuccessfulHandoverReportItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SuccessfulHandoverReportItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SuccessfulHandoverReport-Item-ExtIEsEntry", v)
}

// DecodeSuccessfulHandoverReportItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSuccessfulHandoverReportItemExtIEsEntry(data []byte) (SuccessfulHandoverReportItemExtIEsEntry, int, error) {
	var v SuccessfulHandoverReportItemExtIEsEntry
	err := ngapcodec.DecodeParameter("SuccessfulHandoverReport-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SuccessfulHandoverReportItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
