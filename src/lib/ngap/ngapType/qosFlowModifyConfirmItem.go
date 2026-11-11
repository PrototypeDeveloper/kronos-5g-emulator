// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowModifyConfirmItem represents ASN.1 QosFlowModifyConfirmItem.
type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier
	HasIEExtensions   bool
	IEExtensions      []QosFlowModifyConfirmItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowModifyConfirmItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowModifyConfirmItem", v)
}

// DecodeQosFlowModifyConfirmItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowModifyConfirmItem(data []byte) (QosFlowModifyConfirmItem, int, error) {
	var v QosFlowModifyConfirmItem
	err := ngapcodec.DecodeParameter("QosFlowModifyConfirmItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowModifyConfirmItem{}, 0, err
	}
	return v, len(data), err
}
