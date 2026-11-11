// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GlobalWAGFID represents ASN.1 GlobalW-AGF-ID.
type GlobalWAGFID struct {
	PLMNIdentity      PLMNIdentity
	WAGFID            WAGFID
	HasIEExtensions   bool
	IEExtensions      []GlobalWAGFIDExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GlobalWAGFID) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GlobalW-AGF-ID", v) }

// DecodeGlobalWAGFID decodes a standalone APER value and reports the consumed octets.
func DecodeGlobalWAGFID(data []byte) (GlobalWAGFID, int, error) {
	var v GlobalWAGFID
	err := ngapcodec.DecodeParameter("GlobalW-AGF-ID", data, &v)
	if ngapcodec.IsFatal(err) {
		return GlobalWAGFID{}, 0, err
	}
	return v, len(data), err
}
