// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// InitialContextSetupResponse represents ASN.1 InitialContextSetupResponse.
type InitialContextSetupResponse struct {
	AMFUENGAPID                                  ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                  ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceSetupListCxtRes         bool
	PDUSessionResourceSetupListCxtRes            ngapType.PDUSessionResourceSetupListCxtRes `ngap:"id:72,criticality:ignore,presence:optional"`
	HasPDUSessionResourceFailedToSetupListCxtRes bool
	PDUSessionResourceFailedToSetupListCxtRes    ngapType.PDUSessionResourceFailedToSetupListCxtRes `ngap:"id:55,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics                    bool
	CriticalityDiagnostics                       ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                                []ngapType.UnknownIE
	UnknownExtensions                            []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *InitialContextSetupResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("InitialContextSetupResponse", m)
}

// DecodeInitialContextSetupResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeInitialContextSetupResponse(data []byte) (*InitialContextSetupResponse, error) {
	var m InitialContextSetupResponse
	err := ngapcodec.Decode("InitialContextSetupResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
