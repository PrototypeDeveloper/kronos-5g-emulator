// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionActivationResponse represents ASN.1 MulticastSessionActivationResponse.
type MulticastSessionActivationResponse struct {
	MBSSessionID              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionActivationResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionActivationResponse", m)
}

// DecodeMulticastSessionActivationResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionActivationResponse(data []byte) (*MulticastSessionActivationResponse, error) {
	var m MulticastSessionActivationResponse
	err := ngapcodec.Decode("MulticastSessionActivationResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
