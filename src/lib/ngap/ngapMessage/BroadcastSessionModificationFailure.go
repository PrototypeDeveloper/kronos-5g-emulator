// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionModificationFailure represents ASN.1 BroadcastSessionModificationFailure.
type BroadcastSessionModificationFailure struct {
	MBSSessionID                             ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSSessionModificationFailureTransfer bool
	MBSSessionModificationFailureTransfer    ngapType.MBSSessionSetupOrModFailureTransferOCTETSTRING `ngap:"id:348,criticality:ignore,presence:optional"`
	Cause                                    ngapType.Cause                                          `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics                bool
	CriticalityDiagnostics                   ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                            []ngapType.UnknownIE
	UnknownExtensions                        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionModificationFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionModificationFailure", m)
}

// DecodeBroadcastSessionModificationFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionModificationFailure(data []byte) (*BroadcastSessionModificationFailure, error) {
	var m BroadcastSessionModificationFailure
	err := ngapcodec.Decode("BroadcastSessionModificationFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
