// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PacketDelayBudget represents ASN.1 PacketDelayBudget.
type PacketDelayBudget struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:1023"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PacketDelayBudget) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PacketDelayBudget", v)
}

// DecodePacketDelayBudget decodes a standalone APER value and reports the consumed octets.
func DecodePacketDelayBudget(data []byte) (PacketDelayBudget, int, error) {
	var v PacketDelayBudget
	err := ngapcodec.DecodeParameter("PacketDelayBudget", data, &v)
	if ngapcodec.IsFatal(err) {
		return PacketDelayBudget{}, 0, err
	}
	return v, len(data), err
}
