// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PEIPSassistanceInformationExtIEsEntryPresent uint32

const (
	PEIPSassistanceInformationExtIEsEntryPresentNothing PEIPSassistanceInformationExtIEsEntryPresent = iota
	PEIPSassistanceInformationExtIEsEntryPresentUnknown
)

type PEIPSassistanceInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PEIPSassistanceInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PEIPSassistanceInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PEIPSassistanceInformation-ExtIEsEntry", v)
}

// DecodePEIPSassistanceInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePEIPSassistanceInformationExtIEsEntry(data []byte) (PEIPSassistanceInformationExtIEsEntry, int, error) {
	var v PEIPSassistanceInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("PEIPSassistanceInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PEIPSassistanceInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
