// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type A2XPC5QoSFlowItemExtIEsEntryPresent uint32

const (
	A2XPC5QoSFlowItemExtIEsEntryPresentNothing A2XPC5QoSFlowItemExtIEsEntryPresent = iota
	A2XPC5QoSFlowItemExtIEsEntryPresentUnknown
)

type A2XPC5QoSFlowItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      A2XPC5QoSFlowItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v A2XPC5QoSFlowItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("A2X-PC5-QoS-Flow-Item-ExtIEsEntry", v)
}

// DecodeA2XPC5QoSFlowItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeA2XPC5QoSFlowItemExtIEsEntry(data []byte) (A2XPC5QoSFlowItemExtIEsEntry, int, error) {
	var v A2XPC5QoSFlowItemExtIEsEntry
	err := ngapcodec.DecodeParameter("A2X-PC5-QoS-Flow-Item-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return A2XPC5QoSFlowItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
