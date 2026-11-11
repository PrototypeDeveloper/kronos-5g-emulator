// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNTAIListforMDTItem represents ASN.1 SNPNTAIListforMDTItem.
type SNPNTAIListforMDTItem struct {
	TAI               TAI
	NID               NID
	HasIEExtensions   bool
	IEExtensions      []SNPNTAIListforMDTItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNTAIListforMDTItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNTAIListforMDTItem", v)
}

// DecodeSNPNTAIListforMDTItem decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNTAIListforMDTItem(data []byte) (SNPNTAIListforMDTItem, int, error) {
	var v SNPNTAIListforMDTItem
	err := ngapcodec.DecodeParameter("SNPNTAIListforMDTItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNTAIListforMDTItem{}, 0, err
	}
	return v, len(data), err
}
