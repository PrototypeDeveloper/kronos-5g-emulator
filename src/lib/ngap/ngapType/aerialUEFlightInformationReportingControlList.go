// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEFlightInformationReportingControlList represents ASN.1 Aerial-UE-FlightInformationReportingControlList.
type AerialUEFlightInformationReportingControlList struct {
	List []AerialUEFlightInformationReportingControlItem `aper:"sizeLB:1,sizeUB:64"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReportingControlList) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReportingControlList", v)
}

// DecodeAerialUEFlightInformationReportingControlList decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReportingControlList(data []byte) (AerialUEFlightInformationReportingControlList, int, error) {
	var v AerialUEFlightInformationReportingControlList
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReportingControlList", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReportingControlList{}, 0, err
	}
	return v, len(data), err
}
