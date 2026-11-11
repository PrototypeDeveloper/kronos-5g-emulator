// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// AMFConfigurationUpdateAcknowledge represents ASN.1 AMFConfigurationUpdateAcknowledge.
type AMFConfigurationUpdateAcknowledge struct {
	HasAMFTNLAssociationSetupList         bool
	AMFTNLAssociationSetupList            ngapType.AMFTNLAssociationSetupList `ngap:"id:5,criticality:ignore,presence:optional"`
	HasAMFTNLAssociationFailedToSetupList bool
	AMFTNLAssociationFailedToSetupList    ngapType.TNLAssociationList `ngap:"id:4,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics             bool
	CriticalityDiagnostics                ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                         []ngapType.UnknownIE
	UnknownExtensions                     []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *AMFConfigurationUpdateAcknowledge) Encode() ([]byte, error) {
	return ngapcodec.Encode("AMFConfigurationUpdateAcknowledge", m)
}

// DecodeAMFConfigurationUpdateAcknowledge decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeAMFConfigurationUpdateAcknowledge(data []byte) (*AMFConfigurationUpdateAcknowledge, error) {
	var m AMFConfigurationUpdateAcknowledge
	err := ngapcodec.Decode("AMFConfigurationUpdateAcknowledge", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
