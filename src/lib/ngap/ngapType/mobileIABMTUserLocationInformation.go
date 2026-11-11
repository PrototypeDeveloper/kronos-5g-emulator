// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MobileIABMTUserLocationInformation represents ASN.1 MobileIAB-MTUserLocationInformation.
type MobileIABMTUserLocationInformation struct {
	NRCGI             NRCGI
	TAI               TAI
	HasIEExtensions   bool
	IEExtensions      []MobileIABMTUserLocationInformationExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MobileIABMTUserLocationInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MobileIAB-MTUserLocationInformation", v)
}

// DecodeMobileIABMTUserLocationInformation decodes a standalone APER value and reports the consumed octets.
func DecodeMobileIABMTUserLocationInformation(data []byte) (MobileIABMTUserLocationInformation, int, error) {
	var v MobileIABMTUserLocationInformation
	err := ngapcodec.DecodeParameter("MobileIAB-MTUserLocationInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return MobileIABMTUserLocationInformation{}, 0, err
	}
	return v, len(data), err
}
