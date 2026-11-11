// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NBIoTPagingDRX represents ASN.1 NB-IoT-PagingDRX.
const (
	NBIoTPagingDRXPresentRf32   Enumerated = 0
	NBIoTPagingDRXPresentRf64   Enumerated = 1
	NBIoTPagingDRXPresentRf128  Enumerated = 2
	NBIoTPagingDRXPresentRf256  Enumerated = 3
	NBIoTPagingDRXPresentRf512  Enumerated = 4
	NBIoTPagingDRXPresentRf1024 Enumerated = 5
)

type NBIoTPagingDRX struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTPagingDRX) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-PagingDRX", v)
}

// DecodeNBIoTPagingDRX decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTPagingDRX(data []byte) (NBIoTPagingDRX, int, error) {
	var v NBIoTPagingDRX
	err := ngapcodec.DecodeParameter("NB-IoT-PagingDRX", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTPagingDRX{}, 0, err
	}
	return v, len(data), err
}
