// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DistributionSetupFailure represents ASN.1 DistributionSetupFailure.
type DistributionSetupFailure struct {
	MBSSessionID                             ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID                      bool
	MBSAreaSessionID                         ngapType.MBSAreaSessionID                                    `ngap:"id:295,criticality:reject,presence:optional"`
	MBSDistributionSetupUnsuccessfulTransfer ngapType.MBSDistributionSetupUnsuccessfulTransferOCTETSTRING `ngap:"id:303,criticality:ignore,presence:mandatory"`
	Cause                                    ngapType.Cause                                               `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics                bool
	CriticalityDiagnostics                   ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                            []ngapType.UnknownIE
	UnknownExtensions                        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DistributionSetupFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("DistributionSetupFailure", m)
}

// DecodeDistributionSetupFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDistributionSetupFailure(data []byte) (*DistributionSetupFailure, error) {
	var m DistributionSetupFailure
	err := ngapcodec.Decode("DistributionSetupFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
