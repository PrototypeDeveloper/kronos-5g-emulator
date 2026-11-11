// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type NRPagingeDRXInformationExtIEsEntryPresent uint32

const (
	NRPagingeDRXInformationExtIEsEntryPresentNothing NRPagingeDRXInformationExtIEsEntryPresent = iota
	NRPagingeDRXInformationExtIEsEntryPresentUnknown
)

type NRPagingeDRXInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      NRPagingeDRXInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingeDRXInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-PagingeDRXInformation-ExtIEsEntry", v)
}

// DecodeNRPagingeDRXInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingeDRXInformationExtIEsEntry(data []byte) (NRPagingeDRXInformationExtIEsEntry, int, error) {
	var v NRPagingeDRXInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("NR-PagingeDRXInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingeDRXInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
