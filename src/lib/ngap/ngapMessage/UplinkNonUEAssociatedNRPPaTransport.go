// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkNonUEAssociatedNRPPaTransport represents ASN.1 UplinkNonUEAssociatedNRPPaTransport.
type UplinkNonUEAssociatedNRPPaTransport struct {
	RoutingID         ngapType.RoutingID `ngap:"id:89,criticality:reject,presence:mandatory"`
	NRPPaPDU          ngapType.NRPPaPDU  `ngap:"id:46,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkNonUEAssociatedNRPPaTransport) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkNonUEAssociatedNRPPaTransport", m)
}

// DecodeUplinkNonUEAssociatedNRPPaTransport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkNonUEAssociatedNRPPaTransport(data []byte) (*UplinkNonUEAssociatedNRPPaTransport, error) {
	var m UplinkNonUEAssociatedNRPPaTransport
	err := ngapcodec.Decode("UplinkNonUEAssociatedNRPPaTransport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
