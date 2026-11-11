// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// HandoverType represents ASN.1 HandoverType.
const (
	HandoverTypePresentIntra5gs      Enumerated = 0
	HandoverTypePresentFivegsToEps   Enumerated = 1
	HandoverTypePresentEpsTo5gs      Enumerated = 2
	HandoverTypePresentFivegsToUtran Enumerated = 3
)

type HandoverType struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:2"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v HandoverType) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("HandoverType", v) }

// DecodeHandoverType decodes a standalone APER value and reports the consumed octets.
func DecodeHandoverType(data []byte) (HandoverType, int, error) {
	var v HandoverType
	err := ngapcodec.DecodeParameter("HandoverType", data, &v)
	if ngapcodec.IsFatal(err) {
		return HandoverType{}, 0, err
	}
	return v, len(data), err
}
