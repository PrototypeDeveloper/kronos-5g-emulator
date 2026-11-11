// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DistributionSetupRequest represents ASN.1 DistributionSetupRequest.
type DistributionSetupRequest struct {
	MBSSessionID                        ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID                 bool
	MBSAreaSessionID                    ngapType.MBSAreaSessionID                               `ngap:"id:295,criticality:reject,presence:optional"`
	MBSDistributionSetupRequestTransfer ngapType.MBSDistributionSetupRequestTransferOCTETSTRING `ngap:"id:301,criticality:reject,presence:mandatory"`
	AdditionalIEs                       []ngapType.UnknownIE
	UnknownExtensions                   []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DistributionSetupRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("DistributionSetupRequest", m)
}

// DecodeDistributionSetupRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDistributionSetupRequest(data []byte) (*DistributionSetupRequest, error) {
	var m DistributionSetupRequest
	err := ngapcodec.Decode("DistributionSetupRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
