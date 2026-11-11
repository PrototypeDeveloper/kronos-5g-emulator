// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WarningMessageContents represents ASN.1 WarningMessageContents.
type WarningMessageContents struct {
	Value []byte `aper:"sizeLB:1,sizeUB:9600"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WarningMessageContents) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WarningMessageContents", v)
}

// DecodeWarningMessageContents decodes a standalone APER value and reports the consumed octets.
func DecodeWarningMessageContents(data []byte) (WarningMessageContents, int, error) {
	var v WarningMessageContents
	err := ngapcodec.DecodeParameter("WarningMessageContents", data, &v)
	if ngapcodec.IsFatal(err) {
		return WarningMessageContents{}, 0, err
	}
	return v, len(data), err
}
