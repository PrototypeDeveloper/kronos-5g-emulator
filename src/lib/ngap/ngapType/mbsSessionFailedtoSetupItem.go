// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionFailedtoSetupItem represents ASN.1 MBSSessionFailedtoSetupItem.
type MBSSessionFailedtoSetupItem struct {
	MBSSessionID        MBSSessionID
	HasMBSAreaSessionID bool
	MBSAreaSessionID    MBSAreaSessionID
	Cause               Cause
	HasIEExtensions     bool
	IEExtensions        []MBSSessionFailedtoSetupItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionFailedtoSetupItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionFailedtoSetupItem", v)
}

// DecodeMBSSessionFailedtoSetupItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionFailedtoSetupItem(data []byte) (MBSSessionFailedtoSetupItem, int, error) {
	var v MBSSessionFailedtoSetupItem
	err := ngapcodec.DecodeParameter("MBSSessionFailedtoSetupItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionFailedtoSetupItem{}, 0, err
	}
	return v, len(data), err
}
