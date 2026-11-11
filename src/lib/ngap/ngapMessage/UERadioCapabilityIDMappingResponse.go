// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UERadioCapabilityIDMappingResponse represents ASN.1 UERadioCapabilityIDMappingResponse.
type UERadioCapabilityIDMappingResponse struct {
	UERadioCapabilityID       ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:mandatory"`
	UERadioCapability         ngapType.UERadioCapability   `ngap:"id:117,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UERadioCapabilityIDMappingResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("UERadioCapabilityIDMappingResponse", m)
}

// DecodeUERadioCapabilityIDMappingResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUERadioCapabilityIDMappingResponse(data []byte) (*UERadioCapabilityIDMappingResponse, error) {
	var m UERadioCapabilityIDMappingResponse
	err := ngapcodec.Decode("UERadioCapabilityIDMappingResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
