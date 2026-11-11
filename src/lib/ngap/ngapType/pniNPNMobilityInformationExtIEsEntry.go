// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PNINPNMobilityInformationExtIEsEntryPresent uint32

const (
	PNINPNMobilityInformationExtIEsEntryPresentNothing PNINPNMobilityInformationExtIEsEntryPresent = iota
	PNINPNMobilityInformationExtIEsEntryPresentUnknown
)

type PNINPNMobilityInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PNINPNMobilityInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PNINPNMobilityInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PNI-NPN-MobilityInformation-ExtIEsEntry", v)
}

// DecodePNINPNMobilityInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePNINPNMobilityInformationExtIEsEntry(data []byte) (PNINPNMobilityInformationExtIEsEntry, int, error) {
	var v PNINPNMobilityInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("PNI-NPN-MobilityInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PNINPNMobilityInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
