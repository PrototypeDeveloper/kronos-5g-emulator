// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAddOrModifyResponseItem represents ASN.1 QosFlowAddOrModifyResponseItem.
type QosFlowAddOrModifyResponseItem struct {
	QosFlowIdentifier QosFlowIdentifier
	HasIEExtensions   bool
	IEExtensions      []QosFlowAddOrModifyResponseItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAddOrModifyResponseItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAddOrModifyResponseItem", v)
}

// DecodeQosFlowAddOrModifyResponseItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAddOrModifyResponseItem(data []byte) (QosFlowAddOrModifyResponseItem, int, error) {
	var v QosFlowAddOrModifyResponseItem
	err := ngapcodec.DecodeParameter("QosFlowAddOrModifyResponseItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAddOrModifyResponseItem{}, 0, err
	}
	return v, len(data), err
}
