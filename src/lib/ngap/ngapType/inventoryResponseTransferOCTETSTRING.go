// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryResponseTransferOCTETSTRING represents ASN.1 InventoryResponseTransfer-OCTET-STRING.
type InventoryResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryResponseTransfer-OCTET-STRING", v)
}

// DecodeInventoryResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryResponseTransferOCTETSTRING(data []byte) (InventoryResponseTransferOCTETSTRING, int, error) {
	var v InventoryResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("InventoryResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
