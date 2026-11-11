// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceSetupRequest represents ASN.1 PDUSessionResourceSetupRequest.
type PDUSessionResourceSetupRequest struct {
	AMFUENGAPID                      ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                      ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasRANPagingPriority             bool
	RANPagingPriority                ngapType.RANPagingPriority `ngap:"id:83,criticality:ignore,presence:optional"`
	HasNASPDU                        bool
	NASPDU                           ngapType.NASPDU                           `ngap:"id:38,criticality:reject,presence:optional"`
	PDUSessionResourceSetupListSUReq ngapType.PDUSessionResourceSetupListSUReq `ngap:"id:74,criticality:reject,presence:mandatory"`
	HasUEAggregateMaximumBitRate     bool
	UEAggregateMaximumBitRate        ngapType.UEAggregateMaximumBitRate `ngap:"id:110,criticality:ignore,presence:optional"`
	HasUESliceMaximumBitRateList     bool
	UESliceMaximumBitRateList        ngapType.UESliceMaximumBitRateList `ngap:"id:335,criticality:ignore,presence:optional"`
	AdditionalIEs                    []ngapType.UnknownIE
	UnknownExtensions                []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceSetupRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceSetupRequest", m)
}

// DecodePDUSessionResourceSetupRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceSetupRequest(data []byte) (*PDUSessionResourceSetupRequest, error) {
	var m PDUSessionResourceSetupRequest
	err := ngapcodec.Decode("PDUSessionResourceSetupRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
