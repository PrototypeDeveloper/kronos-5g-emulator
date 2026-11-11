// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionTNLInfo5GCItem represents ASN.1 MBS-SessionTNLInfo5GCItem.
type MBSSessionTNLInfo5GCItem struct {
	MBSAreaSessionID                 MBSAreaSessionID
	SharedNGUMulticastTNLInformation SharedNGUMulticastTNLInformation
	HasIEExtensions                  bool
	IEExtensions                     []MBSSessionTNLInfo5GCItemExtIEsEntry
	UnknownExtensions                []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionTNLInfo5GCItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-SessionTNLInfo5GCItem", v)
}

// DecodeMBSSessionTNLInfo5GCItem decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionTNLInfo5GCItem(data []byte) (MBSSessionTNLInfo5GCItem, int, error) {
	var v MBSSessionTNLInfo5GCItem
	err := ngapcodec.DecodeParameter("MBS-SessionTNLInfo5GCItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionTNLInfo5GCItem{}, 0, err
	}
	return v, len(data), err
}
