// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LocationReportingRequestTypeExtIEsEntryPresent uint32

const (
	LocationReportingRequestTypeExtIEsEntryPresentNothing LocationReportingRequestTypeExtIEsEntryPresent = iota
	LocationReportingRequestTypeExtIEsEntryPresentLocationReportingAdditionalInfo
	LocationReportingRequestTypeExtIEsEntryPresentAdditionalCancelledlocationReportingReferenceIDList
	LocationReportingRequestTypeExtIEsEntryPresentAerialUEFlightInformationReportingControlList
	LocationReportingRequestTypeExtIEsEntryPresentUnknown
)

type LocationReportingRequestTypeExtIEsEntry struct {
	ID                                                  ProtocolExtensionID
	Criticality                                         Criticality
	Present                                             LocationReportingRequestTypeExtIEsEntryPresent
	LocationReportingAdditionalInfo                     LocationReportingAdditionalInfo                     `ngap:"id:170,criticality:ignore,presence:optional"`
	AdditionalCancelledlocationReportingReferenceIDList AdditionalCancelledlocationReportingReferenceIDList `ngap:"id:368,criticality:reject,presence:optional"`
	AerialUEFlightInformationReportingControlList       AerialUEFlightInformationReportingControlList       `ngap:"id:494,criticality:ignore,presence:optional"`
	UnknownValue                                        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LocationReportingRequestTypeExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LocationReportingRequestType-ExtIEsEntry", v)
}

// DecodeLocationReportingRequestTypeExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLocationReportingRequestTypeExtIEsEntry(data []byte) (LocationReportingRequestTypeExtIEsEntry, int, error) {
	var v LocationReportingRequestTypeExtIEsEntry
	err := ngapcodec.DecodeParameter("LocationReportingRequestType-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LocationReportingRequestTypeExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
