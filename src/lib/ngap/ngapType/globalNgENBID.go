// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalNgENBID represents ASN.1 GlobalNgENB-ID.
type GlobalNgENBID struct {
	PLMNIdentity      PLMNIdentity
	NgENBID           NgENBID
	HasIEExtensions   bool
	IEExtensions      []GlobalNgENBIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalNgENBID) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GlobalNgENB-ID", v)
}

// DecodeGlobalNgENBID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalNgENBID(data []byte) (GlobalNgENBID, int, error) {
	var v GlobalNgENBID
	err := ngapcodec.DecodeParameter("GlobalNgENB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalNgENBID{}, 0, err
	}
	return v, len(data), err
}
