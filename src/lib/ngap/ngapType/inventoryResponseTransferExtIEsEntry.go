// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InventoryResponseTransferExtIEsEntryPresent uint32

const (
	InventoryResponseTransferExtIEsEntryPresentNothing InventoryResponseTransferExtIEsEntryPresent = iota
	InventoryResponseTransferExtIEsEntryPresentUnknown
)

type InventoryResponseTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      InventoryResponseTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryResponseTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryResponseTransfer-ExtIEsEntry", v)
}

// DecodeInventoryResponseTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryResponseTransferExtIEsEntry(data []byte) (InventoryResponseTransferExtIEsEntry, int, error) {
	var v InventoryResponseTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("InventoryResponseTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryResponseTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
