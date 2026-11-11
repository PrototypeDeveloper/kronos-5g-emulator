// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextSuspendResponse represents ASN.1 UEContextSuspendResponse.
type UEContextSuspendResponse struct {
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasSecurityContext        bool
	SecurityContext           ngapType.SecurityContext `ngap:"id:93,criticality:reject,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextSuspendResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextSuspendResponse", m)
}

// DecodeUEContextSuspendResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextSuspendResponse(data []byte) (*UEContextSuspendResponse, error) {
	var m UEContextSuspendResponse
	err := ngapcodec.Decode("UEContextSuspendResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
