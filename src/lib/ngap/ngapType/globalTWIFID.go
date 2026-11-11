// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalTWIFID represents ASN.1 GlobalTWIF-ID.
type GlobalTWIFID struct {
	PLMNIdentity      PLMNIdentity
	TWIFID            TWIFID
	HasIEExtensions   bool
	IEExtensions      []GlobalTWIFIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalTWIFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GlobalTWIF-ID", v) }

// DecodeGlobalTWIFID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalTWIFID(data []byte) (GlobalTWIFID, int, error) {
	var v GlobalTWIFID
	err := ngapcodec.DecodeParameter("GlobalTWIF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalTWIFID{}, 0, err
	}
	return v, len(data), err
}
