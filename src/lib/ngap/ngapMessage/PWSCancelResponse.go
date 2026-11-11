// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PWSCancelResponse represents ASN.1 PWSCancelResponse.
type PWSCancelResponse struct {
	MessageIdentifier             ngapType.MessageIdentifier `ngap:"id:35,criticality:reject,presence:mandatory"`
	SerialNumber                  ngapType.SerialNumber      `ngap:"id:95,criticality:reject,presence:mandatory"`
	HasBroadcastCancelledAreaList bool
	BroadcastCancelledAreaList    ngapType.BroadcastCancelledAreaList `ngap:"id:12,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics     bool
	CriticalityDiagnostics        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                 []ngapType.UnknownIE
	UnknownExtensions             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PWSCancelResponse) Encode() ([]byte, error) { return ngapcodec.Encode("PWSCancelResponse", m) }

// DecodePWSCancelResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePWSCancelResponse(data []byte) (*PWSCancelResponse, error) {
	var m PWSCancelResponse
	err := ngapcodec.Decode("PWSCancelResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
