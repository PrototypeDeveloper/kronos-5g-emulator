// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ResetAll represents ASN.1 ResetAll.
const (
	ResetAllPresentResetAll Enumerated = 0
)

type ResetAll struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ResetAll) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("ResetAll", v) }

// DecodeResetAll decodes a standalone APER value and reports the consumed octets.
func DecodeResetAll(data []byte) (ResetAll, int, error) {
	var v ResetAll
	err := ngapcodec.DecodeParameter("ResetAll", data, &v)
	if ngapcodec.IsFatal(err) {
		return ResetAll{}, 0, err
	}
	return v, len(data), err
}
