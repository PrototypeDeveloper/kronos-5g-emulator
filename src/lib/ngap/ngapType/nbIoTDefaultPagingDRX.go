// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// NBIoTDefaultPagingDRX represents ASN.1 NB-IoT-DefaultPagingDRX.
const (
	NBIoTDefaultPagingDRXPresentRf128  Enumerated = 0
	NBIoTDefaultPagingDRXPresentRf256  Enumerated = 1
	NBIoTDefaultPagingDRXPresentRf512  Enumerated = 2
	NBIoTDefaultPagingDRXPresentRf1024 Enumerated = 3
)

type NBIoTDefaultPagingDRX struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v NBIoTDefaultPagingDRX) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("NB-IoT-DefaultPagingDRX", v)
}

// DecodeNBIoTDefaultPagingDRX decodes a standalone APER value and reports the consumed octets.
func DecodeNBIoTDefaultPagingDRX(data []byte) (NBIoTDefaultPagingDRX, int, error) {
	var v NBIoTDefaultPagingDRX
	err := ngapcodec.DecodeParameter("NB-IoT-DefaultPagingDRX", data, &v)
	if ngapcodec.IsFatal(err) {
		return NBIoTDefaultPagingDRX{}, 0, err
	}
	return v, len(data), err
}
