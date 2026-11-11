// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// WLANMeasConfigNameItem represents ASN.1 WLANMeasConfigNameItem.
type WLANMeasConfigNameItem struct {
	WLANName          WLANName
	HasIEExtensions   bool
	IEExtensions      []WLANMeasConfigNameItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v WLANMeasConfigNameItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("WLANMeasConfigNameItem", v)
}

// DecodeWLANMeasConfigNameItem decodes a standalone APER value and reports the consumed octets.
func DecodeWLANMeasConfigNameItem(data []byte) (WLANMeasConfigNameItem, int, error) {
	var v WLANMeasConfigNameItem
	err := ngapcodec.DecodeParameter("WLANMeasConfigNameItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return WLANMeasConfigNameItem{}, 0, err
	}
	return v, len(data), err
}
