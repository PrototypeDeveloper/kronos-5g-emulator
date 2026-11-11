// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UERetentionInformation represents ASN.1 UERetentionInformation.
const (
	UERetentionInformationPresentUesRetained Enumerated = 0
)

type UERetentionInformation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UERetentionInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UERetentionInformation", v)
}

// DecodeUERetentionInformation decodes a standalone APER value and reports the consumed octets.
func DecodeUERetentionInformation(data []byte) (UERetentionInformation, int, error) {
	var v UERetentionInformation
	err := ngapcodec.DecodeParameter("UERetentionInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return UERetentionInformation{}, 0, err
	}
	return v, len(data), err
}
