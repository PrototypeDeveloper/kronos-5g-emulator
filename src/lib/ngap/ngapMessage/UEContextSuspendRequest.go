// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextSuspendRequest represents ASN.1 UEContextSuspendRequest.
type UEContextSuspendRequest struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasInfoOnRecommendedCellsAndRANNodesForPaging bool
	InfoOnRecommendedCellsAndRANNodesForPaging    ngapType.InfoOnRecommendedCellsAndRANNodesForPaging `ngap:"id:32,criticality:ignore,presence:optional"`
	HasPagingAssisDataforCEcapabUE                bool
	PagingAssisDataforCEcapabUE                   ngapType.PagingAssisDataforCEcapabUE `ngap:"id:207,criticality:ignore,presence:optional"`
	HasPDUSessionResourceSuspendListSUSReq        bool
	PDUSessionResourceSuspendListSUSReq           ngapType.PDUSessionResourceSuspendListSUSReq `ngap:"id:231,criticality:reject,presence:optional"`
	HasUserLocationInformation                    bool
	UserLocationInformation                       ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextSuspendRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextSuspendRequest", m)
}

// DecodeUEContextSuspendRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextSuspendRequest(data []byte) (*UEContextSuspendRequest, error) {
	var m UEContextSuspendRequest
	err := ngapcodec.Decode("UEContextSuspendRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
