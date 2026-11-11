// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowItemWithDataForwardingExtIEsEntryPresent uint32

const (
	QosFlowItemWithDataForwardingExtIEsEntryPresentNothing QosFlowItemWithDataForwardingExtIEsEntryPresent = iota
	QosFlowItemWithDataForwardingExtIEsEntryPresentCurrentQoSParaSetIndex
	QosFlowItemWithDataForwardingExtIEsEntryPresentUnknown
)

type QosFlowItemWithDataForwardingExtIEsEntry struct {
	ID                     ProtocolExtensionID
	Criticality            Criticality
	Present                QosFlowItemWithDataForwardingExtIEsEntryPresent
	CurrentQoSParaSetIndex AlternativeQoSParaSetIndex `ngap:"id:221,criticality:ignore,presence:optional"`
	UnknownValue           []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowItemWithDataForwardingExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowItemWithDataForwarding-ExtIEsEntry", v)
}

// DecodeQosFlowItemWithDataForwardingExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowItemWithDataForwardingExtIEsEntry(data []byte) (QosFlowItemWithDataForwardingExtIEsEntry, int, error) {
	var v QosFlowItemWithDataForwardingExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowItemWithDataForwarding-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowItemWithDataForwardingExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
