// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DRBStatusDL represents ASN.1 DRBStatusDL.
type DRBStatusDLPresent uint32

const (
	DRBStatusDLPresentNothing DRBStatusDLPresent = iota
	DRBStatusDLPresentDRBStatusDL12
	DRBStatusDLPresentDRBStatusDL18
	DRBStatusDLPresentChoiceExtensions
)

type DRBStatusDL struct {
	Present          DRBStatusDLPresent
	DRBStatusDL12    DRBStatusDL12
	DRBStatusDL18    DRBStatusDL18
	ChoiceExtensions DRBStatusDLExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DRBStatusDL) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("DRBStatusDL", v) }

// DecodeDRBStatusDL decodes a standalone APER value and reports the consumed octets.
func DecodeDRBStatusDL(data []byte) (DRBStatusDL, int, error) {
	var v DRBStatusDL
	err := ngapcodec.DecodeParameter("DRBStatusDL", data, &v)
	if ngapcodec.IsFatal(err) {
		return DRBStatusDL{}, 0, err
	}
	return v, len(data), err
}
