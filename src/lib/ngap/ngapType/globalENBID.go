// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalENBID represents ASN.1 GlobalENB-ID.
type GlobalENBID struct {
	PLMNidentity      PLMNIdentity
	ENBID             ENBID
	HasIEExtensions   bool
	IEExtensions      []GlobalENBIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalENBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GlobalENB-ID", v) }

// DecodeGlobalENBID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalENBID(data []byte) (GlobalENBID, int, error) {
	var v GlobalENBID
	err := ngapcodec.DecodeParameter("GlobalENB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalENBID{}, 0, err
	}
	return v, len(data), err
}
