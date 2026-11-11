// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RoutingID represents ASN.1 RoutingID.
type RoutingID struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RoutingID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("RoutingID", v) }

// DecodeRoutingID decodes a standalone APER value and reports the consumed octets.
func DecodeRoutingID(data []byte) (RoutingID, int, error) {
	var v RoutingID
	err := ngapcodec.DecodeParameter("RoutingID", data, &v)
	if ngapcodec.IsFatal(err) {
		return RoutingID{}, 0, err
	}
	return v, len(data), err
}
