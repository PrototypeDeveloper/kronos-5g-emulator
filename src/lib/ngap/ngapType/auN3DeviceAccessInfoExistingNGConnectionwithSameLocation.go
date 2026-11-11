// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation represents ASN.1 AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation.
const (
	AUN3DeviceAccessInfoExistingNGConnectionwithSameLocationPresentTrue  Enumerated = 0
	AUN3DeviceAccessInfoExistingNGConnectionwithSameLocationPresentFalse Enumerated = 1
)

type AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}

// Encode encodes this parameter as a standalone APER value.
func (v AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation", v)
}

// DecodeAUN3DeviceAccessInfoExistingNGConnectionwithSameLocation decodes a standalone APER value and reports consumed octets.
func DecodeAUN3DeviceAccessInfoExistingNGConnectionwithSameLocation(data []byte) (AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation, int, error) {
	var v AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation
	err := ngapcodec.DecodeParameter("inline:AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation", data, &v)
	if ngapcodec.IsFatal(err) {
		return AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation{}, 0, err
	}
	return v, len(data), err
}
