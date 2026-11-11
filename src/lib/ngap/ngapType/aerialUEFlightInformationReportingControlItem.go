// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AerialUEFlightInformationReportingControlItem represents ASN.1 Aerial-UE-FlightInformationReportingControlItem.
type AerialUEFlightInformationReportingControlItem struct {
	AerialUEFlightInformationReportingControl AerialUEFlightInformationReportingControl
	HasIEExtensions                           bool
	IEExtensions                              []AerialUEFlightInformationReportingControlItemExtIEsEntry
	UnknownExtensions                         []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReportingControlItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReportingControlItem", v)
}

// DecodeAerialUEFlightInformationReportingControlItem decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReportingControlItem(data []byte) (AerialUEFlightInformationReportingControlItem, int, error) {
	var v AerialUEFlightInformationReportingControlItem
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReportingControlItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReportingControlItem{}, 0, err
	}
	return v, len(data), err
}
