// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationTNGFExtIEsEntryPresent uint32

const (
	UserLocationInformationTNGFExtIEsEntryPresentNothing UserLocationInformationTNGFExtIEsEntryPresent = iota
	UserLocationInformationTNGFExtIEsEntryPresentTAI
	UserLocationInformationTNGFExtIEsEntryPresentUnknown
)

type UserLocationInformationTNGFExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UserLocationInformationTNGFExtIEsEntryPresent
	TAI          TAI `ngap:"id:213,criticality:ignore,presence:optional"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationTNGFExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationTNGF-ExtIEsEntry", v)
}

// DecodeUserLocationInformationTNGFExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationTNGFExtIEsEntry(data []byte) (UserLocationInformationTNGFExtIEsEntry, int, error) {
	var v UserLocationInformationTNGFExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationTNGF-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationTNGFExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
