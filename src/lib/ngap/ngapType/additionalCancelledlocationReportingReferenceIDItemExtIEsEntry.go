// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntryPresent uint32

const (
	AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntryPresentNothing AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntryPresent = iota
	AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntryPresentUnknown
)

type AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AdditionalCancelledlocationReportingReferenceIDItem-ExtIEsEntry", v)
}

// DecodeAdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry(data []byte) (AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry, int, error) {
	var v AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry
	err := ngapcodec.DecodeParameter("AdditionalCancelledlocationReportingReferenceIDItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AdditionalCancelledlocationReportingReferenceIDItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
