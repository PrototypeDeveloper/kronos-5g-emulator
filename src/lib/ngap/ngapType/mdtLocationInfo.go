// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MDTLocationInfo represents ASN.1 MDT-Location-Info.
type MDTLocationInfo struct {
	MDTLocationInformation MDTLocationInformation
	HasIEExtensions        bool
	IEExtensions           []MDTLocationInfoExtIEsEntry
	UnknownExtensions      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MDTLocationInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MDT-Location-Info", v)
}

// DecodeMDTLocationInfo decodes a standalone APER value and reports the consumed octets.
func DecodeMDTLocationInfo(data []byte) (MDTLocationInfo, int, error) {
	var v MDTLocationInfo
	err := ngapcodec.DecodeParameter("MDT-Location-Info", data, &v)
	if ngapcodec.IsFatal(err) {
		return MDTLocationInfo{}, 0, err
	}
	return v, len(data), err
}
