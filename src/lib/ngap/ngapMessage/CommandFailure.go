// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// CommandFailure represents ASN.1 CommandFailure.
type CommandFailure struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                   `ngap:"id:467,criticality:ignore,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier         `ngap:"id:466,criticality:ignore,presence:mandatory"`
	RANAIOTDeviceNGAPID       ngapType.RANAIOTDeviceNGAPID               `ngap:"id:472,criticality:ignore,presence:mandatory"`
	CommandFailureTransfer    ngapType.CommandFailureTransferOCTETSTRING `ngap:"id:462,criticality:ignore,presence:mandatory"`
	Cause                     ngapType.Cause                             `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *CommandFailure) Encode() ([]byte, error) { return ngapcodec.Encode("CommandFailure", m) }

// DecodeCommandFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeCommandFailure(data []byte) (*CommandFailure, error) {
	var m CommandFailure
	err := ngapcodec.Decode("CommandFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
