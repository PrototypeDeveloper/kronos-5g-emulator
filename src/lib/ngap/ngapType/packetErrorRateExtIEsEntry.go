// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type PacketErrorRateExtIEsEntryPresent uint32

const (
	PacketErrorRateExtIEsEntryPresentNothing PacketErrorRateExtIEsEntryPresent = iota
	PacketErrorRateExtIEsEntryPresentUnknown
)

type PacketErrorRateExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      PacketErrorRateExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v PacketErrorRateExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("PacketErrorRate-ExtIEsEntry", v)
}

// DecodePacketErrorRateExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodePacketErrorRateExtIEsEntry(data []byte) (PacketErrorRateExtIEsEntry, int, error) {
	var v PacketErrorRateExtIEsEntry
	err := ngapcodec.DecodeParameter("PacketErrorRate-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return PacketErrorRateExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
