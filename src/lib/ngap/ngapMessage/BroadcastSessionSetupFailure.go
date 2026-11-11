// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionSetupFailure represents ASN.1 BroadcastSessionSetupFailure.
type BroadcastSessionSetupFailure struct {
	MBSSessionID                      ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSSessionSetupFailureTransfer bool
	MBSSessionSetupFailureTransfer    ngapType.MBSSessionSetupOrModFailureTransferOCTETSTRING `ngap:"id:314,criticality:ignore,presence:optional"`
	Cause                             ngapType.Cause                                          `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics         bool
	CriticalityDiagnostics            ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                     []ngapType.UnknownIE
	UnknownExtensions                 []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionSetupFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionSetupFailure", m)
}

// DecodeBroadcastSessionSetupFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionSetupFailure(data []byte) (*BroadcastSessionSetupFailure, error) {
	var m BroadcastSessionSetupFailure
	err := ngapcodec.Decode("BroadcastSessionSetupFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
