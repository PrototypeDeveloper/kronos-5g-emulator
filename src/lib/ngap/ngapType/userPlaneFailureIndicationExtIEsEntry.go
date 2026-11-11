// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserPlaneFailureIndicationExtIEsEntryPresent uint32

const (
	UserPlaneFailureIndicationExtIEsEntryPresentNothing UserPlaneFailureIndicationExtIEsEntryPresent = iota
	UserPlaneFailureIndicationExtIEsEntryPresentUnknown
)

type UserPlaneFailureIndicationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UserPlaneFailureIndicationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserPlaneFailureIndicationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserPlaneFailureIndication-ExtIEsEntry", v)
}

// DecodeUserPlaneFailureIndicationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserPlaneFailureIndicationExtIEsEntry(data []byte) (UserPlaneFailureIndicationExtIEsEntry, int, error) {
	var v UserPlaneFailureIndicationExtIEsEntry
	err := ngapcodec.DecodeParameter("UserPlaneFailureIndication-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserPlaneFailureIndicationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
