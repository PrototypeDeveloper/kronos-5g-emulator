// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NGRANNumberOfActiveUEs represents ASN.1 NGRAN-NumberOfActiveUEs.
type NGRANNumberOfActiveUEs struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:16777215"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NGRANNumberOfActiveUEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NGRAN-NumberOfActiveUEs", v)
}

// DecodeNGRANNumberOfActiveUEs decodes a standalone APER value and reports the consumed octets.
func DecodeNGRANNumberOfActiveUEs(data []byte) (NGRANNumberOfActiveUEs, int, error) {
	var v NGRANNumberOfActiveUEs
	err := ngapcodec.DecodeParameter("NGRAN-NumberOfActiveUEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return NGRANNumberOfActiveUEs{}, 0, err
	}
	return v, len(data), err
}
