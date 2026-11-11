// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AUN3DeviceAccessInfo represents ASN.1 AUN3DeviceAccessInfo.
type AUN3DeviceAccessInfo struct {
	AUN3DeviceAccess                     AUN3DeviceAccessInfoAUN3DeviceAccess                     `aper:"valueExt,valueLB:0,valueUB:0"`
	ExistingNGConnectionwithSameLocation AUN3DeviceAccessInfoExistingNGConnectionwithSameLocation `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions                      bool
	IEExtensions                         []AUN3DeviceAccessInfoExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AUN3DeviceAccessInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AUN3DeviceAccessInfo", v)
}

// DecodeAUN3DeviceAccessInfo decodes a standalone APER value and reports the consumed octets.
func DecodeAUN3DeviceAccessInfo(data []byte) (AUN3DeviceAccessInfo, int, error) {
	var v AUN3DeviceAccessInfo
	err := ngapcodec.DecodeParameter("AUN3DeviceAccessInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return AUN3DeviceAccessInfo{}, 0, err
	}
	return v, len(data), err
}
