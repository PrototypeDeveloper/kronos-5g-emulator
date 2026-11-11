// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// M1ThresholdType represents ASN.1 M1ThresholdType.
type M1ThresholdTypePresent uint32

const (
	M1ThresholdTypePresentNothing M1ThresholdTypePresent = iota
	M1ThresholdTypePresentThresholdRSRP
	M1ThresholdTypePresentThresholdRSRQ
	M1ThresholdTypePresentThresholdSINR
	M1ThresholdTypePresentChoiceExtensions
)

type M1ThresholdType struct {
	Present          M1ThresholdTypePresent
	ThresholdRSRP    ThresholdRSRP
	ThresholdRSRQ    ThresholdRSRQ
	ThresholdSINR    ThresholdSINR
	ChoiceExtensions M1ThresholdTypeExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v M1ThresholdType) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("M1ThresholdType", v)
}

// DecodeM1ThresholdType decodes a standalone APER value and reports the consumed octets.
func DecodeM1ThresholdType(data []byte) (M1ThresholdType, int, error) {
	var v M1ThresholdType
	err := ngapcodec.DecodeParameter("M1ThresholdType", data, &v)
	if ngapcodec.IsFatal(err) {
		return M1ThresholdType{}, 0, err
	}
	return v, len(data), err
}
