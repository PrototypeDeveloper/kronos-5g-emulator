// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAcceptedItem represents ASN.1 QosFlowAcceptedItem.
type QosFlowAcceptedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	HasIEExtensions   bool
	IEExtensions      []QosFlowAcceptedItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAcceptedItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAcceptedItem", v)
}

// DecodeQosFlowAcceptedItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAcceptedItem(data []byte) (QosFlowAcceptedItem, int, error) {
	var v QosFlowAcceptedItem
	err := ngapcodec.DecodeParameter("QosFlowAcceptedItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAcceptedItem{}, 0, err
	}
	return v, len(data), err
}
