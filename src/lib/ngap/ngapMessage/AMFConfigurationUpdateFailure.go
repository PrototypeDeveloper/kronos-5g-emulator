// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AMFConfigurationUpdateFailure represents ASN.1 AMFConfigurationUpdateFailure.
type AMFConfigurationUpdateFailure struct {
	Cause                     ngapType.Cause `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasTimeToWait             bool
	TimeToWait                ngapType.TimeToWait `ngap:"id:107,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AMFConfigurationUpdateFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("AMFConfigurationUpdateFailure", m)
}

// DecodeAMFConfigurationUpdateFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAMFConfigurationUpdateFailure(data []byte) (*AMFConfigurationUpdateFailure, error) {
	var m AMFConfigurationUpdateFailure
	err := ngapcodec.Decode("AMFConfigurationUpdateFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
