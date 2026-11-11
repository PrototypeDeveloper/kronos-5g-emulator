// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEFlightInformationReporting represents ASN.1 Aerial-UE-FlightInformationReporting.
type AerialUEFlightInformationReporting struct {
	Altitude          Altitude
	TimeStamp         TimeStamp
	HasIEExtensions   bool
	IEExtensions      []AerialUEFlightInformationReportingExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReporting) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReporting", v)
}

// DecodeAerialUEFlightInformationReporting decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReporting(data []byte) (AerialUEFlightInformationReporting, int, error) {
	var v AerialUEFlightInformationReporting
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReporting", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReporting{}, 0, err
	}
	return v, len(data), err
}
