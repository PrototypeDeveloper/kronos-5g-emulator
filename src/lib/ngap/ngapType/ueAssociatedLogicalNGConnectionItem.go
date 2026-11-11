// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UEAssociatedLogicalNGConnectionItem represents ASN.1 UE-associatedLogicalNG-connectionItem.
type UEAssociatedLogicalNGConnectionItem struct {
	HasAMFUENGAPID    bool
	AMFUENGAPID       AMFUENGAPID
	HasRANUENGAPID    bool
	RANUENGAPID       RANUENGAPID
	HasIEExtensions   bool
	IEExtensions      []UEAssociatedLogicalNGConnectionItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UEAssociatedLogicalNGConnectionItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UE-associatedLogicalNG-connectionItem", v)
}

// DecodeUEAssociatedLogicalNGConnectionItem decodes a standalone APER value and reports the consumed octets.
func DecodeUEAssociatedLogicalNGConnectionItem(data []byte) (UEAssociatedLogicalNGConnectionItem, int, error) {
	var v UEAssociatedLogicalNGConnectionItem
	err := ngapcodec.DecodeParameter("UE-associatedLogicalNG-connectionItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return UEAssociatedLogicalNGConnectionItem{}, 0, err
	}
	return v, len(data), err
}
