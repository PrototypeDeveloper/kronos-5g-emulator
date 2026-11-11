// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextSuspendFailure represents ASN.1 UEContextSuspendFailure.
type UEContextSuspendFailure struct {
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	Cause                     ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextSuspendFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextSuspendFailure", m)
}

// DecodeUEContextSuspendFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextSuspendFailure(data []byte) (*UEContextSuspendFailure, error) {
	var m UEContextSuspendFailure
	err := ngapcodec.Decode("UEContextSuspendFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
