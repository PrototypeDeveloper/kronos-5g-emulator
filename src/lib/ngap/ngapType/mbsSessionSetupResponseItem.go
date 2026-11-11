// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupResponseItem represents ASN.1 MBSSessionSetupResponseItem.
type MBSSessionSetupResponseItem struct {
	MBSSessionID        MBSSessionID
	HasMBSAreaSessionID bool
	MBSAreaSessionID    MBSAreaSessionID
	HasIEExtensions     bool
	IEExtensions        []MBSSessionSetupResponseItemExtIEsEntry
	UnknownExtensions   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupResponseItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupResponseItem", v)
}

// DecodeMBSSessionSetupResponseItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupResponseItem(data []byte) (MBSSessionSetupResponseItem, int, error) {
	var v MBSSessionSetupResponseItem
	err := ngapcodec.DecodeParameter("MBSSessionSetupResponseItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupResponseItem{}, 0, err
	}
	return v, len(data), err
}
