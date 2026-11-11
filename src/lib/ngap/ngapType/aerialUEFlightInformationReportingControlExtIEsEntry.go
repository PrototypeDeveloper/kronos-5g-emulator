// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AerialUEFlightInformationReportingControlExtIEsEntryPresent uint32

const (
	AerialUEFlightInformationReportingControlExtIEsEntryPresentNothing AerialUEFlightInformationReportingControlExtIEsEntryPresent = iota
	AerialUEFlightInformationReportingControlExtIEsEntryPresentUnknown
)

type AerialUEFlightInformationReportingControlExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AerialUEFlightInformationReportingControlExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReportingControlExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReportingControl-ExtIEsEntry", v)
}

// DecodeAerialUEFlightInformationReportingControlExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReportingControlExtIEsEntry(data []byte) (AerialUEFlightInformationReportingControlExtIEsEntry, int, error) {
	var v AerialUEFlightInformationReportingControlExtIEsEntry
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReportingControl-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReportingControlExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
