// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SecondaryRATDataUsageReportTransferExtIEsEntryPresent uint32

const (
	SecondaryRATDataUsageReportTransferExtIEsEntryPresentNothing SecondaryRATDataUsageReportTransferExtIEsEntryPresent = iota
	SecondaryRATDataUsageReportTransferExtIEsEntryPresentUnknown
)

type SecondaryRATDataUsageReportTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SecondaryRATDataUsageReportTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecondaryRATDataUsageReportTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecondaryRATDataUsageReportTransfer-ExtIEsEntry", v)
}

// DecodeSecondaryRATDataUsageReportTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSecondaryRATDataUsageReportTransferExtIEsEntry(data []byte) (SecondaryRATDataUsageReportTransferExtIEsEntry, int, error) {
	var v SecondaryRATDataUsageReportTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("SecondaryRATDataUsageReportTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecondaryRATDataUsageReportTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
