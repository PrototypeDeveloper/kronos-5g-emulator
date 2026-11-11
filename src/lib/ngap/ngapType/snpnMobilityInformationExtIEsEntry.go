// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SNPNMobilityInformationExtIEsEntryPresent uint32

const (
	SNPNMobilityInformationExtIEsEntryPresentNothing SNPNMobilityInformationExtIEsEntryPresent = iota
	SNPNMobilityInformationExtIEsEntryPresentEquivalentSNPNsList
	SNPNMobilityInformationExtIEsEntryPresentUnknown
)

type SNPNMobilityInformationExtIEsEntry struct {
	ID                  ProtocolExtensionID
	Criticality         Criticality
	Present             SNPNMobilityInformationExtIEsEntryPresent
	EquivalentSNPNsList EquivalentSNPNsList `ngap:"id:370,criticality:reject,presence:optional"`
	UnknownValue        []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNMobilityInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPN-MobilityInformation-ExtIEsEntry", v)
}

// DecodeSNPNMobilityInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNMobilityInformationExtIEsEntry(data []byte) (SNPNMobilityInformationExtIEsEntry, int, error) {
	var v SNPNMobilityInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("SNPN-MobilityInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNMobilityInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
