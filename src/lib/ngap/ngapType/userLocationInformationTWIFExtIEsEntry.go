// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationTWIFExtIEsEntryPresent uint32

const (
	UserLocationInformationTWIFExtIEsEntryPresentNothing UserLocationInformationTWIFExtIEsEntryPresent = iota
	UserLocationInformationTWIFExtIEsEntryPresentTAI
	UserLocationInformationTWIFExtIEsEntryPresentUnknown
)

type UserLocationInformationTWIFExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UserLocationInformationTWIFExtIEsEntryPresent
	TAI          TAI `ngap:"id:213,criticality:ignore,presence:optional"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationTWIFExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationTWIF-ExtIEsEntry", v)
}

// DecodeUserLocationInformationTWIFExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationTWIFExtIEsEntry(data []byte) (UserLocationInformationTWIFExtIEsEntry, int, error) {
	var v UserLocationInformationTWIFExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationTWIF-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationTWIFExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
