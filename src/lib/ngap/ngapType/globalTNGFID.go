// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalTNGFID represents ASN.1 GlobalTNGF-ID.
type GlobalTNGFID struct {
	PLMNIdentity      PLMNIdentity
	TNGFID            TNGFID
	HasIEExtensions   bool
	IEExtensions      []GlobalTNGFIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalTNGFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GlobalTNGF-ID", v) }

// DecodeGlobalTNGFID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalTNGFID(data []byte) (GlobalTNGFID, int, error) {
	var v GlobalTNGFID
	err := ngapcodec.DecodeParameter("GlobalTNGF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalTNGFID{}, 0, err
	}
	return v, len(data), err
}
