// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRAPagingeDRXInformation represents ASN.1 EUTRA-PagingeDRXInformation.
type EUTRAPagingeDRXInformation struct {
	EUTRAPagingEDRXCycle     EUTRAPagingEDRXCycle
	HasEUTRAPagingTimeWindow bool
	EUTRAPagingTimeWindow    EUTRAPagingTimeWindow
	HasIEExtensions          bool
	IEExtensions             []EUTRAPagingeDRXInformationExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRAPagingeDRXInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRA-PagingeDRXInformation", v)
}

// DecodeEUTRAPagingeDRXInformation decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRAPagingeDRXInformation(data []byte) (EUTRAPagingeDRXInformation, int, error) {
	var v EUTRAPagingeDRXInformation
	err := ngapcodec.DecodeParameter("EUTRA-PagingeDRXInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRAPagingeDRXInformation{}, 0, err
	}
	return v, len(data), err
}
