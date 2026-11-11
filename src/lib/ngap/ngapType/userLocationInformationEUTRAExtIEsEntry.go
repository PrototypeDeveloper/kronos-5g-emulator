// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type UserLocationInformationEUTRAExtIEsEntryPresent uint32

const (
	UserLocationInformationEUTRAExtIEsEntryPresentNothing UserLocationInformationEUTRAExtIEsEntryPresent = iota
	UserLocationInformationEUTRAExtIEsEntryPresentPSCellInformation
	UserLocationInformationEUTRAExtIEsEntryPresentUnknown
)

type UserLocationInformationEUTRAExtIEsEntry struct {
	ID                ProtocolExtensionID
	Criticality       Criticality
	Present           UserLocationInformationEUTRAExtIEsEntryPresent
	PSCellInformation NGRANCGI `ngap:"id:149,criticality:ignore,presence:optional"`
	UnknownValue      []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UserLocationInformationEUTRAExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UserLocationInformationEUTRA-ExtIEsEntry", v)
}

// DecodeUserLocationInformationEUTRAExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeUserLocationInformationEUTRAExtIEsEntry(data []byte) (UserLocationInformationEUTRAExtIEsEntry, int, error) {
	var v UserLocationInformationEUTRAExtIEsEntry
	err := ngapcodec.DecodeParameter("UserLocationInformationEUTRA-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return UserLocationInformationEUTRAExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
