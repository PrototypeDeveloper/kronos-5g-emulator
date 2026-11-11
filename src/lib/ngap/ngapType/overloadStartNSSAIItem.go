// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// OverloadStartNSSAIItem represents ASN.1 OverloadStartNSSAIItem.
type OverloadStartNSSAIItem struct {
	SliceOverloadList                      SliceOverloadList
	HasSliceOverloadResponse               bool
	SliceOverloadResponse                  OverloadResponse
	HasSliceTrafficLoadReductionIndication bool
	SliceTrafficLoadReductionIndication    TrafficLoadReductionIndication
	HasIEExtensions                        bool
	IEExtensions                           []OverloadStartNSSAIItemExtIEsEntry
	UnknownExtensions                      []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v OverloadStartNSSAIItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("OverloadStartNSSAIItem", v)
}

// DecodeOverloadStartNSSAIItem decodes a standalone APER value and reports the consumed octets.
func DecodeOverloadStartNSSAIItem(data []byte) (OverloadStartNSSAIItem, int, error) {
	var v OverloadStartNSSAIItem
	err := ngapcodec.DecodeParameter("OverloadStartNSSAIItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return OverloadStartNSSAIItem{}, 0, err
	}
	return v, len(data), err
}
