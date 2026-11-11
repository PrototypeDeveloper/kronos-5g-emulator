// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionReleaseResponse represents ASN.1 BroadcastSessionReleaseResponse.
type BroadcastSessionReleaseResponse struct {
	MBSSessionID                         ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSSessionReleaseResponseTransfer bool
	MBSSessionReleaseResponseTransfer    ngapType.MBSSessionReleaseResponseTransferOCTETSTRING `ngap:"id:358,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics            bool
	CriticalityDiagnostics               ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                        []ngapType.UnknownIE
	UnknownExtensions                    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionReleaseResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionReleaseResponse", m)
}

// DecodeBroadcastSessionReleaseResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionReleaseResponse(data []byte) (*BroadcastSessionReleaseResponse, error) {
	var m BroadcastSessionReleaseResponse
	err := ngapcodec.Decode("BroadcastSessionReleaseResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
