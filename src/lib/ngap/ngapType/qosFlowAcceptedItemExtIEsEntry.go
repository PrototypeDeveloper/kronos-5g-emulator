// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowAcceptedItemExtIEsEntryPresent uint32

const (
	QosFlowAcceptedItemExtIEsEntryPresentNothing QosFlowAcceptedItemExtIEsEntryPresent = iota
	QosFlowAcceptedItemExtIEsEntryPresentCurrentQoSParaSetIndex
	QosFlowAcceptedItemExtIEsEntryPresentUnknown
)

type QosFlowAcceptedItemExtIEsEntry struct {
	ID                     ProtocolExtensionID
	Criticality            Criticality
	Present                QosFlowAcceptedItemExtIEsEntryPresent
	CurrentQoSParaSetIndex AlternativeQoSParaSetIndex `ngap:"id:221,criticality:ignore,presence:optional"`
	UnknownValue           []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAcceptedItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAcceptedItem-ExtIEsEntry", v)
}

// DecodeQosFlowAcceptedItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAcceptedItemExtIEsEntry(data []byte) (QosFlowAcceptedItemExtIEsEntry, int, error) {
	var v QosFlowAcceptedItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowAcceptedItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAcceptedItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
