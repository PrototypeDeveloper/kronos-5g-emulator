// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGRemovalResponse represents ASN.1 NGRemovalResponse.
type NGRemovalResponse struct {
	AMFName                   ngapType.AMFName `ngap:"id:1,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasExtendedAMFName        bool
	ExtendedAMFName           ngapType.ExtendedAMFName `ngap:"id:274,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGRemovalResponse) Encode() ([]byte, error) { return ngapcodec.Encode("NGRemovalResponse", m) }

// DecodeNGRemovalResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGRemovalResponse(data []byte) (*NGRemovalResponse, error) {
	var m NGRemovalResponse
	err := ngapcodec.Decode("NGRemovalResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
