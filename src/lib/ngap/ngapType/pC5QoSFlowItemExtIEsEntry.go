// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PC5QoSFlowItemExtIEsEntryPresent uint32

const (
	PC5QoSFlowItemExtIEsEntryPresentNothing PC5QoSFlowItemExtIEsEntryPresent = iota
	PC5QoSFlowItemExtIEsEntryPresentUnknown
)

type PC5QoSFlowItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PC5QoSFlowItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PC5QoSFlowItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PC5QoSFlowItem-ExtIEsEntry", v)
}

// DecodePC5QoSFlowItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePC5QoSFlowItemExtIEsEntry(data []byte) (PC5QoSFlowItemExtIEsEntry, int, error) {
	var v PC5QoSFlowItemExtIEsEntry
	err := ngapcodec.DecodeParameter("PC5QoSFlowItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PC5QoSFlowItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
