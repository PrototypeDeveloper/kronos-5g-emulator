// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NPNMobilityInformationExtIEsEntryPresent uint32

const (
	NPNMobilityInformationExtIEsEntryPresentNothing NPNMobilityInformationExtIEsEntryPresent = iota
	NPNMobilityInformationExtIEsEntryPresentUnknown
)

type NPNMobilityInformationExtIEsEntry struct {
	ID           ProtocolIEID
	Criticality  Criticality
	Present      NPNMobilityInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NPNMobilityInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NPN-MobilityInformation-ExtIEsEntry", v)
}

// DecodeNPNMobilityInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNPNMobilityInformationExtIEsEntry(data []byte) (NPNMobilityInformationExtIEsEntry, int, error) {
	var v NPNMobilityInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("NPN-MobilityInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NPNMobilityInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
