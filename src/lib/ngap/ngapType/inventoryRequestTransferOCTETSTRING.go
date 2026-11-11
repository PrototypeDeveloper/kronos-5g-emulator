// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryRequestTransferOCTETSTRING represents ASN.1 InventoryRequestTransfer-OCTET-STRING.
type InventoryRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryRequestTransfer-OCTET-STRING", v)
}

// DecodeInventoryRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryRequestTransferOCTETSTRING(data []byte) (InventoryRequestTransferOCTETSTRING, int, error) {
	var v InventoryRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("InventoryRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
