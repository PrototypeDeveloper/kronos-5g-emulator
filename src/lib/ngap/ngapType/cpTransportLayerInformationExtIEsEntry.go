// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type CPTransportLayerInformationExtIEsEntryPresent uint32

const (
	CPTransportLayerInformationExtIEsEntryPresentNothing CPTransportLayerInformationExtIEsEntryPresent = iota
	CPTransportLayerInformationExtIEsEntryPresentEndpointIPAddressAndPort
	CPTransportLayerInformationExtIEsEntryPresentUnknown
)

type CPTransportLayerInformationExtIEsEntry struct {
	ID                       ProtocolIEID
	Criticality              Criticality
	Present                  CPTransportLayerInformationExtIEsEntryPresent
	EndpointIPAddressAndPort EndpointIPAddressAndPort `ngap:"id:169,criticality:reject,presence:mandatory"`
	UnknownValue             []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CPTransportLayerInformationExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CPTransportLayerInformation-ExtIEsEntry", v)
}

// DecodeCPTransportLayerInformationExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeCPTransportLayerInformationExtIEsEntry(data []byte) (CPTransportLayerInformationExtIEsEntry, int, error) {
	var v CPTransportLayerInformationExtIEsEntry
	err := ngapcodec.DecodeParameter("CPTransportLayerInformation-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return CPTransportLayerInformationExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
