// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowAdditionalInfoItemRelComExtIEsEntryPresent uint32

const (
	QosFlowAdditionalInfoItemRelComExtIEsEntryPresentNothing QosFlowAdditionalInfoItemRelComExtIEsEntryPresent = iota
	QosFlowAdditionalInfoItemRelComExtIEsEntryPresentUnknown
)

type QosFlowAdditionalInfoItemRelComExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowAdditionalInfoItemRelComExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowAdditionalInfoItemRelComExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowAdditionalInfoItemRelCom-ExtIEsEntry", v)
}

// DecodeQosFlowAdditionalInfoItemRelComExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowAdditionalInfoItemRelComExtIEsEntry(data []byte) (QosFlowAdditionalInfoItemRelComExtIEsEntry, int, error) {
	var v QosFlowAdditionalInfoItemRelComExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowAdditionalInfoItemRelCom-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowAdditionalInfoItemRelComExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
