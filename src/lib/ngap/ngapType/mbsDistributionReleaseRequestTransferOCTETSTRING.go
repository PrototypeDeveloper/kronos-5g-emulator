// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionReleaseRequestTransferOCTETSTRING represents ASN.1 MBS-DistributionReleaseRequestTransfer-OCTET-STRING.
type MBSDistributionReleaseRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionReleaseRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionReleaseRequestTransfer-OCTET-STRING", v)
}

// DecodeMBSDistributionReleaseRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionReleaseRequestTransferOCTETSTRING(data []byte) (MBSDistributionReleaseRequestTransferOCTETSTRING, int, error) {
	var v MBSDistributionReleaseRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBS-DistributionReleaseRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionReleaseRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
