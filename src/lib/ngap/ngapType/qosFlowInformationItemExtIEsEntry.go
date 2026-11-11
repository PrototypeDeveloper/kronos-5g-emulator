// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowInformationItemExtIEsEntryPresent uint32

const (
	QosFlowInformationItemExtIEsEntryPresentNothing QosFlowInformationItemExtIEsEntryPresent = iota
	QosFlowInformationItemExtIEsEntryPresentULForwarding
	QosFlowInformationItemExtIEsEntryPresentSourceTNLAddrInfo
	QosFlowInformationItemExtIEsEntryPresentSourceNodeTNLAddrInfo
	QosFlowInformationItemExtIEsEntryPresentUnknown
)

type QosFlowInformationItemExtIEsEntry struct {
	ID                    ProtocolExtensionID
	Criticality           Criticality
	Present               QosFlowInformationItemExtIEsEntryPresent
	ULForwarding          ULForwarding          `ngap:"id:163,criticality:ignore,presence:optional"`
	SourceTNLAddrInfo     TransportLayerAddress `ngap:"id:284,criticality:ignore,presence:optional"`
	SourceNodeTNLAddrInfo TransportLayerAddress `ngap:"id:354,criticality:ignore,presence:optional"`
	UnknownValue          []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowInformationItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowInformationItem-ExtIEsEntry", v)
}

// DecodeQosFlowInformationItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowInformationItemExtIEsEntry(data []byte) (QosFlowInformationItemExtIEsEntry, int, error) {
	var v QosFlowInformationItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowInformationItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowInformationItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
