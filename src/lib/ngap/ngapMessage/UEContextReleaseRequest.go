// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextReleaseRequest represents ASN.1 UEContextReleaseRequest.
type UEContextReleaseRequest struct {
	AMFUENGAPID                        ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID                        ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	HasPDUSessionResourceListCxtRelReq bool
	PDUSessionResourceListCxtRelReq    ngapType.PDUSessionResourceListCxtRelReq `ngap:"id:133,criticality:reject,presence:optional"`
	Cause                              ngapType.Cause                           `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasGWContextReleaseIndication      bool
	GWContextReleaseIndication         ngapType.GWContextReleaseIndication `ngap:"id:452,criticality:reject,presence:optional"`
	AdditionalIEs                      []ngapType.UnknownIE
	UnknownExtensions                  []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextReleaseRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextReleaseRequest", m)
}

// DecodeUEContextReleaseRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextReleaseRequest(data []byte) (*UEContextReleaseRequest, error) {
	var m UEContextReleaseRequest
	err := ngapcodec.Decode("UEContextReleaseRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
