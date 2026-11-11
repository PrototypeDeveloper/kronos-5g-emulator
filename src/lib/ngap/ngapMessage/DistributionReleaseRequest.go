// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DistributionReleaseRequest represents ASN.1 DistributionReleaseRequest.
type DistributionReleaseRequest struct {
	MBSSessionID                          ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID                   bool
	MBSAreaSessionID                      ngapType.MBSAreaSessionID                                 `ngap:"id:295,criticality:reject,presence:optional"`
	MBSDistributionReleaseRequestTransfer ngapType.MBSDistributionReleaseRequestTransferOCTETSTRING `ngap:"id:300,criticality:reject,presence:mandatory"`
	Cause                                 ngapType.Cause                                            `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs                         []ngapType.UnknownIE
	UnknownExtensions                     []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DistributionReleaseRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("DistributionReleaseRequest", m)
}

// DecodeDistributionReleaseRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDistributionReleaseRequest(data []byte) (*DistributionReleaseRequest, error) {
	var m DistributionReleaseRequest
	err := ngapcodec.Decode("DistributionReleaseRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
