// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PacketLossRate represents ASN.1 PacketLossRate.
type PacketLossRate struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:1000"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PacketLossRate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PacketLossRate", v)
}

// DecodePacketLossRate decodes a standalone APER value and reports the consumed octets.
func DecodePacketLossRate(data []byte) (PacketLossRate, int, error) {
	var v PacketLossRate
	err := ngapcodec.DecodeParameter("PacketLossRate", data, &v)
	if ngapcodec.IsFatal(err) {
		return PacketLossRate{}, 0, err
	}
	return v, len(data), err
}
