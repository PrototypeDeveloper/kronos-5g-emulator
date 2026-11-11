// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryReportTransferInventoryCompleteIndication represents ASN.1 InventoryReportTransferInventoryCompleteIndication.
const (
	InventoryReportTransferInventoryCompleteIndicationPresentTrue Enumerated = 0
)

type InventoryReportTransferInventoryCompleteIndication struct {
	Value Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}

// Encode encodes this parameter as a standalone APER value.
func (v InventoryReportTransferInventoryCompleteIndication) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("inline:InventoryReportTransferInventoryCompleteIndication", v)
}

// DecodeInventoryReportTransferInventoryCompleteIndication decodes a standalone APER value and reports consumed octets.
func DecodeInventoryReportTransferInventoryCompleteIndication(data []byte) (InventoryReportTransferInventoryCompleteIndication, int, error) {
	var v InventoryReportTransferInventoryCompleteIndication
	err := ngapcodec.DecodeParameter("inline:InventoryReportTransferInventoryCompleteIndication", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryReportTransferInventoryCompleteIndication{}, 0, err
	}
	return v, len(data), err
}
