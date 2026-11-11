// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GNBSetID represents ASN.1 GNBSetID.
type GNBSetID struct {
	Value BitString `aper:"sizeLB:22,sizeUB:22"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GNBSetID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GNBSetID", v) }

// DecodeGNBSetID decodes a standalone APER value and reports the consumed octets.
func DecodeGNBSetID(data []byte) (GNBSetID, int, error) {
	var v GNBSetID
	err := ngapcodec.DecodeParameter("GNBSetID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GNBSetID{}, 0, err
	}
	return v, len(data), err
}
