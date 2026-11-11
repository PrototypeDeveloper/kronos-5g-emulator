// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalN3IWFID represents ASN.1 GlobalN3IWF-ID.
type GlobalN3IWFID struct {
	PLMNIdentity      PLMNIdentity
	N3IWFID           N3IWFID
	HasIEExtensions   bool
	IEExtensions      []GlobalN3IWFIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalN3IWFID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalN3IWF-ID", v)
}

// DecodeGlobalN3IWFID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalN3IWFID(data []byte) (GlobalN3IWFID, int, error) {
	var v GlobalN3IWFID
	err := ngapcodec.DecodeParameter("GlobalN3IWF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalN3IWFID{}, 0, err
	}
	return v, len(data), err
}
