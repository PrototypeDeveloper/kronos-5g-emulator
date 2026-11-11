// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// RequestedTNLInfo represents ASN.1 RequestedTNLInfo.
type RequestedTNLInfo struct {
	PLMNIdentity      PLMNIdentity
	HasIEExtensions   bool
	IEExtensions      []RequestedTNLInfoExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v RequestedTNLInfo) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("RequestedTNLInfo", v)
}

// DecodeRequestedTNLInfo decodes a standalone APER value and reports the consumed octets.
func DecodeRequestedTNLInfo(data []byte) (RequestedTNLInfo, int, error) {
	var v RequestedTNLInfo
	err := ngapcodec.DecodeParameter("RequestedTNLInfo", data, &v)
	if ngapcodec.IsFatal(err) {
		return RequestedTNLInfo{}, 0, err
	}
	return v, len(data), err
}
