// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// TypeOfError represents ASN.1 TypeOfError.
const (
	TypeOfErrorPresentNotUnderstood Enumerated = 0
	TypeOfErrorPresentMissing       Enumerated = 1
)

type TypeOfError struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v TypeOfError) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("TypeOfError", v) }

// DecodeTypeOfError decodes a standalone APER value and reports the consumed octets.
func DecodeTypeOfError(data []byte) (TypeOfError, int, error) {
	var v TypeOfError
	err := ngapcodec.DecodeParameter("TypeOfError", data, &v)
	if ngapcodec.IsFatal(err) {
		return TypeOfError{}, 0, err
	}
	return v, len(data), err
}
