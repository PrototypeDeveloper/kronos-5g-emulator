// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBStatusUL represents ASN.1 DRBStatusUL.
type DRBStatusULPresent uint32

const (
	DRBStatusULPresentNothing DRBStatusULPresent = iota
	DRBStatusULPresentDRBStatusUL12
	DRBStatusULPresentDRBStatusUL18
	DRBStatusULPresentChoiceExtensions
)

type DRBStatusUL struct {
	Present          DRBStatusULPresent
	DRBStatusUL12    DRBStatusUL12
	DRBStatusUL18    DRBStatusUL18
	ChoiceExtensions DRBStatusULExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusUL) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRBStatusUL", v) }

// DecodeDRBStatusUL decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusUL(data []byte) (DRBStatusUL, int, error) {
	var v DRBStatusUL
	err := ngapcodec.DecodeParameter("DRBStatusUL", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusUL{}, 0, err
	}
	return v, len(data), err
}
