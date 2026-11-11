// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTSessionReleaseCommandTransferOCTETSTRING represents ASN.1 AIOTSessionReleaseCommandTransfer-OCTET-STRING.
type AIOTSessionReleaseCommandTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseCommandTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseCommandTransfer-OCTET-STRING", v)
}

// DecodeAIOTSessionReleaseCommandTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseCommandTransferOCTETSTRING(data []byte) (AIOTSessionReleaseCommandTransferOCTETSTRING, int, error) {
	var v AIOTSessionReleaseCommandTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseCommandTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseCommandTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
