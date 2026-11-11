// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionModificationResponse represents ASN.1 BroadcastSessionModificationResponse.
type BroadcastSessionModificationResponse struct {
	MBSSessionID                              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSSessionModificationResponseTransfer bool
	MBSSessionModificationResponseTransfer    ngapType.MBSSessionSetupOrModResponseTransferOCTETSTRING `ngap:"id:350,criticality:reject,presence:optional"`
	HasCriticalityDiagnostics                 bool
	CriticalityDiagnostics                    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                             []ngapType.UnknownIE
	UnknownExtensions                         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionModificationResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionModificationResponse", m)
}

// DecodeBroadcastSessionModificationResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionModificationResponse(data []byte) (*BroadcastSessionModificationResponse, error) {
	var m BroadcastSessionModificationResponse
	err := ngapcodec.Decode("BroadcastSessionModificationResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
