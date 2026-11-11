// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserPlaneSecurityInformationExtIEsEntryPresent uint32

const (
	UserPlaneSecurityInformationExtIEsEntryPresentNothing UserPlaneSecurityInformationExtIEsEntryPresent = iota
	UserPlaneSecurityInformationExtIEsEntryPresentUnknown
)

type UserPlaneSecurityInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UserPlaneSecurityInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneSecurityInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneSecurityInformation-ExtIEsEntry", v)
}

// DecodeUserPlaneSecurityInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneSecurityInformationExtIEsEntry(data []byte) (UserPlaneSecurityInformationExtIEsEntry, int, error) {
	var v UserPlaneSecurityInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("UserPlaneSecurityInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneSecurityInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
