// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionTransportFailure represents ASN.1 BroadcastSessionTransportFailure.
type BroadcastSessionTransportFailure struct {
	MBSSessionID                      ngapType.MBSSessionID                                 `ngap:"id:299,criticality:reject,presence:mandatory"`
	BroadcastTransportFailureTransfer ngapType.BroadcastTransportFailureTransferOCTETSTRING `ngap:"id:417,criticality:ignore,presence:mandatory"`
	Cause                             ngapType.Cause                                        `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics         bool
	CriticalityDiagnostics            ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                     []ngapType.UnknownIE
	UnknownExtensions                 []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionTransportFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionTransportFailure", m)
}

// DecodeBroadcastSessionTransportFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionTransportFailure(data []byte) (*BroadcastSessionTransportFailure, error) {
	var m BroadcastSessionTransportFailure
	err := ngapcodec.Decode("BroadcastSessionTransportFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
