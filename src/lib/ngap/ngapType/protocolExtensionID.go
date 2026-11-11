// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ProtocolExtensionID represents ASN.1 ProtocolExtensionID.
type ProtocolExtensionID struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ProtocolExtensionID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ProtocolExtensionID", v)
}

// DecodeProtocolExtensionID decodes a standalone APER value and reports the consumed octets.
func DecodeProtocolExtensionID(data []byte) (ProtocolExtensionID, int, error) {
	var v ProtocolExtensionID
	err := ngapcodec.DecodeParameter("ProtocolExtensionID", data, &v)
	if ngapcodec.IsFatal(err) {
		return ProtocolExtensionID{}, 0, err
	}
	return v, len(data), err
}
