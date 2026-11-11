// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowParametersItemExtIEsEntryPresent uint32

const (
	QosFlowParametersItemExtIEsEntryPresentNothing QosFlowParametersItemExtIEsEntryPresent = iota
	QosFlowParametersItemExtIEsEntryPresentCNPacketDelayBudgetDL
	QosFlowParametersItemExtIEsEntryPresentCNPacketDelayBudgetUL
	QosFlowParametersItemExtIEsEntryPresentBurstArrivalTimeDownlink
	QosFlowParametersItemExtIEsEntryPresentUnknown
)

type QosFlowParametersItemExtIEsEntry struct {
	ID                       ProtocolExtensionID
	Criticality              Criticality
	Present                  QosFlowParametersItemExtIEsEntryPresent
	CNPacketDelayBudgetDL    ExtendedPacketDelayBudget `ngap:"id:187,criticality:ignore,presence:optional"`
	CNPacketDelayBudgetUL    ExtendedPacketDelayBudget `ngap:"id:188,criticality:ignore,presence:optional"`
	BurstArrivalTimeDownlink BurstArrivalTime          `ngap:"id:279,criticality:ignore,presence:optional"`
	UnknownValue             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowParametersItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowParametersItem-ExtIEsEntry", v)
}

// DecodeQosFlowParametersItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowParametersItemExtIEsEntry(data []byte) (QosFlowParametersItemExtIEsEntry, int, error) {
	var v QosFlowParametersItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowParametersItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowParametersItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
