// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InventoryReportTransferExtIEsEntryPresent uint32

const (
	InventoryReportTransferExtIEsEntryPresentNothing InventoryReportTransferExtIEsEntryPresent = iota
	InventoryReportTransferExtIEsEntryPresentUnknown
)

type InventoryReportTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      InventoryReportTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryReportTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryReportTransfer-ExtIEsEntry", v)
}

// DecodeInventoryReportTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryReportTransferExtIEsEntry(data []byte) (InventoryReportTransferExtIEsEntry, int, error) {
	var v InventoryReportTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("InventoryReportTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryReportTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
