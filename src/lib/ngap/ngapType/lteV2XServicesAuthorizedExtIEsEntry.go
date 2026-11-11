// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type LTEV2XServicesAuthorizedExtIEsEntryPresent uint32

const (
	LTEV2XServicesAuthorizedExtIEsEntryPresentNothing LTEV2XServicesAuthorizedExtIEsEntryPresent = iota
	LTEV2XServicesAuthorizedExtIEsEntryPresentUnknown
)

type LTEV2XServicesAuthorizedExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      LTEV2XServicesAuthorizedExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v LTEV2XServicesAuthorizedExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("LTEV2XServicesAuthorized-ExtIEsEntry", v)
}

// DecodeLTEV2XServicesAuthorizedExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeLTEV2XServicesAuthorizedExtIEsEntry(data []byte) (LTEV2XServicesAuthorizedExtIEsEntry, int, error) {
	var v LTEV2XServicesAuthorizedExtIEsEntry
	err := ngapcodec.DecodeParameter("LTEV2XServicesAuthorized-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return LTEV2XServicesAuthorizedExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
