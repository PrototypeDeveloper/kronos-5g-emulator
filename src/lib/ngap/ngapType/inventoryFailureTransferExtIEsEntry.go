// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

type InventoryFailureTransferExtIEsEntryPresent uint32

const (
	InventoryFailureTransferExtIEsEntryPresentNothing InventoryFailureTransferExtIEsEntryPresent = iota
	InventoryFailureTransferExtIEsEntryPresentUnknown
)

type InventoryFailureTransferExtIEsEntry struct {
	ID           ProtocolExtensionID
	Criticality  Criticality
	Present      InventoryFailureTransferExtIEsEntryPresent
	UnknownValue []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryFailureTransferExtIEsEntry) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryFailureTransfer-ExtIEsEntry", v)
}

// DecodeInventoryFailureTransferExtIEsEntry decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryFailureTransferExtIEsEntry(data []byte) (InventoryFailureTransferExtIEsEntry, int, error) {
	var v InventoryFailureTransferExtIEsEntry
	err := ngapcodec.DecodeParameter("InventoryFailureTransfer-ExtIEsEntry", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryFailureTransferExtIEsEntry{}, 0, err
	}
	return v, len(data), err
}
