// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionSetupOrModFailureTransferOCTETSTRING represents ASN.1 MBSSessionSetupOrModFailureTransfer-OCTET-STRING.
type MBSSessionSetupOrModFailureTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionSetupOrModFailureTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionSetupOrModFailureTransfer-OCTET-STRING", v)
}

// DecodeMBSSessionSetupOrModFailureTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionSetupOrModFailureTransferOCTETSTRING(data []byte) (MBSSessionSetupOrModFailureTransferOCTETSTRING, int, error) {
	var v MBSSessionSetupOrModFailureTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBSSessionSetupOrModFailureTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionSetupOrModFailureTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
