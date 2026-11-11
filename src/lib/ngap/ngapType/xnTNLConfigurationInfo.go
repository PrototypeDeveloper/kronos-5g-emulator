// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// XnTNLConfigurationInfo represents ASN.1 XnTNLConfigurationInfo.
type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses            XnTLAs
	HasXnExtendedTransportLayerAddresses bool
	XnExtendedTransportLayerAddresses    XnExtTLAs
	HasIEExtensions                      bool
	IEExtensions                         []XnTNLConfigurationInfoExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v XnTNLConfigurationInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("XnTNLConfigurationInfo", v)
}

// DecodeXnTNLConfigurationInfo decodes a standalone APER value and reports the consumed octets.
func DecodeXnTNLConfigurationInfo(data []byte) (XnTNLConfigurationInfo, int, error) {
	var v XnTNLConfigurationInfo
	err := ngapcodec.DecodeParameter("XnTNLConfigurationInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return XnTNLConfigurationInfo{}, 0, err
	}
	return v, len(data), err
}
