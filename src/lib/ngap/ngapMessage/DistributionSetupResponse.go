// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DistributionSetupResponse represents ASN.1 DistributionSetupResponse.
type DistributionSetupResponse struct {
	MBSSessionID                         ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID                  bool
	MBSAreaSessionID                     ngapType.MBSAreaSessionID                                `ngap:"id:295,criticality:reject,presence:optional"`
	MBSDistributionSetupResponseTransfer ngapType.MBSDistributionSetupResponseTransferOCTETSTRING `ngap:"id:302,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics            bool
	CriticalityDiagnostics               ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                        []ngapType.UnknownIE
	UnknownExtensions                    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DistributionSetupResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("DistributionSetupResponse", m)
}

// DecodeDistributionSetupResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDistributionSetupResponse(data []byte) (*DistributionSetupResponse, error) {
	var m DistributionSetupResponse
	err := ngapcodec.Decode("DistributionSetupResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
