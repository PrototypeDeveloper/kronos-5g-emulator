// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AIOTSessionReleaseRequest represents ASN.1 AIOTSessionReleaseRequest.
type AIOTSessionReleaseRequest struct {
	AIOTFIdentifier                   ngapType.AIOTFIdentifier                              `ngap:"id:467,criticality:reject,presence:mandatory"`
	AIoTCorrelationIdentifier         ngapType.AIoTCorrelationIdentifier                    `ngap:"id:466,criticality:reject,presence:mandatory"`
	AIOTSessionReleaseRequestTransfer ngapType.AIOTSessionReleaseRequestTransferOCTETSTRING `ngap:"id:465,criticality:reject,presence:mandatory"`
	AdditionalIEs                     []ngapType.UnknownIE
	UnknownExtensions                 []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AIOTSessionReleaseRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("AIOTSessionReleaseRequest", m)
}

// DecodeAIOTSessionReleaseRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAIOTSessionReleaseRequest(data []byte) (*AIOTSessionReleaseRequest, error) {
	var m AIOTSessionReleaseRequest
	err := ngapcodec.Decode("AIOTSessionReleaseRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
