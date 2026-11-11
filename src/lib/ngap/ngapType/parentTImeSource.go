// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ParentTImeSource represents ASN.1 ParentTImeSource.
const (
	ParentTImeSourcePresentSynce            Enumerated = 0
	ParentTImeSourcePresentPtp              Enumerated = 1
	ParentTImeSourcePresentGnss             Enumerated = 2
	ParentTImeSourcePresentAtomicclock      Enumerated = 3
	ParentTImeSourcePresentTerrestrialradio Enumerated = 4
	ParentTImeSourcePresentSerialtimecode   Enumerated = 5
	ParentTImeSourcePresentNtp              Enumerated = 6
	ParentTImeSourcePresentHandset          Enumerated = 7
	ParentTImeSourcePresentOther            Enumerated = 8
)

type ParentTImeSource struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:8"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ParentTImeSource) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ParentTImeSource", v)
}

// DecodeParentTImeSource decodes a standalone APER value and reports the consumed octets.
func DecodeParentTImeSource(data []byte) (ParentTImeSource, int, error) {
	var v ParentTImeSource
	err := ngapcodec.DecodeParameter("ParentTImeSource", data, &v)
	if ngapcodec.IsFatal(err) {
		return ParentTImeSource{}, 0, err
	}
	return v, len(data), err
}
