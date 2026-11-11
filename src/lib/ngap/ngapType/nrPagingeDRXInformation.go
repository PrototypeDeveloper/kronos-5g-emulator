// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NRPagingeDRXInformation represents ASN.1 NR-PagingeDRXInformation.
type NRPagingeDRXInformation struct {
	NRPagingEDRXCycle     NRPagingEDRXCycle
	HasNRPagingTimeWindow bool
	NRPagingTimeWindow    NRPagingTimeWindow
	HasIEExtensions       bool
	IEExtensions          []NRPagingeDRXInformationExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NRPagingeDRXInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NR-PagingeDRXInformation", v)
}

// DecodeNRPagingeDRXInformation decodes a standalone APER value and reports the consumed octets.
func DecodeNRPagingeDRXInformation(data []byte) (NRPagingeDRXInformation, int, error) {
	var v NRPagingeDRXInformation
	err := ngapcodec.DecodeParameter("NR-PagingeDRXInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return NRPagingeDRXInformation{}, 0, err
	}
	return v, len(data), err
}
