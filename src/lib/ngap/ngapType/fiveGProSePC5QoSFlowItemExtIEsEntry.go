// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type FiveGProSePC5QoSFlowItemExtIEsEntryPresent uint32

const (
	FiveGProSePC5QoSFlowItemExtIEsEntryPresentNothing FiveGProSePC5QoSFlowItemExtIEsEntryPresent = iota
	FiveGProSePC5QoSFlowItemExtIEsEntryPresentUnknown
)

type FiveGProSePC5QoSFlowItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      FiveGProSePC5QoSFlowItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v FiveGProSePC5QoSFlowItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("FiveGProSePC5QoSFlowItem-ExtIEsEntry", v)
}

// DecodeFiveGProSePC5QoSFlowItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeFiveGProSePC5QoSFlowItemExtIEsEntry(data []byte) (FiveGProSePC5QoSFlowItemExtIEsEntry, int, error) {
	var v FiveGProSePC5QoSFlowItemExtIEsEntry
	err := ngapcodec.DecodeParameter("FiveGProSePC5QoSFlowItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return FiveGProSePC5QoSFlowItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
