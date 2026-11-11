// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionSetupResponseTransferOCTETSTRING represents ASN.1 MBS-DistributionSetupResponseTransfer-OCTET-STRING.
type MBSDistributionSetupResponseTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupResponseTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupResponseTransfer-OCTET-STRING", v)
}

// DecodeMBSDistributionSetupResponseTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupResponseTransferOCTETSTRING(data []byte) (MBSDistributionSetupResponseTransferOCTETSTRING, int, error) {
	var v MBSDistributionSetupResponseTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupResponseTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupResponseTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
