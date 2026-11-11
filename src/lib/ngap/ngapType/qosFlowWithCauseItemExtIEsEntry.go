// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowWithCauseItemExtIEsEntryPresent uint32

const (
	QosFlowWithCauseItemExtIEsEntryPresentNothing QosFlowWithCauseItemExtIEsEntryPresent = iota
	QosFlowWithCauseItemExtIEsEntryPresentUnknown
)

type QosFlowWithCauseItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowWithCauseItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowWithCauseItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowWithCauseItem-ExtIEsEntry", v)
}

// DecodeQosFlowWithCauseItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowWithCauseItemExtIEsEntry(data []byte) (QosFlowWithCauseItemExtIEsEntry, int, error) {
	var v QosFlowWithCauseItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowWithCauseItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowWithCauseItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
