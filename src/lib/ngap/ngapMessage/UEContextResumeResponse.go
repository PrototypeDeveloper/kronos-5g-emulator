// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UEContextResumeResponse represents ASN.1 UEContextResumeResponse.
type UEContextResumeResponse struct {
	AMFUENGAPID                                   ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                   ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceResumeListRESRes         bool
	PDUSessionResourceResumeListRESRes            ngapType.PDUSessionResourceResumeListRESRes `ngap:"id:233,criticality:reject,presence:optional"`
	HasPDUSessionResourceFailedToResumeListRESRes bool
	PDUSessionResourceFailedToResumeListRESRes    ngapType.PDUSessionResourceFailedToResumeListRESRes `ngap:"id:230,criticality:reject,presence:optional"`
	HasSecurityContext                            bool
	SecurityContext                               ngapType.SecurityContext `ngap:"id:93,criticality:reject,presence:optional"`
	HasSuspendResponseIndication                  bool
	SuspendResponseIndication                     ngapType.SuspendResponseIndication `ngap:"id:236,criticality:ignore,presence:optional"`
	HasExtendedConnectedTime                      bool
	ExtendedConnectedTime                         ngapType.ExtendedConnectedTime `ngap:"id:206,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics                     bool
	CriticalityDiagnostics                        ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                                 []ngapType.UnknownIE
	UnknownExtensions                             []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UEContextResumeResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("UEContextResumeResponse", m)
}

// DecodeUEContextResumeResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUEContextResumeResponse(data []byte) (*UEContextResumeResponse, error) {
	var m UEContextResumeResponse
	err := ngapcodec.Decode("UEContextResumeResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
