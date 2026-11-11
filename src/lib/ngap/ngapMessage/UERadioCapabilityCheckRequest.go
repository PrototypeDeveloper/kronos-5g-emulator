// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UERadioCapabilityCheckRequest represents ASN.1 UERadioCapabilityCheckRequest.
type UERadioCapabilityCheckRequest struct {
	AMFUENGAPID            ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID            ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasUERadioCapability   bool
	UERadioCapability      ngapType.UERadioCapability `ngap:"id:117,criticality:ignore,presence:optional"`
	HasUERadioCapabilityID bool
	UERadioCapabilityID    ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:optional"`
	AdditionalIEs          []ngapType.UnknownIE
	UnknownExtensions      []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UERadioCapabilityCheckRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UERadioCapabilityCheckRequest", m)
}

// DecodeUERadioCapabilityCheckRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUERadioCapabilityCheckRequest(data []byte) (*UERadioCapabilityCheckRequest, error) {
	var m UERadioCapabilityCheckRequest
	err := ngapcodec.Decode("UERadioCapabilityCheckRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
