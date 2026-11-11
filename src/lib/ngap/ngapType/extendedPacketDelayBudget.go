// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExtendedPacketDelayBudget represents ASN.1 ExtendedPacketDelayBudget.
type ExtendedPacketDelayBudget struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:65535"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExtendedPacketDelayBudget) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExtendedPacketDelayBudget", v)
}

// DecodeExtendedPacketDelayBudget decodes a standalone APER value and reports the consumed octets.
func DecodeExtendedPacketDelayBudget(data []byte) (ExtendedPacketDelayBudget, int, error) {
	var v ExtendedPacketDelayBudget
	err := ngapcodec.DecodeParameter("ExtendedPacketDelayBudget", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExtendedPacketDelayBudget{}, 0, err
	}
	return v, len(data), err
}
