// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestItem represents ASN.1 AreaOfInterestItem.
type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest
	LocationReportingReferenceID LocationReportingReferenceID
	HasIEExtensions              bool
	IEExtensions                 []AreaOfInterestItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestItem", v)
}

// DecodeAreaOfInterestItem decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestItem(data []byte) (AreaOfInterestItem, int, error) {
	var v AreaOfInterestItem
	err := ngapcodec.DecodeParameter("AreaOfInterestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestItem{}, 0, err
	}
	return v, len(data), err
}
