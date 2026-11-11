// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextReleaseCommand represents ASN.1 UEContextReleaseCommand.
type UEContextReleaseCommand struct {
	UENGAPIDs         ngapType.UENGAPIDs `ngap:"id:114,criticality:reject,presence:mandatory"`
	Cause             ngapType.Cause     `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextReleaseCommand) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextReleaseCommand", m)
}

// DecodeUEContextReleaseCommand decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextReleaseCommand(data []byte) (*UEContextReleaseCommand, error) {
	var m UEContextReleaseCommand
	err := ngapcodec.Decode("UEContextReleaseCommand", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
