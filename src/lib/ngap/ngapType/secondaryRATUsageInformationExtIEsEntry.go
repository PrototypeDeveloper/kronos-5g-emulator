// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type SecondaryRATUsageInformationExtIEsEntryPresent uint32

const (
	SecondaryRATUsageInformationExtIEsEntryPresentNothing SecondaryRATUsageInformationExtIEsEntryPresent = iota
	SecondaryRATUsageInformationExtIEsEntryPresentUnknown
)

type SecondaryRATUsageInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      SecondaryRATUsageInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SecondaryRATUsageInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SecondaryRATUsageInformation-ExtIEsEntry", v)
}

// DecodeSecondaryRATUsageInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeSecondaryRATUsageInformationExtIEsEntry(data []byte) (SecondaryRATUsageInformationExtIEsEntry, int, error) {
	var v SecondaryRATUsageInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("SecondaryRATUsageInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return SecondaryRATUsageInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
