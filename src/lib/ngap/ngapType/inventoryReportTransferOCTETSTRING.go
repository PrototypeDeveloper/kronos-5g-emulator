// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// InventoryReportTransferOCTETSTRING represents ASN.1 InventoryReportTransfer-OCTET-STRING.
type InventoryReportTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v InventoryReportTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("InventoryReportTransfer-OCTET-STRING", v)
}

// DecodeInventoryReportTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeInventoryReportTransferOCTETSTRING(data []byte) (InventoryReportTransferOCTETSTRING, int, error) {
	var v InventoryReportTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("InventoryReportTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return InventoryReportTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
