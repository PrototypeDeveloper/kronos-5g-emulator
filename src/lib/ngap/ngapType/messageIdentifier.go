// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MessageIdentifier represents ASN.1 MessageIdentifier.
type MessageIdentifier struct {
	Value BitString `aper:"sizeLB:16,sizeUB:16"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MessageIdentifier) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MessageIdentifier", v)
}

// DecodeMessageIdentifier decodes a standalone APER value and reports the consumed octets.
func DecodeMessageIdentifier(data []byte) (MessageIdentifier, int, error) {
	var v MessageIdentifier
	err := ngapcodec.DecodeParameter("MessageIdentifier", data, &v)
	if ngapcodec.IsFatal(err) {
		return MessageIdentifier{}, 0, err
	}
	return v, len(data), err
}
