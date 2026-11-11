// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QosFlowFeedbackItem represents ASN.1 QosFlowFeedbackItem.
type QosFlowFeedbackItem struct {
	QosFlowIdentifier        QosFlowIdentifier
	HasUpdateFeedback        bool
	UpdateFeedback           UpdateFeedback
	HasCNpacketDelayBudgetDL bool
	CNpacketDelayBudgetDL    ExtendedPacketDelayBudget
	HasCNpacketDelayBudgetUL bool
	CNpacketDelayBudgetUL    ExtendedPacketDelayBudget
	HasIEExtensions          bool
	IEExtensions             []QosFlowFeedbackItemExtIEsEntry
	UnknownExtensions        []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowFeedbackItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowFeedbackItem", v)
}

// DecodeQosFlowFeedbackItem decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowFeedbackItem(data []byte) (QosFlowFeedbackItem, int, error) {
	var v QosFlowFeedbackItem
	err := ngapcodec.DecodeParameter("QosFlowFeedbackItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowFeedbackItem{}, 0, err
	}
	return v, len(data), err
}
