// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RANConfigurationUpdateAcknowledge represents ASN.1 RANConfigurationUpdateAcknowledge.
type RANConfigurationUpdateAcknowledge struct {
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RANConfigurationUpdateAcknowledge) Encode() ([]byte, error) {
	return ngapcodec.Encode("RANConfigurationUpdateAcknowledge", m)
}

// DecodeRANConfigurationUpdateAcknowledge decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRANConfigurationUpdateAcknowledge(data []byte) (*RANConfigurationUpdateAcknowledge, error) {
	var m RANConfigurationUpdateAcknowledge
	err := ngapcodec.Decode("RANConfigurationUpdateAcknowledge", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
