// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ForbiddenAreaInformationItem represents ASN.1 ForbiddenAreaInformation-Item.
type ForbiddenAreaInformationItem struct {
	PLMNIdentity      PLMNIdentity
	ForbiddenTACs     ForbiddenTACs
	HasIEExtensions   bool
	IEExtensions      []ForbiddenAreaInformationItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ForbiddenAreaInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ForbiddenAreaInformation-Item", v)
}

// DecodeForbiddenAreaInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeForbiddenAreaInformationItem(data []byte) (ForbiddenAreaInformationItem, int, error) {
	var v ForbiddenAreaInformationItem
	err := ngapcodec.DecodeParameter("ForbiddenAreaInformation-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return ForbiddenAreaInformationItem{}, 0, err
	}
	return v, len(data), err
}
