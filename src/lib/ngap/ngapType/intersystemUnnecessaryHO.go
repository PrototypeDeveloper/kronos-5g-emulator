// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// IntersystemUnnecessaryHO represents ASN.1 IntersystemUnnecessaryHO.
type IntersystemUnnecessaryHO struct {
	SourcecellID      NGRANCGI
	TargetcellID      EUTRACGI
	EarlyIRATHO       IntersystemUnnecessaryHOEarlyIRATHO `aper:"valueExt,valueLB:0,valueUB:1"`
	CandidateCellList CandidateCellList
	HasIEExtensions   bool
	IEExtensions      []IntersystemUnnecessaryHOExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v IntersystemUnnecessaryHO) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("IntersystemUnnecessaryHO", v)
}

// DecodeIntersystemUnnecessaryHO decodes a standalone APER value and reports the consumed octets.
func DecodeIntersystemUnnecessaryHO(data []byte) (IntersystemUnnecessaryHO, int, error) {
	var v IntersystemUnnecessaryHO
	err := ngapcodec.DecodeParameter("IntersystemUnnecessaryHO", data, &v)
	if ngapcodec.IsFatal(err) {
		return IntersystemUnnecessaryHO{}, 0, err
	}
	return v, len(data), err
}
