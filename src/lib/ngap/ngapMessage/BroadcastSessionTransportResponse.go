// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionTransportResponse represents ASN.1 BroadcastSessionTransportResponse.
type BroadcastSessionTransportResponse struct {
	MBSSessionID                       ngapType.MBSSessionID                                  `ngap:"id:299,criticality:reject,presence:mandatory"`
	BroadcastTransportResponseTransfer ngapType.BroadcastTransportResponseTransferOCTETSTRING `ngap:"id:419,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics          bool
	CriticalityDiagnostics             ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionTransportResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionTransportResponse", m)
}

// DecodeBroadcastSessionTransportResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionTransportResponse(data []byte) (*BroadcastSessionTransportResponse, error) {
	var m BroadcastSessionTransportResponse
	err := ngapcodec.Decode("BroadcastSessionTransportResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
