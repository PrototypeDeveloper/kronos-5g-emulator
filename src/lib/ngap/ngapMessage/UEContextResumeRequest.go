// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextResumeRequest represents ASN.1 UEContextResumeRequest.
type UEContextResumeRequest struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID           `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID           `ngap:"id:85,criticality:reject,presence:mandatory"`
	RRCResumeCause                                ngapType.RRCEstablishmentCause `ngap:"id:237,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceResumeListRESReq         bool
	PDUSessionResourceResumeListRESReq            ngapType.PDUSessionResourceResumeListRESReq `ngap:"id:232,criticality:reject,presence:optional"`
	HasPDUSessionResourceFailedToResumeListRESReq bool
	PDUSessionResourceFailedToResumeListRESReq    ngapType.PDUSessionResourceFailedToResumeListRESReq `ngap:"id:229,criticality:reject,presence:optional"`
	HasSuspendRequestIndication                   bool
	SuspendRequestIndication                      ngapType.SuspendRequestIndication `ngap:"id:235,criticality:ignore,presence:optional"`
	HasInfoOnRecommendedCellsAndRANNodesForPaging bool
	InfoOnRecommendedCellsAndRANNodesForPaging    ngapType.InfoOnRecommendedCellsAndRANNodesForPaging `ngap:"id:32,criticality:ignore,presence:optional"`
	HasPagingAssisDataforCEcapabUE                bool
	PagingAssisDataforCEcapabUE                   ngapType.PagingAssisDataforCEcapabUE `ngap:"id:207,criticality:ignore,presence:optional"`
	HasUserLocationInformation                    bool
	UserLocationInformation                       ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextResumeRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextResumeRequest", m)
}

// DecodeUEContextResumeRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextResumeRequest(data []byte) (*UEContextResumeRequest, error) {
	var m UEContextResumeRequest
	err := ngapcodec.Decode("UEContextResumeRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
