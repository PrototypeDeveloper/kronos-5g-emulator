// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRA2XServicesAuthorizedExtIEsEntryPresent uint32

const (
	NRA2XServicesAuthorizedExtIEsEntryPresentNothing NRA2XServicesAuthorizedExtIEsEntryPresent = iota
	NRA2XServicesAuthorizedExtIEsEntryPresentUnknown
)

type NRA2XServicesAuthorizedExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRA2XServicesAuthorizedExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRA2XServicesAuthorizedExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-A2X-ServicesAuthorized-ExtIEsEntry", v)
}

// DecodeNRA2XServicesAuthorizedExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRA2XServicesAuthorizedExtIEsEntry(data []byte) (NRA2XServicesAuthorizedExtIEsEntry, int, error) {
	var v NRA2XServicesAuthorizedExtIEsEntry
	err := ngapcodec.DecodeParameter("NR-A2X-ServicesAuthorized-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRA2XServicesAuthorizedExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
