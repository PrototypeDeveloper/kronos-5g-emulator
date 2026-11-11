// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionToReleaseItem represents ASN.1 MBSSessionToReleaseItem.
type MBSSessionToReleaseItem struct {
	MBSSessionID      MBSSessionID
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []MBSSessionToReleaseItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionToReleaseItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionToReleaseItem", v)
}

// DecodeMBSSessionToReleaseItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionToReleaseItem(data []byte) (MBSSessionToReleaseItem, int, error) {
	var v MBSSessionToReleaseItem
	err := ngapcodec.DecodeParameter("MBSSessionToReleaseItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionToReleaseItem{}, 0, err
	}
	return v, len(data), err
}
