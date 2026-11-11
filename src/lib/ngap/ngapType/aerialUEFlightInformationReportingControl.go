// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEFlightInformationReportingControl represents ASN.1 Aerial-UE-FlightInformationReportingControl.
type AerialUEFlightInformationReportingControl struct {
	HigherAltitudeThreshold      Altitude
	LowerAltitudeThreshold       Altitude
	HasAerialUEReportPeriodicity bool
	AerialUEReportPeriodicity    AerialUEReportingPeriodicity
	AreaID                       AreaID
	HasIEExtensions              bool
	IEExtensions                 []AerialUEFlightInformationReportingControlExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReportingControl) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReportingControl", v)
}

// DecodeAerialUEFlightInformationReportingControl decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReportingControl(data []byte) (AerialUEFlightInformationReportingControl, int, error) {
	var v AerialUEFlightInformationReportingControl
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReportingControl", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReportingControl{}, 0, err
	}
	return v, len(data), err
}
