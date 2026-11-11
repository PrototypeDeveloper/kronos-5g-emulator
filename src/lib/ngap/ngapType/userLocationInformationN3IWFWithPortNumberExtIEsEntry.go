// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationN3IWFWithPortNumberExtIEsEntryPresent uint32

const (
	UserLocationInformationN3IWFWithPortNumberExtIEsEntryPresentNothing UserLocationInformationN3IWFWithPortNumberExtIEsEntryPresent = iota
	UserLocationInformationN3IWFWithPortNumberExtIEsEntryPresentTAI
	UserLocationInformationN3IWFWithPortNumberExtIEsEntryPresentUnknown
)

type UserLocationInformationN3IWFWithPortNumberExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UserLocationInformationN3IWFWithPortNumberExtIEsEntryPresent
	TAI          TAI `ngap:"id:213,criticality:ignore,presence:optional"`
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationN3IWFWithPortNumberExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationN3IWF-with-PortNumber-ExtIEsEntry", v)
}

// DecodeUserLocationInformationN3IWFWithPortNumberExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationN3IWFWithPortNumberExtIEsEntry(data []byte) (UserLocationInformationN3IWFWithPortNumberExtIEsEntry, int, error) {
	var v UserLocationInformationN3IWFWithPortNumberExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationN3IWF-with-PortNumber-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationN3IWFWithPortNumberExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
