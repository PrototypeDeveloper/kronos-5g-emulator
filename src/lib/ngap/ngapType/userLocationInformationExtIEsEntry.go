// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationExtIEsEntryPresent uint32

const (
	UserLocationInformationExtIEsEntryPresentNothing UserLocationInformationExtIEsEntryPresent = iota
	UserLocationInformationExtIEsEntryPresentUserLocationInformationTNGF
	UserLocationInformationExtIEsEntryPresentUserLocationInformationTWIF
	UserLocationInformationExtIEsEntryPresentUserLocationInformationWAGF
	UserLocationInformationExtIEsEntryPresentUserLocationInformationN3IWFWithoutPortNumber
	UserLocationInformationExtIEsEntryPresentUnknown
)

type UserLocationInformationExtIEsEntry struct {
	ID                                            ProtocolIEID
	Criticality                                   Criticality
	Present                                       UserLocationInformationExtIEsEntryPresent
	UserLocationInformationTNGF                   UserLocationInformationTNGF                   `ngap:"id:244,criticality:ignore,presence:mandatory"`
	UserLocationInformationTWIF                   UserLocationInformationTWIF                   `ngap:"id:248,criticality:ignore,presence:mandatory"`
	UserLocationInformationWAGF                   UserLocationInformationWAGF                   `ngap:"id:243,criticality:ignore,presence:mandatory"`
	UserLocationInformationN3IWFWithoutPortNumber UserLocationInformationN3IWFWithoutPortNumber `ngap:"id:439,criticality:ignore,presence:mandatory"`
	UnknownValue                                  []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformation-ExtIEsEntry", v)
}

// DecodeUserLocationInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationExtIEsEntry(data []byte) (UserLocationInformationExtIEsEntry, int, error) {
	var v UserLocationInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
