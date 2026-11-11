// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NewSecurityContextInd represents ASN.1 NewSecurityContextInd.
const (
	NewSecurityContextIndPresentTrue Enumerated = 0
)

type NewSecurityContextInd struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NewSecurityContextInd) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NewSecurityContextInd", v)
}

// DecodeNewSecurityContextInd decodes a standalone APER value and reports the consumed octets.
func DecodeNewSecurityContextInd(data []byte) (NewSecurityContextInd, int, error) {
	var v NewSecurityContextInd
	err := ngapcodec.DecodeParameter("NewSecurityContextInd", data, &v)
	if ngapcodec.IsFatal(err) {
		return NewSecurityContextInd{}, 0, err
	}
	return v, len(data), err
}
