// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationN3IWFWithoutPortNumberExtIEsEntryPresent uint32

const (
	UserLocationInformationN3IWFWithoutPortNumberExtIEsEntryPresentNothing UserLocationInformationN3IWFWithoutPortNumberExtIEsEntryPresent = iota
	UserLocationInformationN3IWFWithoutPortNumberExtIEsEntryPresentUnknown
)

type UserLocationInformationN3IWFWithoutPortNumberExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      UserLocationInformationN3IWFWithoutPortNumberExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationN3IWFWithoutPortNumberExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationN3IWF-without-PortNumber-ExtIEsEntry", v)
}

// DecodeUserLocationInformationN3IWFWithoutPortNumberExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationN3IWFWithoutPortNumberExtIEsEntry(data []byte) (UserLocationInformationN3IWFWithoutPortNumberExtIEsEntry, int, error) {
	var v UserLocationInformationN3IWFWithoutPortNumberExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationN3IWF-without-PortNumber-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationN3IWFWithoutPortNumberExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
