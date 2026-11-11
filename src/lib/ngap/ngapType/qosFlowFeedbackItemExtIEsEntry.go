// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowFeedbackItemExtIEsEntryPresent uint32

const (
	QosFlowFeedbackItemExtIEsEntryPresentNothing QosFlowFeedbackItemExtIEsEntryPresent = iota
	QosFlowFeedbackItemExtIEsEntryPresentUnknown
)

type QosFlowFeedbackItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowFeedbackItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowFeedbackItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowFeedbackItem-ExtIEsEntry", v)
}

// DecodeQosFlowFeedbackItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowFeedbackItemExtIEsEntry(data []byte) (QosFlowFeedbackItemExtIEsEntry, int, error) {
	var v QosFlowFeedbackItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowFeedbackItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowFeedbackItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
