// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// EndpointIPAddressAndPort represents ASN.1 EndpointIPAddressAndPort.
type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress
	PortNumber        PortNumber
	HasIEExtensions   bool
	IEExtensions      []EndpointIPAddressAndPortExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EndpointIPAddressAndPort) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EndpointIPAddressAndPort", v)
}

// DecodeEndpointIPAddressAndPort decodes a standalone APER value and reports the consumed octets.
func DecodeEndpointIPAddressAndPort(data []byte) (EndpointIPAddressAndPort, int, error) {
	var v EndpointIPAddressAndPort
	err := ngapcodec.DecodeParameter("EndpointIPAddressAndPort", data, &v)
	if ngapcodec.IsFatal(err) {
		return EndpointIPAddressAndPort{}, 0, err
	}
	return v, len(data), err
}
