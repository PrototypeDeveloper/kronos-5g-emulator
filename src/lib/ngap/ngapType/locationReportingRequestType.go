// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// LocationReportingRequestType represents ASN.1 LocationReportingRequestType.
type LocationReportingRequestType struct {
	EventType                                    EventType
	ReportArea                                   ReportArea
	HasAreaOfInterestList                        bool
	AreaOfInterestList                           AreaOfInterestList
	HasLocationReportingReferenceIDToBeCancelled bool
	LocationReportingReferenceIDToBeCancelled    LocationReportingReferenceID
	HasIEExtensions                              bool
	IEExtensions                                 []LocationReportingRequestTypeExtIEsEntry
	UnknownExtensions                            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LocationReportingRequestType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LocationReportingRequestType", v)
}

// DecodeLocationReportingRequestType decodes a standalone APER value and reports the consumed octets.
func DecodeLocationReportingRequestType(data []byte) (LocationReportingRequestType, int, error) {
	var v LocationReportingRequestType
	err := ngapcodec.DecodeParameter("LocationReportingRequestType", data, &v)
	if ngapcodec.IsFatal(err) {
		return LocationReportingRequestType{}, 0, err
	}
	return v, len(data), err
}
