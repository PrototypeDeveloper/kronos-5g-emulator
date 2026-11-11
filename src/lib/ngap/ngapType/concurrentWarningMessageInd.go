// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ConcurrentWarningMessageInd represents ASN.1 ConcurrentWarningMessageInd.
const (
	ConcurrentWarningMessageIndPresentTrue Enumerated = 0
)

type ConcurrentWarningMessageInd struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ConcurrentWarningMessageInd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ConcurrentWarningMessageInd", v)
}

// DecodeConcurrentWarningMessageInd decodes a standalone APER value and reports the consumed octets.
func DecodeConcurrentWarningMessageInd(data []byte) (ConcurrentWarningMessageInd, int, error) {
	var v ConcurrentWarningMessageInd
	err := ngapcodec.DecodeParameter("ConcurrentWarningMessageInd", data, &v)
	if ngapcodec.IsFatal(err) {
		return ConcurrentWarningMessageInd{}, 0, err
	}
	return v, len(data), err
}
