// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EUTRAPagingeDRXInformationExtIEsEntryPresent uint32

const (
	EUTRAPagingeDRXInformationExtIEsEntryPresentNothing EUTRAPagingeDRXInformationExtIEsEntryPresent = iota
	EUTRAPagingeDRXInformationExtIEsEntryPresentUnknown
)

type EUTRAPagingeDRXInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EUTRAPagingeDRXInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRAPagingeDRXInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRA-PagingeDRXInformation-ExtIEsEntry", v)
}

// DecodeEUTRAPagingeDRXInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRAPagingeDRXInformationExtIEsEntry(data []byte) (EUTRAPagingeDRXInformationExtIEsEntry, int, error) {
	var v EUTRAPagingeDRXInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("EUTRA-PagingeDRXInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRAPagingeDRXInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
