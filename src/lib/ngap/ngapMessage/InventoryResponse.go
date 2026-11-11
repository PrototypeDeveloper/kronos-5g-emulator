// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InventoryResponse represents ASN.1 InventoryResponse.
type InventoryResponse struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                      `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier            `ngap:"id:466,criticality:reject,presence:mandatory"`
	InventoryResponseTransfer ngapType.InventoryResponseTransferOCTETSTRING `ngap:"id:457,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InventoryResponse) Encode() ([]byte, error) { return ngapcodec.Encode("InventoryResponse", m) }

// DecodeInventoryResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInventoryResponse(data []byte) (*InventoryResponse, error) {
	var m InventoryResponse
	err := ngapcodec.Decode("InventoryResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
