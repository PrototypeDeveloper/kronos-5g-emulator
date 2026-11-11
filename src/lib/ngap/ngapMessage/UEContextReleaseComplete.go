// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextReleaseComplete represents ASN.1 UEContextReleaseComplete.
type UEContextReleaseComplete struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasUserLocationInformation                    bool
	UserLocationInformation                       ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	HasInfoOnRecommendedCellsAndRANNodesForPaging bool
	InfoOnRecommendedCellsAndRANNodesForPaging    ngapType.InfoOnRecommendedCellsAndRANNodesForPaging `ngap:"id:32,criticality:ignore,presence:optional"`
	HasPDUSessionResourceListCxtRelCpl            bool
	PDUSessionResourceListCxtRelCpl               ngapType.PDUSessionResourceListCxtRelCpl `ngap:"id:60,criticality:reject,presence:optional"`
	HasCriticalityDiagnostics                     bool
	CriticalityDiagnostics                        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasPagingAssisDataforCEcapabUE                bool
	PagingAssisDataforCEcapabUE                   ngapType.PagingAssisDataforCEcapabUE `ngap:"id:207,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextReleaseComplete) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextReleaseComplete", m)
}

// DecodeUEContextReleaseComplete decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextReleaseComplete(data []byte) (*UEContextReleaseComplete, error) {
	var m UEContextReleaseComplete
	err := ngapcodec.Decode("UEContextReleaseComplete", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
