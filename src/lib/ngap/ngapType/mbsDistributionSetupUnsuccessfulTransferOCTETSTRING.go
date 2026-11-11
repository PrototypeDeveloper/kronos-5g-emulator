// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionSetupUnsuccessfulTransferOCTETSTRING represents ASN.1 MBS-DistributionSetupUnsuccessfulTransfer-OCTET-STRING.
type MBSDistributionSetupUnsuccessfulTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupUnsuccessfulTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupUnsuccessfulTransfer-OCTET-STRING", v)
}

// DecodeMBSDistributionSetupUnsuccessfulTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupUnsuccessfulTransferOCTETSTRING(data []byte) (MBSDistributionSetupUnsuccessfulTransferOCTETSTRING, int, error) {
	var v MBSDistributionSetupUnsuccessfulTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupUnsuccessfulTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupUnsuccessfulTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
