// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverCancelAcknowledge represents ASN.1 HandoverCancelAcknowledge.
type HandoverCancelAcknowledge struct {
	AMFUENGAPID               ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID               ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverCancelAcknowledge) Encode() ([]byte, error) {
	return ngapcodec.Encode("HandoverCancelAcknowledge", m)
}

// DecodeHandoverCancelAcknowledge decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverCancelAcknowledge(data []byte) (*HandoverCancelAcknowledge, error) {
	var m HandoverCancelAcknowledge
	err := ngapcodec.Decode("HandoverCancelAcknowledge", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
