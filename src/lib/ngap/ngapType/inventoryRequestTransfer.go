// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryRequestTransfer represents ASN.1 InventoryRequestTransfer.
type InventoryRequestTransfer struct {
	ProtocolIEs       []InventoryRequestTransferIEsEntry
	UnknownExtensions []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryRequestTransfer", v)
}

// DecodeInventoryRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryRequestTransfer(data []byte) (InventoryRequestTransfer, int, error) {
	var v InventoryRequestTransfer
	err := ngapcodec.DecodeParameter("InventoryRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
