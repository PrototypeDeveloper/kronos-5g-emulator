// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowPerTNLInformationItemExtIEsEntryPresent uint32

const (
	QosFlowPerTNLInformationItemExtIEsEntryPresentNothing QosFlowPerTNLInformationItemExtIEsEntryPresent = iota
	QosFlowPerTNLInformationItemExtIEsEntryPresentUnknown
)

type QosFlowPerTNLInformationItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowPerTNLInformationItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowPerTNLInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowPerTNLInformationItem-ExtIEsEntry", v)
}

// DecodeQosFlowPerTNLInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowPerTNLInformationItemExtIEsEntry(data []byte) (QosFlowPerTNLInformationItemExtIEsEntry, int, error) {
	var v QosFlowPerTNLInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowPerTNLInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowPerTNLInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
