// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InventoryRequest represents ASN.1 InventoryRequest.
type InventoryRequest struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                     `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier           `ngap:"id:466,criticality:reject,presence:mandatory"`
	InventoryRequestTransfer  ngapType.InventoryRequestTransferOCTETSTRING `ngap:"id:456,criticality:reject,presence:mandatory"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InventoryRequest) Encode() ([]byte, error) { return ngapcodec.Encode("InventoryRequest", m) }

// DecodeInventoryRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInventoryRequest(data []byte) (*InventoryRequest, error) {
	var m InventoryRequest
	err := ngapcodec.Decode("InventoryRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
