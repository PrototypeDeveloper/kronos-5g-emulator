// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowNotifyItem represents ASN.1 QosFlowNotifyItem.
type QosFlowNotifyItem struct {
	QosFlowIdentifier QosFlowIdentifier
	NotificationCause NotificationCause
	HasIEExtensions   bool
	IEExtensions      []QosFlowNotifyItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowNotifyItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowNotifyItem", v)
}

// DecodeQosFlowNotifyItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowNotifyItem(data []byte) (QosFlowNotifyItem, int, error) {
	var v QosFlowNotifyItem
	err := ngapcodec.DecodeParameter("QosFlowNotifyItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowNotifyItem{}, 0, err
	}
	return v, len(data), err
}
