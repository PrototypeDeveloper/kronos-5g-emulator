// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// ExcessPacketDelayThresholdItem represents ASN.1 ExcessPacketDelayThresholdItem.
type ExcessPacketDelayThresholdItem struct {
	FiveQi                          FiveQI
	ExcessPacketDelayThresholdValue ExcessPacketDelayThresholdValue
	HasIEExtensions                 bool
	IEExtensions                    []ExcessPacketDelayThresholdItemExtIEsEntry
	UnknownExtensions               []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExcessPacketDelayThresholdItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExcessPacketDelayThresholdItem", v)
}

// DecodeExcessPacketDelayThresholdItem decodes a standalone APER value and reports the consumed octets.
func DecodeExcessPacketDelayThresholdItem(data []byte) (ExcessPacketDelayThresholdItem, int, error) {
	var v ExcessPacketDelayThresholdItem
	err := ngapcodec.DecodeParameter("ExcessPacketDelayThresholdItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExcessPacketDelayThresholdItem{}, 0, err
	}
	return v, len(data), err
}
