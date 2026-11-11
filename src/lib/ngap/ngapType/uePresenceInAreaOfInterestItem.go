// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEPresenceInAreaOfInterestItem represents ASN.1 UEPresenceInAreaOfInterestItem.
type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID LocationReportingReferenceID
	UEPresence                   UEPresence
	HasIEExtensions              bool
	IEExtensions                 []UEPresenceInAreaOfInterestItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEPresenceInAreaOfInterestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UEPresenceInAreaOfInterestItem", v)
}

// DecodeUEPresenceInAreaOfInterestItem decodes a standalone APER value and reports the consumed octets.
func DecodeUEPresenceInAreaOfInterestItem(data []byte) (UEPresenceInAreaOfInterestItem, int, error) {
	var v UEPresenceInAreaOfInterestItem
	err := ngapcodec.DecodeParameter("UEPresenceInAreaOfInterestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEPresenceInAreaOfInterestItem{}, 0, err
	}
	return v, len(data), err
}
