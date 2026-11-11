// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// SNPNListforMDTItem represents ASN.1 SNPNListforMDTItem.
type SNPNListforMDTItem struct {
	PLMNIdentity      PLMNIdentity
	NID               NID
	HasIEExtensions   bool
	IEExtensions      []SNPNListforMDTItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v SNPNListforMDTItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("SNPNListforMDTItem", v)
}

// DecodeSNPNListforMDTItem decodes a standalone APER value and reports the consumed octets.
func DecodeSNPNListforMDTItem(data []byte) (SNPNListforMDTItem, int, error) {
	var v SNPNListforMDTItem
	err := ngapcodec.DecodeParameter("SNPNListforMDTItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return SNPNListforMDTItem{}, 0, err
	}
	return v, len(data), err
}
