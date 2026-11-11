// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionSetupRequestTransfer represents ASN.1 MBS-DistributionSetupRequestTransfer.
type MBSDistributionSetupRequestTransfer struct {
	MBSSessionID                      MBSSessionID
	HasMBSAreaSessionID               bool
	MBSAreaSessionID                  MBSAreaSessionID
	HasSharedNGUUnicastTNLInformation bool
	SharedNGUUnicastTNLInformation    UPTransportLayerInformation
	HasIEExtensions                   bool
	IEExtensions                      []MBSDistributionSetupRequestTransferExtIEsEntry
	UnknownExtensions                 []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupRequestTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupRequestTransfer", v)
}

// DecodeMBSDistributionSetupRequestTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupRequestTransfer(data []byte) (MBSDistributionSetupRequestTransfer, int, error) {
	var v MBSDistributionSetupRequestTransfer
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupRequestTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupRequestTransfer{}, 0, err
	}
	return v, len(data), err
}
