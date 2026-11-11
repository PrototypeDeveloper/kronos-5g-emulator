// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AreaOfInterestRANNodeItem represents ASN.1 AreaOfInterestRANNodeItem.
type AreaOfInterestRANNodeItem struct {
	GlobalRANNodeID   GlobalRANNodeID
	HasIEExtensions   bool
	IEExtensions      []AreaOfInterestRANNodeItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AreaOfInterestRANNodeItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AreaOfInterestRANNodeItem", v)
}

// DecodeAreaOfInterestRANNodeItem decodes a standalone APER value and reports the consumed octets.
func DecodeAreaOfInterestRANNodeItem(data []byte) (AreaOfInterestRANNodeItem, int, error) {
	var v AreaOfInterestRANNodeItem
	err := ngapcodec.DecodeParameter("AreaOfInterestRANNodeItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return AreaOfInterestRANNodeItem{}, 0, err
	}
	return v, len(data), err
}
