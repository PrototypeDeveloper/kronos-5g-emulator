// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ServiceAreaInformationItem represents ASN.1 ServiceAreaInformation-Item.
type ServiceAreaInformationItem struct {
	PLMNIdentity      PLMNIdentity
	HasAllowedTACs    bool
	AllowedTACs       AllowedTACs
	HasNotAllowedTACs bool
	NotAllowedTACs    NotAllowedTACs
	HasIEExtensions   bool
	IEExtensions      []ServiceAreaInformationItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ServiceAreaInformationItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ServiceAreaInformation-Item", v)
}

// DecodeServiceAreaInformationItem decodes a standalone APER value and reports the consumed octets.
func DecodeServiceAreaInformationItem(data []byte) (ServiceAreaInformationItem, int, error) {
	var v ServiceAreaInformationItem
	err := ngapcodec.DecodeParameter("ServiceAreaInformation-Item", data, &v)
	if ngapcodec.IsFatal(err) {
		return ServiceAreaInformationItem{}, 0, err
	}
	return v, len(data), err
}
