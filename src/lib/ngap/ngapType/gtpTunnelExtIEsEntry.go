// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type GTPTunnelExtIEsEntryPresent uint32

const (
	GTPTunnelExtIEsEntryPresentNothing GTPTunnelExtIEsEntryPresent = iota
	GTPTunnelExtIEsEntryPresentUnknown
)

type GTPTunnelExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      GTPTunnelExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GTPTunnelExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("GTPTunnel-ExtIEsEntry", v)
}

// DecodeGTPTunnelExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeGTPTunnelExtIEsEntry(data []byte) (GTPTunnelExtIEsEntry, int, error) {
	var v GTPTunnelExtIEsEntry
	err := ngapcodec.DecodeParameter("GTPTunnel-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return GTPTunnelExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
