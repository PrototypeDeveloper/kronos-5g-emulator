// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RANConfigurationUpdateFailure represents ASN.1 RANConfigurationUpdateFailure.
type RANConfigurationUpdateFailure struct {
	Cause                     ngapType.Cause `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasTimeToWait             bool
	TimeToWait                ngapType.TimeToWait `ngap:"id:107,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RANConfigurationUpdateFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("RANConfigurationUpdateFailure", m)
}

// DecodeRANConfigurationUpdateFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRANConfigurationUpdateFailure(data []byte) (*RANConfigurationUpdateFailure, error) {
	var m RANConfigurationUpdateFailure
	err := ngapcodec.Decode("RANConfigurationUpdateFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
