// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type QosFlowToBeForwardedItemExtIEsEntryPresent uint32

const (
	QosFlowToBeForwardedItemExtIEsEntryPresentNothing QosFlowToBeForwardedItemExtIEsEntryPresent = iota
	QosFlowToBeForwardedItemExtIEsEntryPresentUnknown
)

type QosFlowToBeForwardedItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      QosFlowToBeForwardedItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v QosFlowToBeForwardedItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("QosFlowToBeForwardedItem-ExtIEsEntry", v)
}

// DecodeQosFlowToBeForwardedItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeQosFlowToBeForwardedItemExtIEsEntry(data []byte) (QosFlowToBeForwardedItemExtIEsEntry, int, error) {
	var v QosFlowToBeForwardedItemExtIEsEntry
	err := ngapcodec.DecodeParameter("QosFlowToBeForwardedItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return QosFlowToBeForwardedItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
