// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverCommand represents ASN.1 HandoverCommand.
type HandoverCommand struct {
	AMFUENGAPID                             ngapType.AMFUENGAPID  `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                             ngapType.RANUENGAPID  `ngap:"id:85,criticality:reject,presence:mandatory"`
	HandoverType                            ngapType.HandoverType `ngap:"id:29,criticality:reject,presence:mandatory"`
	HasNASSecurityParametersFromNGRAN       bool
	NASSecurityParametersFromNGRAN          ngapType.NASSecurityParametersFromNGRAN `ngap:"id:39,criticality:reject,presence:conditional"`
	HasPDUSessionResourceHandoverList       bool
	PDUSessionResourceHandoverList          ngapType.PDUSessionResourceHandoverList `ngap:"id:59,criticality:ignore,presence:optional"`
	HasPDUSessionResourceToReleaseListHOCmd bool
	PDUSessionResourceToReleaseListHOCmd    ngapType.PDUSessionResourceToReleaseListHOCmd `ngap:"id:78,criticality:ignore,presence:optional"`
	TargetToSourceTransparentContainer      ngapType.TargetToSourceTransparentContainer   `ngap:"id:106,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics               bool
	CriticalityDiagnostics                  ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                           []ngapType.UnknownIE
	UnknownExtensions                       []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverCommand) Encode() ([]byte, error) { return ngapcodec.Encode("HandoverCommand", m) }

// DecodeHandoverCommand decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverCommand(data []byte) (*HandoverCommand, error) {
	var m HandoverCommand
	err := ngapcodec.Decode("HandoverCommand", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
