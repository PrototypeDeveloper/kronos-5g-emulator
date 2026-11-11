// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ForbiddenAreaInformationItemExtIEsEntryPresent uint32

const (
	ForbiddenAreaInformationItemExtIEsEntryPresentNothing ForbiddenAreaInformationItemExtIEsEntryPresent = iota
	ForbiddenAreaInformationItemExtIEsEntryPresentUnknown
)

type ForbiddenAreaInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ForbiddenAreaInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ForbiddenAreaInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ForbiddenAreaInformation-Item-ExtIEsEntry", v)
}

// DecodeForbiddenAreaInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeForbiddenAreaInformationItemExtIEsEntry(data []byte) (ForbiddenAreaInformationItemExtIEsEntry, int, error) {
	var v ForbiddenAreaInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("ForbiddenAreaInformation-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ForbiddenAreaInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
