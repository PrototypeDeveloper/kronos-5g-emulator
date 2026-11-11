// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AerialUEFlightInformationReportingControlItemExtIEsEntryPresent uint32

const (
	AerialUEFlightInformationReportingControlItemExtIEsEntryPresentNothing AerialUEFlightInformationReportingControlItemExtIEsEntryPresent = iota
	AerialUEFlightInformationReportingControlItemExtIEsEntryPresentUnknown
)

type AerialUEFlightInformationReportingControlItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AerialUEFlightInformationReportingControlItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReportingControlItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReportingControlItem-ExtIEsEntry", v)
}

// DecodeAerialUEFlightInformationReportingControlItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReportingControlItemExtIEsEntry(data []byte) (AerialUEFlightInformationReportingControlItemExtIEsEntry, int, error) {
	var v AerialUEFlightInformationReportingControlItemExtIEsEntry
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReportingControlItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReportingControlItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
