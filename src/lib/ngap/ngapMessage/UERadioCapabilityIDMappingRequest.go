// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UERadioCapabilityIDMappingRequest represents ASN.1 UERadioCapabilityIDMappingRequest.
type UERadioCapabilityIDMappingRequest struct {
	UERadioCapabilityID ngapType.UERadioCapabilityID `ngap:"id:264,criticality:reject,presence:mandatory"`
	AdditionalIEs       []ngapType.UnknownIE
	UnknownExtensions   []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UERadioCapabilityIDMappingRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UERadioCapabilityIDMappingRequest", m)
}

// DecodeUERadioCapabilityIDMappingRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUERadioCapabilityIDMappingRequest(data []byte) (*UERadioCapabilityIDMappingRequest, error) {
	var m UERadioCapabilityIDMappingRequest
	err := ngapcodec.Decode("UERadioCapabilityIDMappingRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
