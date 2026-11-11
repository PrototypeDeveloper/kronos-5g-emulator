// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PathSwitchRequestFailure represents ASN.1 PathSwitchRequestFailure.
type PathSwitchRequestFailure struct {
	AMFUENGAPID                          ngapType.AMFUENGAPID                          `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                          ngapType.RANUENGAPID                          `ngap:"id:85,criticality:ignore,presence:mandatory"`
	PDUSessionResourceReleasedListPSFail ngapType.PDUSessionResourceReleasedListPSFail `ngap:"id:69,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics            bool
	CriticalityDiagnostics               ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                        []ngapType.UnknownIE
	UnknownExtensions                    []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PathSwitchRequestFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("PathSwitchRequestFailure", m)
}

// DecodePathSwitchRequestFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePathSwitchRequestFailure(data []byte) (*PathSwitchRequestFailure, error) {
	var m PathSwitchRequestFailure
	err := ngapcodec.Decode("PathSwitchRequestFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
