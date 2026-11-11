// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UERadioCapabilityCheckResponse represents ASN.1 UERadioCapabilityCheckResponse.
type UERadioCapabilityCheckResponse struct {
	AMFUENGAPID               ngapType.AMFUENGAPID              `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID              `ngap:"id:85,criticality:ignore,presence:mandatory"`
	IMSVoiceSupportIndicator  ngapType.IMSVoiceSupportIndicator `ngap:"id:30,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UERadioCapabilityCheckResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("UERadioCapabilityCheckResponse", m)
}

// DecodeUERadioCapabilityCheckResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUERadioCapabilityCheckResponse(data []byte) (*UERadioCapabilityCheckResponse, error) {
	var m UERadioCapabilityCheckResponse
	err := ngapcodec.Decode("UERadioCapabilityCheckResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
