// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGResetAcknowledge represents ASN.1 NGResetAcknowledge.
type NGResetAcknowledge struct {
	HasUEAssociatedLogicalNGConnectionList bool
	UEAssociatedLogicalNGConnectionList    ngapType.UEAssociatedLogicalNGConnectionList `ngap:"id:111,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics              bool
	CriticalityDiagnostics                 ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs                          []ngapType.UnknownIE
	UnknownExtensions                      []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGResetAcknowledge) Encode() ([]byte, error) {
	return ngapcodec.Encode("NGResetAcknowledge", m)
}

// DecodeNGResetAcknowledge decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGResetAcknowledge(data []byte) (*NGResetAcknowledge, error) {
	var m NGResetAcknowledge
	err := ngapcodec.Decode("NGResetAcknowledge", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
