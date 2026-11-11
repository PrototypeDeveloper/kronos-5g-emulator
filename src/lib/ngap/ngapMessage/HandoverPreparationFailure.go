// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverPreparationFailure represents ASN.1 HandoverPreparationFailure.
type HandoverPreparationFailure struct {
	AMFUENGAPID                                  ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                                  ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	Cause                                        ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics                    bool
	CriticalityDiagnostics                       ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	HasTargettoSourceFailureTransparentContainer bool
	TargettoSourceFailureTransparentContainer    ngapType.TargettoSourceFailureTransparentContainer `ngap:"id:262,criticality:ignore,presence:optional"`
	AdditionalIEs                                []ngapType.UnknownIE
	UnknownExtensions                            []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverPreparationFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("HandoverPreparationFailure", m)
}

// DecodeHandoverPreparationFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverPreparationFailure(data []byte) (*HandoverPreparationFailure, error) {
	var m HandoverPreparationFailure
	err := ngapcodec.Decode("HandoverPreparationFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
