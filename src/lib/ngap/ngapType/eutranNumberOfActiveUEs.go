// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EUTRANNumberOfActiveUEs represents ASN.1 EUTRAN-NumberOfActiveUEs.
type EUTRANNumberOfActiveUEs struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:16777215"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EUTRANNumberOfActiveUEs) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EUTRAN-NumberOfActiveUEs", v)
}

// DecodeEUTRANNumberOfActiveUEs decodes a standalone APER value and reports the consumed octets.
func DecodeEUTRANNumberOfActiveUEs(data []byte) (EUTRANNumberOfActiveUEs, int, error) {
	var v EUTRANNumberOfActiveUEs
	err := ngapcodec.DecodeParameter("EUTRAN-NumberOfActiveUEs", data, &v)
	if ngapcodec.IsFatal(err) {
		return EUTRANNumberOfActiveUEs{}, 0, err
	}
	return v, len(data), err
}
