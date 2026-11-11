// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AllowedPNINPNItem represents ASN.1 Allowed-PNI-NPN-Item.
type AllowedPNINPNItem struct {
	PLMNIdentity          PLMNIdentity
	PNINPNRestricted      AllowedPNINPNItemPNINPNRestricted `aper:"valueExt,valueLB:0,valueUB:1"`
	AllowedCAGListPerPLMN AllowedCAGListPerPLMN
	HasIEExtensions       bool
	IEExtensions          []AllowedPNINPNItemExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AllowedPNINPNItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("Allowed-PNI-NPN-Item", v)
}

// DecodeAllowedPNINPNItem decodes a standalone APER value and reports the consumed octets.
func DecodeAllowedPNINPNItem(data []byte) (AllowedPNINPNItem, int, error) {
	var v AllowedPNINPNItem
	err := ngapcodec.DecodeParameter("Allowed-PNI-NPN-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return AllowedPNINPNItem{}, 0, err
	}
	return v, len(data), err
}
