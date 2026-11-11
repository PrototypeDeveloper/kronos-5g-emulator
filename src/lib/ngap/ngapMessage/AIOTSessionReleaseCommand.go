// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AIOTSessionReleaseCommand represents ASN.1 AIOTSessionReleaseCommand.
type AIOTSessionReleaseCommand struct {
	AIOTFIdentifier                   ngapType.AIOTFIdentifier                              `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier         ngapType.AIoTCorrelationIdentifier                    `ngap:"id:466,criticality:reject,presence:mandatory"`
	AIOTSessionReleaseCommandTransfer ngapType.AIOTSessionReleaseCommandTransferOCTETSTRING `ngap:"id:463,criticality:reject,presence:mandatory"`
	AdditionalIEs                     []ngapType.UnknownIE
	UnknownExtensions                 []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AIOTSessionReleaseCommand) Encode() ([]byte, error) {
	return ngapcodec.Encode("AIOTSessionReleaseCommand", m)
}

// DecodeAIOTSessionReleaseCommand decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAIOTSessionReleaseCommand(data []byte) (*AIOTSessionReleaseCommand, error) {
	var m AIOTSessionReleaseCommand
	err := ngapcodec.Decode("AIOTSessionReleaseCommand", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
