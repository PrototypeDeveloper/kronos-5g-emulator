// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UERLFReportContainerExtIEsEntryPresent uint32

const (
	UERLFReportContainerExtIEsEntryPresentNothing UERLFReportContainerExtIEsEntryPresent = iota
	UERLFReportContainerExtIEsEntryPresentUnknown
)

type UERLFReportContainerExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      UERLFReportContainerExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERLFReportContainerExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERLFReportContainer-ExtIEsEntry", v)
}

// DecodeUERLFReportContainerExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUERLFReportContainerExtIEsEntry(data []byte) (UERLFReportContainerExtIEsEntry, int, error) {
	var v UERLFReportContainerExtIEsEntry
	err := ngapcodec.DecodeParameter("UERLFReportContainer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERLFReportContainerExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
