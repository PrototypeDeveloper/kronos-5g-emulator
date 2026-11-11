// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapType

import "libngap/ngapcodec"

// MBSDistributionSetupResponseTransfer represents ASN.1 MBS-DistributionSetupResponseTransfer.
type MBSDistributionSetupResponseTransfer struct {
	MBSSessionID                        MBSSessionID
	HasMBSAreaSessionID                 bool
	MBSAreaSessionID                    MBSAreaSessionID
	HasSharedNGUMulticastTNLInformation bool
	SharedNGUMulticastTNLInformation    SharedNGUMulticastTNLInformation
	MBSQoSFlowsToBeSetupList            MBSQoSFlowsToBeSetupList
	MBSSessionStatus                    MBSSessionStatus
	HasMBSServiceArea                   bool
	MBSServiceArea                      MBSServiceArea
	HasIEExtensions                     bool
	IEExtensions                        []MBSDistributionSetupResponseTransferExtIEsEntry
	UnknownExtensions                   []UnknownExtension
}

// Encode encodes this parameter as a standalone APER value, without an IE header.
func (v MBSDistributionSetupResponseTransfer) Encode() ([]byte, error) {
	return ngapcodec.EncodeParameter("MBS-DistributionSetupResponseTransfer", v)
}

// DecodeMBSDistributionSetupResponseTransfer decodes a standalone APER value and reports the consumed octets.
func DecodeMBSDistributionSetupResponseTransfer(data []byte) (MBSDistributionSetupResponseTransfer, int, error) {
	var v MBSDistributionSetupResponseTransfer
	err := ngapcodec.DecodeParameter("MBS-DistributionSetupResponseTransfer", data, &v)
	if ngapcodec.IsFatal(err) {
		return MBSDistributionSetupResponseTransfer{}, 0, err
	}
	return v, len(data), err
}
