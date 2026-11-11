// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEFlightInformationOngoingReportingFailed represents ASN.1 Aerial-UE-FlightInformationOngoingReporting-failed.
const (
	AerialUEFlightInformationOngoingReportingFailedPresentTrue Enumerated = 0
)

type AerialUEFlightInformationOngoingReportingFailed struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationOngoingReportingFailed) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationOngoingReporting-failed", v)
}

// DecodeAerialUEFlightInformationOngoingReportingFailed decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationOngoingReportingFailed(data []byte) (AerialUEFlightInformationOngoingReportingFailed, int, error) {
	var v AerialUEFlightInformationOngoingReportingFailed
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationOngoingReporting-failed", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationOngoingReportingFailed{}, 0, err
	}
	return v, len(data), err
}
