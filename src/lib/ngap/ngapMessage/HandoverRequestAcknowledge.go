// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverRequestAcknowledge represents ASN.1 HandoverRequestAcknowledge.
type HandoverRequestAcknowledge struct {
	AMFUENGAPID                                 ngapType.AMFUENGAPID                    `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                 ngapType.RANUENGAPID                    `ngap:"id:85,criticality:ignore,presence:mandatory"`
	PDUSessionResourceAdmittedList              ngapType.PDUSessionResourceAdmittedList `ngap:"id:53,criticality:ignore,presence:mandatory"`
	HasPDUSessionResourceFailedToSetupListHOAck bool
	PDUSessionResourceFailedToSetupListHOAck    ngapType.PDUSessionResourceFailedToSetupListHOAck `ngap:"id:56,criticality:ignore,presence:optional"`
	TargetToSourceTransparentContainer          ngapType.TargetToSourceTransparentContainer       `ngap:"id:106,criticality:reject,presence:mandatory"`
	HasCriticalityDiagnostics                   bool
	CriticalityDiagnostics                      ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasNPNAccessInformation                     bool
	NPNAccessInformation                        ngapType.NPNAccessInformation `ngap:"id:259,criticality:reject,presence:optional"`
	HasRedCapIndication                         bool
	RedCapIndication                            ngapType.RedCapIndication `ngap:"id:333,criticality:ignore,presence:optional"`
	HasERedCapIndication                        bool
	ERedCapIndication                           ngapType.ERedCapIndication `ngap:"id:427,criticality:ignore,presence:optional"`
	AdditionalIEs                               []ngapType.UnknownIE
	UnknownExtensions                           []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverRequestAcknowledge) Encode() ([]byte, error) {
	return ngapcodec.Encode("HandoverRequestAcknowledge", m)
}

// DecodeHandoverRequestAcknowledge decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverRequestAcknowledge(data []byte) (*HandoverRequestAcknowledge, error) {
	var m HandoverRequestAcknowledge
	err := ngapcodec.Decode("HandoverRequestAcknowledge", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
