// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowAddOrModifyRequestItem represents ASN.1 QosFlowAddOrModifyRequestItem.
type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier            QosFlowIdentifier
	HasQosFlowLevelQosParameters bool
	QosFlowLevelQosParameters    QosFlowLevelQosParameters
	HasERABID                    bool
	ERABID                       ERABID
	HasIEExtensions              bool
	IEExtensions                 []QosFlowAddOrModifyRequestItemExtIEsEntry
	UnknownExtensions            []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAddOrModifyRequestItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAddOrModifyRequestItem", v)
}

// DecodeQosFlowAddOrModifyRequestItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAddOrModifyRequestItem(data []byte) (QosFlowAddOrModifyRequestItem, int, error) {
	var v QosFlowAddOrModifyRequestItem
	err := ngapcodec.DecodeParameter("QosFlowAddOrModifyRequestItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAddOrModifyRequestItem{}, 0, err
	}
	return v, len(data), err
}
