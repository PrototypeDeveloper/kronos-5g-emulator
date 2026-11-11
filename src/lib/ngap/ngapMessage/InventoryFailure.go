// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InventoryFailure represents ASN.1 InventoryFailure.
type InventoryFailure struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                     `ngap:"id:467,criticality:ignore,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier           `ngap:"id:466,criticality:ignore,presence:mandatory"`
	InventoryFailureTransfer  ngapType.InventoryFailureTransferOCTETSTRING `ngap:"id:458,criticality:ignore,presence:mandatory"`
	Cause                     ngapType.Cause                               `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InventoryFailure) Encode() ([]byte, error) { return ngapcodec.Encode("InventoryFailure", m) }

// DecodeInventoryFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInventoryFailure(data []byte) (*InventoryFailure, error) {
	var m InventoryFailure
	err := ngapcodec.Decode("InventoryFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
