// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CancelAllWarningMessages represents ASN.1 CancelAllWarningMessages.
const (
	CancelAllWarningMessagesPresentTrue Enumerated = 0
)

type CancelAllWarningMessages struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CancelAllWarningMessages) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CancelAllWarningMessages", v)
}

// DecodeCancelAllWarningMessages decodes a standalone APER value and reports the consumed octets.
func DecodeCancelAllWarningMessages(data []byte) (CancelAllWarningMessages, int, error) {
	var v CancelAllWarningMessages
	err := ngapcodec.DecodeParameter("CancelAllWarningMessages", data, &v)
	if ngapcodec.IsFatal(err) {
		return CancelAllWarningMessages{}, 0, err
	}
	return v, len(data), err
}
