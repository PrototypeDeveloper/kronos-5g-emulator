// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowAdditionalInfoItemRelResExtIEsEntryPresent uint32

const (
	QosFlowAdditionalInfoItemRelResExtIEsEntryPresentNothing QosFlowAdditionalInfoItemRelResExtIEsEntryPresent = iota
	QosFlowAdditionalInfoItemRelResExtIEsEntryPresentUnknown
)

type QosFlowAdditionalInfoItemRelResExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowAdditionalInfoItemRelResExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAdditionalInfoItemRelResExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAdditionalInfoItemRelRes-ExtIEsEntry", v)
}

// DecodeQosFlowAdditionalInfoItemRelResExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAdditionalInfoItemRelResExtIEsEntry(data []byte) (QosFlowAdditionalInfoItemRelResExtIEsEntry, int, error) {
	var v QosFlowAdditionalInfoItemRelResExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowAdditionalInfoItemRelRes-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAdditionalInfoItemRelResExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
