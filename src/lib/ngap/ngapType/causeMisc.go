// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CauseMisc represents ASN.1 CauseMisc.
const (
	CauseMiscPresentControlProcessingOverload             Enumerated = 0
	CauseMiscPresentNotEnoughUserPlaneProcessingResources Enumerated = 1
	CauseMiscPresentHardwareFailure                       Enumerated = 2
	CauseMiscPresentOmIntervention                        Enumerated = 3
	CauseMiscPresentUnknownPLMNOrSNPN                     Enumerated = 4
	CauseMiscPresentUnspecified                           Enumerated = 5
)

type CauseMisc struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CauseMisc) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("CauseMisc", v) }

// DecodeCauseMisc decodes a standalone APER value and reports the consumed octets.
func DecodeCauseMisc(data []byte) (CauseMisc, int, error) {
	var v CauseMisc
	err := ngapcodec.DecodeParameter("CauseMisc", data, &v)
	if ngapcodec.IsFatal(err) {
		return CauseMisc{}, 0, err
	}
	return v, len(data), err
}
