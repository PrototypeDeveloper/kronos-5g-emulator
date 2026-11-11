// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryReportTransfer represents ASN.1 InventoryReportTransfer.
type InventoryReportTransfer struct {
	CorrelationIdentifier          AIoTCorrelationIdentifier
	GlobalgNBID                    GlobalGNBID
	HasReaderReportList            bool
	ReaderReportList               AIoTReaderReportList
	HasInventoryCompleteIndication bool
	InventoryCompleteIndication    InventoryReportTransferInventoryCompleteIndication `aper:"valueExt,valueLB:0,valueUB:0"`
	HasIEExtensions                bool
	IEExtensions                   []InventoryReportTransferExtIEsEntry
	UnknownExtensions              []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryReportTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryReportTransfer", v)
}

// DecodeInventoryReportTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryReportTransfer(data []byte) (InventoryReportTransfer, int, error) {
	var v InventoryReportTransfer
	err := ngapcodec.DecodeParameter("InventoryReportTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryReportTransfer{}, 0, err
	}
	return v, len(data), err
}
