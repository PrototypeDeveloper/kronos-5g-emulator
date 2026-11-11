// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// QoSFlowTSCItem represents ASN.1 QoSFlowTSCItem.
type QoSFlowTSCItem struct {
	QosFlowIdentifier                    QosFlowIdentifier
	HasTSCTrafficCharacteristicsFeedback bool
	TSCTrafficCharacteristicsFeedback    TSCTrafficCharacteristicsFeedback
	HasANPacketDelayBudgetUL             bool
	ANPacketDelayBudgetUL                ExtendedPacketDelayBudget
	HasIEExtensions                      bool
	IEExtensions                         []QoSFlowTSCItemExtIEsEntry
	UnknownExtensions                    []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QoSFlowTSCItem) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QoSFlowTSCItem", v)
}

// DecodeQoSFlowTSCItem decodes a standalone APER value and reports the consumed octets.
func DecodeQoSFlowTSCItem(data []byte) (QoSFlowTSCItem, int, error) {
	var v QoSFlowTSCItem
	err := ngapcodec.DecodeParameter("QoSFlowTSCItem", data, &v)
	if ngapcodec.IsFatal(err) {
		return QoSFlowTSCItem{}, 0, err
	}
	return v, len(data), err
}
