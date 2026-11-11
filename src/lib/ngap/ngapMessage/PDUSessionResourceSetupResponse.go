// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PDUSessionResourceSetupResponse represents ASN.1 PDUSessionResourceSetupResponse.
type PDUSessionResourceSetupResponse struct {
	AMFUENGAPID                                 ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                 ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceSetupListSURes         bool
	PDUSessionResourceSetupListSURes            ngapType.PDUSessionResourceSetupListSURes `ngap:"id:75,criticality:ignore,presence:optional"`
	HasPDUSessionResourceFailedToSetupListSURes bool
	PDUSessionResourceFailedToSetupListSURes    ngapType.PDUSessionResourceFailedToSetupListSURes `ngap:"id:58,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics                   bool
	CriticalityDiagnostics                      ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasUserLocationInformation                  bool
	UserLocationInformation                     ngapType.UserLocationInformation `ngap:"id:121,criticality:ignore,presence:optional"`
	AdditionalIEs                               []ngapType.UnknownIE
	UnknownExtensions                           []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PDUSessionResourceSetupResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("PDUSessionResourceSetupResponse", m)
}

// DecodePDUSessionResourceSetupResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePDUSessionResourceSetupResponse(data []byte) (*PDUSessionResourceSetupResponse, error) {
	var m PDUSessionResourceSetupResponse
	err := ngapcodec.Decode("PDUSessionResourceSetupResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
