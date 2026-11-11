// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEUPCIoTSupport represents ASN.1 UE-UP-CIoT-Support.
const (
	UEUPCIoTSupportPresentSupported Enumerated = 0
)

type UEUPCIoTSupport struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEUPCIoTSupport) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-UP-CIoT-Support", v)
}

// DecodeUEUPCIoTSupport decodes a standalone APER value and reports the consumed octets.
func DecodeUEUPCIoTSupport(data []byte) (UEUPCIoTSupport, int, error) {
	var v UEUPCIoTSupport
	err := ngapcodec.DecodeParameter("UE-UP-CIoT-Support", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEUPCIoTSupport{}, 0, err
	}
	return v, len(data), err
}
