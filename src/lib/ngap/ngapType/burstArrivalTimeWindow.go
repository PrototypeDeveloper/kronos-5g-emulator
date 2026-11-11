// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// BurstArrivalTimeWindow represents ASN.1 BurstArrivalTimeWindow.
type BurstArrivalTimeWindow struct {
	BurstArrivalTimeWindowStart int64 `aper:"valueExt,valueLB:0,valueUB:640000"`
	BurstArrivalTimeWindowEnd   int64 `aper:"valueExt,valueLB:0,valueUB:640000"`
	HasIEExtension              bool
	IEExtension                 []BurstArrivalTimeWindowExtIEsEntry
	UnknownExtensions           []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v BurstArrivalTimeWindow) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("BurstArrivalTimeWindow", v)
}

// DecodeBurstArrivalTimeWindow decodes a standalone APER value and reports the consumed octets.
func DecodeBurstArrivalTimeWindow(data []byte) (BurstArrivalTimeWindow, int, error) {
	var v BurstArrivalTimeWindow
	err := ngapcodec.DecodeParameter("BurstArrivalTimeWindow", data, &v)
	if ngapcodec.IsFatal(err) {
		return BurstArrivalTimeWindow{}, 0, err
	}
	return v, len(data), err
}
