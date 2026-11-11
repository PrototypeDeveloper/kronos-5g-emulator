// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type EndpointIPAddressAndPortExtIEsEntryPresent uint32

const (
	EndpointIPAddressAndPortExtIEsEntryPresentNothing EndpointIPAddressAndPortExtIEsEntryPresent = iota
	EndpointIPAddressAndPortExtIEsEntryPresentUnknown
)

type EndpointIPAddressAndPortExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      EndpointIPAddressAndPortExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v EndpointIPAddressAndPortExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("EndpointIPAddressAndPort-ExtIEsEntry", v)
}

// DecodeEndpointIPAddressAndPortExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeEndpointIPAddressAndPortExtIEsEntry(data []byte) (EndpointIPAddressAndPortExtIEsEntry, int, error) {
	var v EndpointIPAddressAndPortExtIEsEntry
	err := ngapcodec.DecodeParameter("EndpointIPAddressAndPort-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return EndpointIPAddressAndPortExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
