// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// XrDeviceWith2Rx represents ASN.1 XrDeviceWith2Rx.
const (
	XrDeviceWith2RxPresentTrue Enumerated = 0
)

type XrDeviceWith2Rx struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XrDeviceWith2Rx) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("XrDeviceWith2Rx", v)
}

// DecodeXrDeviceWith2Rx decodes a standalone APER value and reports the consumed octets.
func DecodeXrDeviceWith2Rx(data []byte) (XrDeviceWith2Rx, int, error) {
	var v XrDeviceWith2Rx
	err := ngapcodec.DecodeParameter("XrDeviceWith2Rx", data, &v)
	if ngapcodec.IsFatal(err) {
		return XrDeviceWith2Rx{}, 0, err
	}
	return v, len(data), err
}
