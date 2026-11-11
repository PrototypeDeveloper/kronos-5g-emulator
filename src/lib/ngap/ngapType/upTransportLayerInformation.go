// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// UPTransportLayerInformation represents ASN.1 UPTransportLayerInformation.
type UPTransportLayerInformationPresent uint32

const (
	UPTransportLayerInformationPresentNothing UPTransportLayerInformationPresent = iota
	UPTransportLayerInformationPresentGTPTunnel
	UPTransportLayerInformationPresentChoiceExtensions
)

type UPTransportLayerInformation struct {
	Present          UPTransportLayerInformationPresent
	GTPTunnel        GTPTunnel
	ChoiceExtensions UPTransportLayerInformationExtIEsEntry
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v UPTransportLayerInformation) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("UPTransportLayerInformation", v)
}

// DecodeUPTransportLayerInformation decodes a standalone APER value and reports the consumed octets.
func DecodeUPTransportLayerInformation(data []byte) (UPTransportLayerInformation, int, error) {
	var v UPTransportLayerInformation
	err := ngapcodec.DecodeParameter("UPTransportLayerInformation", data, &v)
	if ngapcodec.IsFatal(err) {
		return UPTransportLayerInformation{}, 0, err
	}
	return v, len(data), err
}
