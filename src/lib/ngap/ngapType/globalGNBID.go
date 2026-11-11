// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalGNBID represents ASN.1 GlobalGNB-ID.
type GlobalGNBID struct {
	PLMNIdentity      PLMNIdentity
	GNBID             GNBID
	HasIEExtensions   bool
	IEExtensions      []GlobalGNBIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalGNBID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GlobalGNB-ID", v) }

// DecodeGlobalGNBID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalGNBID(data []byte) (GlobalGNBID, int, error) {
	var v GlobalGNBID
	err := ngapcodec.DecodeParameter("GlobalGNB-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalGNBID{}, 0, err
	}
	return v, len(data), err
}
