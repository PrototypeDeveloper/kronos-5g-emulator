// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MTCommunicationHandlingResponse represents ASN.1 MTCommunicationHandlingResponse.
type MTCommunicationHandlingResponse struct {
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MTCommunicationHandlingResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("MTCommunicationHandlingResponse", m)
}

// DecodeMTCommunicationHandlingResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMTCommunicationHandlingResponse(data []byte) (*MTCommunicationHandlingResponse, error) {
	var m MTCommunicationHandlingResponse
	err := ngapcodec.Decode("MTCommunicationHandlingResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
