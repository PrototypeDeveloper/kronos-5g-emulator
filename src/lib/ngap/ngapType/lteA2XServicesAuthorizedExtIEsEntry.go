// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LTEA2XServicesAuthorizedExtIEsEntryPresent uint32

const (
	LTEA2XServicesAuthorizedExtIEsEntryPresentNothing LTEA2XServicesAuthorizedExtIEsEntryPresent = iota
	LTEA2XServicesAuthorizedExtIEsEntryPresentUnknown
)

type LTEA2XServicesAuthorizedExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      LTEA2XServicesAuthorizedExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEA2XServicesAuthorizedExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTE-A2X-ServicesAuthorized-ExtIEsEntry", v)
}

// DecodeLTEA2XServicesAuthorizedExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLTEA2XServicesAuthorizedExtIEsEntry(data []byte) (LTEA2XServicesAuthorizedExtIEsEntry, int, error) {
	var v LTEA2XServicesAuthorizedExtIEsEntry
	err := ngapcodec.DecodeParameter("LTE-A2X-ServicesAuthorized-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEA2XServicesAuthorizedExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
