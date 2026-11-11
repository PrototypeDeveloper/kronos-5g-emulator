// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowPerTNLInformationExtIEsEntryPresent uint32

const (
	QosFlowPerTNLInformationExtIEsEntryPresentNothing QosFlowPerTNLInformationExtIEsEntryPresent = iota
	QosFlowPerTNLInformationExtIEsEntryPresentUnknown
)

type QosFlowPerTNLInformationExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowPerTNLInformationExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowPerTNLInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowPerTNLInformation-ExtIEsEntry", v)
}

// DecodeQosFlowPerTNLInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowPerTNLInformationExtIEsEntry(data []byte) (QosFlowPerTNLInformationExtIEsEntry, int, error) {
	var v QosFlowPerTNLInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowPerTNLInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowPerTNLInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
