// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextResumeFailure represents ASN.1 UEContextResumeFailure.
type UEContextResumeFailure struct {
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	Cause                     ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextResumeFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextResumeFailure", m)
}

// DecodeUEContextResumeFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextResumeFailure(data []byte) (*UEContextResumeFailure, error) {
	var m UEContextResumeFailure
	err := ngapcodec.Decode("UEContextResumeFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
