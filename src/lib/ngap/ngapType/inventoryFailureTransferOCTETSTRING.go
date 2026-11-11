// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryFailureTransferOCTETSTRING represents ASN.1 InventoryFailureTransfer-OCTET-STRING.
type InventoryFailureTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryFailureTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryFailureTransfer-OCTET-STRING", v)
}

// DecodeInventoryFailureTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryFailureTransferOCTETSTRING(data []byte) (InventoryFailureTransferOCTETSTRING, int, error) {
	var v InventoryFailureTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("InventoryFailureTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryFailureTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
