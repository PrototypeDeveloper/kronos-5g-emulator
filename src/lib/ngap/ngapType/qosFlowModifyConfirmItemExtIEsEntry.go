// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowModifyConfirmItemExtIEsEntryPresent uint32

const (
	QosFlowModifyConfirmItemExtIEsEntryPresentNothing QosFlowModifyConfirmItemExtIEsEntryPresent = iota
	QosFlowModifyConfirmItemExtIEsEntryPresentUnknown
)

type QosFlowModifyConfirmItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowModifyConfirmItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowModifyConfirmItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowModifyConfirmItem-ExtIEsEntry", v)
}

// DecodeQosFlowModifyConfirmItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowModifyConfirmItemExtIEsEntry(data []byte) (QosFlowModifyConfirmItemExtIEsEntry, int, error) {
	var v QosFlowModifyConfirmItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowModifyConfirmItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowModifyConfirmItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
