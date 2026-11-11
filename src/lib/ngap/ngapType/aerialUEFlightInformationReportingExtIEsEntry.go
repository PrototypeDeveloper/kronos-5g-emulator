// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type AerialUEFlightInformationReportingExtIEsEntryPresent uint32

const (
	AerialUEFlightInformationReportingExtIEsEntryPresentNothing AerialUEFlightInformationReportingExtIEsEntryPresent = iota
	AerialUEFlightInformationReportingExtIEsEntryPresentUnknown
)

type AerialUEFlightInformationReportingExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      AerialUEFlightInformationReportingExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AerialUEFlightInformationReportingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Aerial-UE-FlightInformationReporting-ExtIEsEntry", v)
}

// DecodeAerialUEFlightInformationReportingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeAerialUEFlightInformationReportingExtIEsEntry(data []byte) (AerialUEFlightInformationReportingExtIEsEntry, int, error) {
	var v AerialUEFlightInformationReportingExtIEsEntry
	err := ngapcodec.DecodeParameter("Aerial-UE-FlightInformationReporting-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return AerialUEFlightInformationReportingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
