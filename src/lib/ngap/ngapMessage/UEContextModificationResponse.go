// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextModificationResponse represents ASN.1 UEContextModificationResponse.
type UEContextModificationResponse struct {
	AMFUENGAPID                ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasRRCState                bool
	RRCState                   ngapType.RRCState `ngap:"id:92,criticality:ignore,presence:optional"`
	HasUserLocationInformation bool
	UserLocationInformation    ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics  bool
	CriticalityDiagnostics     ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs              []ngapType.UnknownIE
	UnknownExtensions          []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextModificationResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextModificationResponse", m)
}

// DecodeUEContextModificationResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextModificationResponse(data []byte) (*UEContextModificationResponse, error) {
	var m UEContextModificationResponse
	err := ngapcodec.Decode("UEContextModificationResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
