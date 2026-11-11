// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// GTPTunnel represents ASN.1 GTPTunnel.
type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress
	GTPTEID               GTPTEID
	HasIEExtensions       bool
	IEExtensions          []GTPTunnelExtIEsEntry
	UnknownExtensions     []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v GTPTunnel) Encode() ([]byte, error) { return ngapcodec.EncodeParameter("GTPTunnel", v) }

// DecodeGTPTunnel decodes a standalone APER value and reports the consumed octets.
func DecodeGTPTunnel(data []byte) (GTPTunnel, int, error) {
	var v GTPTunnel
	err := ngapcodec.DecodeParameter("GTPTunnel", data, &v)
	if ngapcodec.IsFatal(err) {
		return GTPTunnel{}, 0, err
	}
	return v, len(data), err
}
