// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InventoryReport represents ASN.1 InventoryReport.
type InventoryReport struct {
	AIOTFIdentifier           ngapType.AIOTFIdentifier                    `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier ngapType.AIoTCorrelationIdentifier          `ngap:"id:466,criticality:reject,presence:mandatory"`
	InventoryReportTransfer   ngapType.InventoryReportTransferOCTETSTRING `ngap:"id:459,criticality:reject,presence:mandatory"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InventoryReport) Encode() ([]byte, error) { return ngapcodec.Encode("InventoryReport", m) }

// DecodeInventoryReport decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInventoryReport(data []byte) (*InventoryReport, error) {
	var m InventoryReport
	err := ngapcodec.Decode("InventoryReport", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
