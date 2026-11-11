// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// CPTransportLayerInformation represents ASN.1 CPTransportLayerInformation.
type CPTransportLayerInformationPresent uint32

const (
	CPTransportLayerInformationPresentNothing CPTransportLayerInformationPresent = iota
	CPTransportLayerInformationPresentEndpointIPAddress
	CPTransportLayerInformationPresentChoiceExtensions
)

type CPTransportLayerInformation struct {
	Present           CPTransportLayerInformationPresent
	EndpointIPAddress TransportLayerAddress
	ChoiceExtensions  CPTransportLayerInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v CPTransportLayerInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("CPTransportLayerInformation", v)
}

// DecodeCPTransportLayerInformation decodes a standalone APER value and reports the consumed octets.
func DecodeCPTransportLayerInformation(data []byte) (CPTransportLayerInformation, int, error) {
	var v CPTransportLayerInformation
	err := ngapcodec.DecodeParameter("CPTransportLayerInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return CPTransportLayerInformation{}, 0, err
	}
	return v, len(data), err
}
