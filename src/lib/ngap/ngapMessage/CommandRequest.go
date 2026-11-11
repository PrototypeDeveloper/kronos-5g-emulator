// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// CommandRequest represents ASN.1 CommandRequest.
type CommandRequest struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                   `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier         `ngap:"id:466,criticality:reject,presence:mandatory"`
	RANAIOTDeviceNGAPID       ngapType.RANAIOTDeviceNGAPID               `ngap:"id:472,criticality:reject,presence:mandatory"`
	CommandRequestTransfer    ngapType.CommandRequestTransferOCTETSTRING `ngap:"id:460,criticality:reject,presence:mandatory"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *CommandRequest) Encode() ([]byte, error) { return ngapcodec.Encode("CommandRequest", m) }

// DecodeCommandRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeCommandRequest(data []byte) (*CommandRequest, error) {
	var m CommandRequest
	err := ngapcodec.Decode("CommandRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
