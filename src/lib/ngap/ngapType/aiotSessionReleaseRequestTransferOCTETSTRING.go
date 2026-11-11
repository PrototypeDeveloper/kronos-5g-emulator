// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// AIOTSessionReleaseRequestTransferOCTETSTRING represents ASN.1 AIOTSessionReleaseRequestTransfer-OCTET-STRING.
type AIOTSessionReleaseRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v AIOTSessionReleaseRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("AIOTSessionReleaseRequestTransfer-OCTET-STRING", v)
}

// DecodeAIOTSessionReleaseRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeAIOTSessionReleaseRequestTransferOCTETSTRING(data []byte) (AIOTSessionReleaseRequestTransferOCTETSTRING, int, error) {
	var v AIOTSessionReleaseRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("AIOTSessionReleaseRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return AIOTSessionReleaseRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
