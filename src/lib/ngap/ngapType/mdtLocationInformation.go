// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTLocationInformation represents ASN.1 MDT-Location-Information.
type MDTLocationInformation struct {
	Value BitString `aper:"sizeLB:8,sizeUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTLocationInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Location-Information", v)
}

// DecodeMDTLocationInformation decodes a standalone APER value and reports the consumed octets.
func DecodeMDTLocationInformation(data []byte) (MDTLocationInformation, int, error) {
	var v MDTLocationInformation
	err := ngapcodec.DecodeParameter("MDT-Location-Information", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTLocationInformation{}, 0, err
	}
	return v, len(data), err
}
