// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemUnnecessaryHOEarlyIRATHO represents ASN.1 IntersystemUnnecessaryHOEarlyIRATHO.
const (
	IntersystemUnnecessaryHOEarlyIRATHOPresentTrue  Enumerated = 0
	IntersystemUnnecessaryHOEarlyIRATHOPresentFalse Enumerated = 1
)

type IntersystemUnnecessaryHOEarlyIRATHO struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v IntersystemUnnecessaryHOEarlyIRATHO) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:IntersystemUnnecessaryHOEarlyIRATHO", v)
}

// DecodeIntersystemUnnecessaryHOEarlyIRATHO decodes a standalone APER value and reports consumed octets.
func DecodeIntersystemUnnecessaryHOEarlyIRATHO(data []byte) (IntersystemUnnecessaryHOEarlyIRATHO, int, error) {
	var v IntersystemUnnecessaryHOEarlyIRATHO
	err := ngapcodec.DecodeParameter("inline:IntersystemUnnecessaryHOEarlyIRATHO", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemUnnecessaryHOEarlyIRATHO{}, 0, err
	}
	return v, len(data), err
}
