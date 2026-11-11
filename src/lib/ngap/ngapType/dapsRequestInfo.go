// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DAPSRequestInfo represents ASN.1 DAPSRequestInfo.
type DAPSRequestInfo struct {
	DAPSIndicator     DAPSRequestInfoDAPSIndicator `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions   bool
	IEExtensions      []DAPSRequestInfoExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSRequestInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSRequestInfo", v)
}

// DecodeDAPSRequestInfo decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSRequestInfo(data []byte) (DAPSRequestInfo, int, error) {
	var v DAPSRequestInfo
	err := ngapcodec.DecodeParameter("DAPSRequestInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSRequestInfo{}, 0, err
	}
	return v, len(data), err
}
