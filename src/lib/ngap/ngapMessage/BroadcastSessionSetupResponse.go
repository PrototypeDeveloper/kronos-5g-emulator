// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionSetupResponse represents ASN.1 BroadcastSessionSetupResponse.
type BroadcastSessionSetupResponse struct {
	MBSSessionID                       ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSSessionSetupResponseTransfer bool
	MBSSessionSetupResponseTransfer    ngapType.MBSSessionSetupOrModResponseTransferOCTETSTRING `ngap:"id:316,criticality:reject,presence:optional"`
	HasCriticalityDiagnostics          bool
	CriticalityDiagnostics             ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionSetupResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionSetupResponse", m)
}

// DecodeBroadcastSessionSetupResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionSetupResponse(data []byte) (*BroadcastSessionSetupResponse, error) {
	var m BroadcastSessionSetupResponse
	err := ngapcodec.Decode("BroadcastSessionSetupResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
