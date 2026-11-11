// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ServiceAreaInformationItemExtIEsEntryPresent uint32

const (
	ServiceAreaInformationItemExtIEsEntryPresentNothing ServiceAreaInformationItemExtIEsEntryPresent = iota
	ServiceAreaInformationItemExtIEsEntryPresentUnknown
)

type ServiceAreaInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ServiceAreaInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServiceAreaInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ServiceAreaInformation-Item-ExtIEsEntry", v)
}

// DecodeServiceAreaInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeServiceAreaInformationItemExtIEsEntry(data []byte) (ServiceAreaInformationItemExtIEsEntry, int, error) {
	var v ServiceAreaInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("ServiceAreaInformation-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServiceAreaInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
