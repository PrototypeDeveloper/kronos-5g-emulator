// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DownlinkNonUEAssociatedNRPPaTransport represents ASN.1 DownlinkNonUEAssociatedNRPPaTransport.
type DownlinkNonUEAssociatedNRPPaTransport struct {
	RoutingID         ngapType.RoutingID `ngap:"id:89,criticality:reject,presence:mandatory"`
	NRPPaPDU          ngapType.NRPPaPDU  `ngap:"id:46,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DownlinkNonUEAssociatedNRPPaTransport) Encode() ([]byte, error) {
	return ngapcodec.Encode("DownlinkNonUEAssociatedNRPPaTransport", m)
}

// DecodeDownlinkNonUEAssociatedNRPPaTransport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDownlinkNonUEAssociatedNRPPaTransport(data []byte) (*DownlinkNonUEAssociatedNRPPaTransport, error) {
	var m DownlinkNonUEAssociatedNRPPaTransport
	err := ngapcodec.Decode("DownlinkNonUEAssociatedNRPPaTransport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
