// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNCellIdListforMDTItem represents ASN.1 SNPNCellIdListforMDTItem.
type SNPNCellIdListforMDTItem struct {
	NRCGI             NRCGI
	NID               NID
	HasIEExtensions   bool
	IEExtensions      []SNPNCellIdListforMDTItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNCellIdListforMDTItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNCellIdListforMDTItem", v)
}

// DecodeSNPNCellIdListforMDTItem decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNCellIdListforMDTItem(data []byte) (SNPNCellIdListforMDTItem, int, error) {
	var v SNPNCellIdListforMDTItem
	err := ngapcodec.DecodeParameter("SNPNCellIdListforMDTItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNCellIdListforMDTItem{}, 0, err
	}
	return v, len(data), err
}
