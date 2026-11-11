// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupOrModResponseTransferOCTETSTRING represents ASN.1 MBSSessionSetupOrModResponseTransfer-OCTET-STRING.
type MBSSessionSetupOrModResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModResponseTransfer-OCTET-STRING", v)
}

// DecodeMBSSessionSetupOrModResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModResponseTransferOCTETSTRING(data []byte) (MBSSessionSetupOrModResponseTransferOCTETSTRING, int, error) {
	var v MBSSessionSetupOrModResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
