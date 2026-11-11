// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowWithCauseItem represents ASN.1 QosFlowWithCauseItem.
type QosFlowWithCauseItem struct {
	QosFlowIdentifier QosFlowIdentifier
	Cause             Cause
	HasIEExtensions   bool
	IEExtensions      []QosFlowWithCauseItemExtIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowWithCauseItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowWithCauseItem", v)
}

// DecodeQosFlowWithCauseItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowWithCauseItem(data []byte) (QosFlowWithCauseItem, int, error) {
	var v QosFlowWithCauseItem
	err := ngapcodec.DecodeParameter("QosFlowWithCauseItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowWithCauseItem{}, 0, err
	}
	return v, len(data), err
}
