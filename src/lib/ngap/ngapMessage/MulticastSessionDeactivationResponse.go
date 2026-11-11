// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionDeactivationResponse represents ASN.1 MulticastSessionDeactivationResponse.
type MulticastSessionDeactivationResponse struct {
	MBSSessionID              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionDeactivationResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionDeactivationResponse", m)
}

// DecodeMulticastSessionDeactivationResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionDeactivationResponse(data []byte) (*MulticastSessionDeactivationResponse, error) {
	var m MulticastSessionDeactivationResponse
	err := ngapcodec.Decode("MulticastSessionDeactivationResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
