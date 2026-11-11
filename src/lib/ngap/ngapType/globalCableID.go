// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalCableID represents ASN.1 GlobalCable-ID.
type GlobalCableID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalCableID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalCable-ID", v)
}

// DecodeGlobalCableID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalCableID(data []byte) (GlobalCableID, int, error) {
	var v GlobalCableID
	err := ngapcodec.DecodeParameter("GlobalCable-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalCableID{}, 0, err
	}
	return v, len(data), err
}
