// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGRemovalFailure represents ASN.1 NGRemovalFailure.
type NGRemovalFailure struct {
	Cause                     ngapType.Cause `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGRemovalFailure) Encode() ([]byte, error) { return ngapcodec.Encode("NGRemovalFailure", m) }

// DecodeNGRemovalFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGRemovalFailure(data []byte) (*NGRemovalFailure, error) {
	var m NGRemovalFailure
	err := ngapcodec.Decode("NGRemovalFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
