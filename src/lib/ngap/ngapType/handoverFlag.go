// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverFlag represents ASN.1 HandoverFlag.
const (
	HandoverFlagPresentHandoverPreparation Enumerated = 0
)

type HandoverFlag struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverFlag) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("HandoverFlag", v) }

// DecodeHandoverFlag decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverFlag(data []byte) (HandoverFlag, int, error) {
	var v HandoverFlag
	err := ngapcodec.DecodeParameter("HandoverFlag", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverFlag{}, 0, err
	}
	return v, len(data), err
}
