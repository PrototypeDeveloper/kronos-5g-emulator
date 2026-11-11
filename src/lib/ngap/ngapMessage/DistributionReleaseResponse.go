// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DistributionReleaseResponse represents ASN.1 DistributionReleaseResponse.
type DistributionReleaseResponse struct {
	MBSSessionID              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID       bool
	MBSAreaSessionID          ngapType.MBSAreaSessionID `ngap:"id:295,criticality:reject,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DistributionReleaseResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("DistributionReleaseResponse", m)
}

// DecodeDistributionReleaseResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDistributionReleaseResponse(data []byte) (*DistributionReleaseResponse, error) {
	var m DistributionReleaseResponse
	err := ngapcodec.Decode("DistributionReleaseResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
