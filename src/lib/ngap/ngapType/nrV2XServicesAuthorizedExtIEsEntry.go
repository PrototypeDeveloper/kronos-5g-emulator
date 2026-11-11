// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRV2XServicesAuthorizedExtIEsEntryPresent uint32

const (
	NRV2XServicesAuthorizedExtIEsEntryPresentNothing NRV2XServicesAuthorizedExtIEsEntryPresent = iota
	NRV2XServicesAuthorizedExtIEsEntryPresentUnknown
)

type NRV2XServicesAuthorizedExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRV2XServicesAuthorizedExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRV2XServicesAuthorizedExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NRV2XServicesAuthorized-ExtIEsEntry", v)
}

// DecodeNRV2XServicesAuthorizedExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRV2XServicesAuthorizedExtIEsEntry(data []byte) (NRV2XServicesAuthorizedExtIEsEntry, int, error) {
	var v NRV2XServicesAuthorizedExtIEsEntry
	err := ngapcodec.DecodeParameter("NRV2XServicesAuthorized-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRV2XServicesAuthorizedExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
