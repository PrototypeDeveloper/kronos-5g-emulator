// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// PacketErrorRate represents ASN.1 PacketErrorRate.
type PacketErrorRate struct {
	PERScalar         int64 `aper:"valueExt,valueLB:0,valueUB:9"`
	PERExponent       int64 `aper:"valueExt,valueLB:0,valueUB:9"`
	HasIEExtensions   bool
	IEExtensions      []PacketErrorRateExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PacketErrorRate) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PacketErrorRate", v)
}

// DecodePacketErrorRate decodes a standalone APER value and reports the consumed octets.
func DecodePacketErrorRate(data []byte) (PacketErrorRate, int, error) {
	var v PacketErrorRate
	err := ngapcodec.DecodeParameter("PacketErrorRate", data, &v)
	if ngapcodec.IsFatal(err) {
		return PacketErrorRate{}, 0, err
	}
	return v, len(data), err
}
