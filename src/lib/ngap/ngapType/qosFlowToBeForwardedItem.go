// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowToBeForwardedItem represents ASN.1 QosFlowToBeForwardedItem.
type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	HasIEExtensions   bool
	IEExtensions      []QosFlowToBeForwardedItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowToBeForwardedItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowToBeForwardedItem", v)
}

// DecodeQosFlowToBeForwardedItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowToBeForwardedItem(data []byte) (QosFlowToBeForwardedItem, int, error) {
	var v QosFlowToBeForwardedItem
	err := ngapcodec.DecodeParameter("QosFlowToBeForwardedItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowToBeForwardedItem{}, 0, err
	}
	return v, len(data), err
}
