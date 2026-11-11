// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// DAPSResponseInfo represents ASN.1 DAPSResponseInfo.
type DAPSResponseInfo struct {
	Dapsresponseindicator DAPSResponseInfoDapsresponseindicator `aper:"valueExt,valueLB:0,valueUB:1"`
	HasIEExtensions       bool
	IEExtensions          []DAPSResponseInfoExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v DAPSResponseInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("DAPSResponseInfo", v)
}

// DecodeDAPSResponseInfo decodes a standalone APER value and reports the consumed octets.
func DecodeDAPSResponseInfo(data []byte) (DAPSResponseInfo, int, error) {
	var v DAPSResponseInfo
	err := ngapcodec.DecodeParameter("DAPSResponseInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return DAPSResponseInfo{}, 0, err
	}
	return v, len(data), err
}
