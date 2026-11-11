// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MaximumDataBurstVolume represents ASN.1 MaximumDataBurstVolume.
type MaximumDataBurstVolume struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:4095"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MaximumDataBurstVolume) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MaximumDataBurstVolume", v)
}

// DecodeMaximumDataBurstVolume decodes a standalone APER value and reports the consumed octets.
func DecodeMaximumDataBurstVolume(data []byte) (MaximumDataBurstVolume, int, error) {
	var v MaximumDataBurstVolume
	err := ngapcodec.DecodeParameter("MaximumDataBurstVolume", data, &v)
	if ngapcodec.IsFatal(err) {
		return MaximumDataBurstVolume{}, 0, err
	}
	return v, len(data), err
}
