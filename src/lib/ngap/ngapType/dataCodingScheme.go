// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DataCodingScheme represents ASN.1 DataCodingScheme.
type DataCodingScheme struct {
	Value BitString `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DataCodingScheme) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DataCodingScheme", v)
}

// DecodeDataCodingScheme decodes a standalone APER value and reports the consumed octets.
func DecodeDataCodingScheme(data []byte) (DataCodingScheme, int, error) {
	var v DataCodingScheme
	err := ngapcodec.DecodeParameter("DataCodingScheme", data, &v)
	if ngapcodec.IsFatal(err) {
		return DataCodingScheme{}, 0, err
	}
	return v, len(data), err
}
