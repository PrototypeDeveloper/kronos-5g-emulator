// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AUN3DeviceAccessInfoAUN3DeviceAccess represents ASN.1 AUN3DeviceAccessInfoAUN3DeviceAccess.
const (
	AUN3DeviceAccessInfoAUN3DeviceAccessPresentTrue Enumerated = 0
)

type AUN3DeviceAccessInfoAUN3DeviceAccess struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v AUN3DeviceAccessInfoAUN3DeviceAccess) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:AUN3DeviceAccessInfoAUN3DeviceAccess", v)
}

// DecodeAUN3DeviceAccessInfoAUN3DeviceAccess decodes a standalone APER value and reports consumed octets.
func DecodeAUN3DeviceAccessInfoAUN3DeviceAccess(data []byte) (AUN3DeviceAccessInfoAUN3DeviceAccess, int, error) {
	var v AUN3DeviceAccessInfoAUN3DeviceAccess
	err := ngapcodec.DecodeParameter("inline:AUN3DeviceAccessInfoAUN3DeviceAccess", data, &v)
	if ngapcodec.IsFatal(err) {
		return AUN3DeviceAccessInfoAUN3DeviceAccess{}, 0, err
	}
	return v, len(data), err
}
