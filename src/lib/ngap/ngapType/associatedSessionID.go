// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AssociatedSessionID represents ASN.1 AssociatedSessionID.
type AssociatedSessionID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AssociatedSessionID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AssociatedSessionID", v)
}

// DecodeAssociatedSessionID decodes a standalone APER value and reports the consumed octets.
func DecodeAssociatedSessionID(data []byte) (AssociatedSessionID, int, error) {
	var v AssociatedSessionID
	err := ngapcodec.DecodeParameter("AssociatedSessionID", data, &v)
	if ngapcodec.IsFatal(err) {
		return AssociatedSessionID{}, 0, err
	}
	return v, len(data), err
}
