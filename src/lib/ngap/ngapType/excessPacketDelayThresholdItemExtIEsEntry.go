// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type ExcessPacketDelayThresholdItemExtIEsEntryPresent uint32

const (
	ExcessPacketDelayThresholdItemExtIEsEntryPresentNothing ExcessPacketDelayThresholdItemExtIEsEntryPresent = iota
	ExcessPacketDelayThresholdItemExtIEsEntryPresentUnknown
)

type ExcessPacketDelayThresholdItemExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      ExcessPacketDelayThresholdItemExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v ExcessPacketDelayThresholdItemExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("ExcessPacketDelayThresholdItem-ExtIEsEntry", v)
}

// DecodeExcessPacketDelayThresholdItemExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeExcessPacketDelayThresholdItemExtIEsEntry(data []byte) (ExcessPacketDelayThresholdItemExtIEsEntry, int, error) {
	var v ExcessPacketDelayThresholdItemExtIEsEntry
	err := ngapcodec.DecodeParameter("ExcessPacketDelayThresholdItem-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return ExcessPacketDelayThresholdItemExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
