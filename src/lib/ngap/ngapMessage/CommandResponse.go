// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// CommandResponse represents ASN.1 CommandResponse.
type CommandResponse struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                    `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier          `ngap:"id:466,criticality:reject,presence:mandatory"`
	RANAIOTDeviceNGAPID       ngapType.RANAIOTDeviceNGAPID                `ngap:"id:472,criticality:reject,presence:mandatory"`
	CommandResponseTransfer   ngapType.CommandResponseTransferOCTETSTRING `ngap:"id:461,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *CommandResponse) Encode() ([]byte, error) { return ngapcodec.Encode("CommandResponse", m) }

// DecodeCommandResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeCommandResponse(data []byte) (*CommandResponse, error) {
	var m CommandResponse
	err := ngapcodec.Decode("CommandResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
