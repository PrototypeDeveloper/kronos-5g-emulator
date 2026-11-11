// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTSessionReleaseCompleteTransferOCTETSTRING represents ASN.1 AIOTSessionReleaseCompleteTransfer-OCTET-STRING.
type AIOTSessionReleaseCompleteTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseCompleteTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseCompleteTransfer-OCTET-STRING", v)
}

// DecodeAIOTSessionReleaseCompleteTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseCompleteTransferOCTETSTRING(data []byte) (AIOTSessionReleaseCompleteTransferOCTETSTRING, int, error) {
	var v AIOTSessionReleaseCompleteTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseCompleteTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseCompleteTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
