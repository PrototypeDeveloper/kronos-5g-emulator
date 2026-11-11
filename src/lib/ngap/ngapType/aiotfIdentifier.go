// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTFIdentifier represents ASN.1 AIOTFIdentifier.
type AIOTFIdentifier struct {
	Value []byte `aper:"sizeExt,sizeLB:16,sizeUB:36"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTFIdentifier) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTFIdentifier", v)
}

// DecodeAIOTFIdentifier decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTFIdentifier(data []byte) (AIOTFIdentifier, int, error) {
	var v AIOTFIdentifier
	err := ngapcodec.DecodeParameter("AIOTFIdentifier", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTFIdentifier{}, 0, err
	}
	return v, len(data), err
}
