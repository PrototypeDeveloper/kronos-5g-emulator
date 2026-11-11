// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ProtocolIEID represents ASN.1 ProtocolIE-ID.
type ProtocolIEID struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ProtocolIEID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ProtocolIE-ID", v) }

// DecodeProtocolIEID decodes a standalone APER value and reports the consumed octets.
func DecodeProtocolIEID(data []byte) (ProtocolIEID, int, error) {
	var v ProtocolIEID
	err := ngapcodec.DecodeParameter("ProtocolIE-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return ProtocolIEID{}, 0, err
	}
	return v, len(data), err
}
