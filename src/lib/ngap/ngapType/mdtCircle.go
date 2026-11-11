// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTCircle represents ASN.1 MDT-circle.
type MDTCircle struct {
	ReferenceLocation []byte
	DistanceRadius    int64 `aper:"valueLB:1,valueUB:65535"`
	HasIEExtensions   bool
	IEExtensions      []MDTCircleExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTCircle) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("MDT-circle", v) }

// DecodeMDTCircle decodes a standalone APER value and reports the consumed octets.
func DecodeMDTCircle(data []byte) (MDTCircle, int, error) {
	var v MDTCircle
	err := ngapcodec.DecodeParameter("MDT-circle", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTCircle{}, 0, err
	}
	return v, len(data), err
}
