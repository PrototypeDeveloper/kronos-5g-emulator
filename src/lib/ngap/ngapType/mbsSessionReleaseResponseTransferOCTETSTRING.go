// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSSessionReleaseResponseTransferOCTETSTRING represents ASN.1 MBSSessionReleaseResponseTransfer-OCTET-STRING.
type MBSSessionReleaseResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSSessionReleaseResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBSSessionReleaseResponseTransfer-OCTET-STRING", v)
}

// DecodeMBSSessionReleaseResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSSessionReleaseResponseTransferOCTETSTRING(data []byte) (MBSSessionReleaseResponseTransferOCTETSTRING, int, error) {
	var v MBSSessionReleaseResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBSSessionReleaseResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSSessionReleaseResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
