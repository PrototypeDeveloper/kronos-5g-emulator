// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PDUSessionUsageReportExtIEsEntryPresent uint32

const (
	PDUSessionUsageReportExtIEsEntryPresentNothing PDUSessionUsageReportExtIEsEntryPresent = iota
	PDUSessionUsageReportExtIEsEntryPresentUnknown
)

type PDUSessionUsageReportExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PDUSessionUsageReportExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PDUSessionUsageReportExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PDUSessionUsageReport-ExtIEsEntry", v)
}

// DecodePDUSessionUsageReportExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePDUSessionUsageReportExtIEsEntry(data []byte) (PDUSessionUsageReportExtIEsEntry, int, error) {
	var v PDUSessionUsageReportExtIEsEntry
	err := ngapcodec.DecodeParameter("PDUSessionUsageReport-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PDUSessionUsageReportExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
