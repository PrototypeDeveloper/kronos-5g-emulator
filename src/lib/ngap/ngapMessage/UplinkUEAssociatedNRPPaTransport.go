// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UplinkUEAssociatedNRPPaTransport represents ASN.1 UplinkUEAssociatedNRPPaTransport.
type UplinkUEAssociatedNRPPaTransport struct {
	AMFUENGAPID       ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	RoutingID         ngapType.RoutingID   `ngap:"id:89,criticality:reject,presence:mandatory"`
	NRPPaPDU          ngapType.NRPPaPDU    `ngap:"id:46,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UplinkUEAssociatedNRPPaTransport) Encode() ([]byte, error) {
	return ngapcodec.Encode("UplinkUEAssociatedNRPPaTransport", m)
}

// DecodeUplinkUEAssociatedNRPPaTransport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUplinkUEAssociatedNRPPaTransport(data []byte) (*UplinkUEAssociatedNRPPaTransport, error) {
	var m UplinkUEAssociatedNRPPaTransport
	err := ngapcodec.Decode("UplinkUEAssociatedNRPPaTransport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
