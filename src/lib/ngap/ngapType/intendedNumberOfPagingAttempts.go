// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntendedNumberOfPagingAttempts represents ASN.1 IntendedNumberOfPagingAttempts.
type IntendedNumberOfPagingAttempts struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntendedNumberOfPagingAttempts) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntendedNumberOfPagingAttempts", v)
}

// DecodeIntendedNumberOfPagingAttempts decodes a standalone APER value and reports the consumed octets.
func DecodeIntendedNumberOfPagingAttempts(data []byte) (IntendedNumberOfPagingAttempts, int, error) {
	var v IntendedNumberOfPagingAttempts
	err := ngapcodec.DecodeParameter("IntendedNumberOfPagingAttempts", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntendedNumberOfPagingAttempts{}, 0, err
	}
	return v, len(data), err
}
