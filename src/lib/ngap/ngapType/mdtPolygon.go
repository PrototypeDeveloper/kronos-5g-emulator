// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTPolygon represents ASN.1 MDT-polygon.
type MDTPolygon struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTPolygon) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MDT-polygon", v) }

// DecodeMDTPolygon decodes a standalone APER value and reports the consumed octets.
func DecodeMDTPolygon(data []byte) (MDTPolygon, int, error) {
	var v MDTPolygon
	err := ngapcodec.DecodeParameter("MDT-polygon", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTPolygon{}, 0, err
	}
	return v, len(data), err
}
