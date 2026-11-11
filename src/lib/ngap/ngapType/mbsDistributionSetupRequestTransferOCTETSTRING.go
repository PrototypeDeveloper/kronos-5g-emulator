// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionSetupRequestTransferOCTETSTRING represents ASN.1 MBS-DistributionSetupRequestTransfer-OCTET-STRING.
type MBSDistributionSetupRequestTransferOCTETSTRING struct {
	Value []byte
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupRequestTransferOCTETSTRING) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupRequestTransfer-OCTET-STRING", v)
}

// DecodeMBSDistributionSetupRequestTransferOCTETSTRING decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupRequestTransferOCTETSTRING(data []byte) (MBSDistributionSetupRequestTransferOCTETSTRING, int, error) {
	var v MBSDistributionSetupRequestTransferOCTETSTRING
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupRequestTransfer-OCTET-STRING", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupRequestTransferOCTETSTRING{}, 0, err
	}
	return v, len(data), err
}
