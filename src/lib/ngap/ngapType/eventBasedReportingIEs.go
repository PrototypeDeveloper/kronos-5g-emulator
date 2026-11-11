// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EventBasedReportingIEs represents ASN.1 EventBasedReportingIEs.
type EventBasedReportingIEs struct {
	IntersystemResourceThresholdLow    IntersystemResourceThreshold
	IntersystemResourceThresholdHigh   IntersystemResourceThreshold
	NumberOfMeasurementReportingLevels NumberOfMeasurementReportingLevels
	HasIEExtensions                    bool
	IEExtensions                       []EventBasedReportingIEsExtIEsEntry
	UnknownExtensions                  []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EventBasedReportingIEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EventBasedReportingIEs", v)
}

// DecodeEventBasedReportingIEs decodes a standalone APER value and reports the consumed octets.
func DecodeEventBasedReportingIEs(data []byte) (EventBasedReportingIEs, int, error) {
	var v EventBasedReportingIEs
	err := ngapcodec.DecodeParameter("EventBasedReportingIEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return EventBasedReportingIEs{}, 0, err
	}
	return v, len(data), err
}
