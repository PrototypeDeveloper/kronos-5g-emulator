// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGSetupFailure represents ASN.1 NGSetupFailure.
type NGSetupFailure struct {
	Cause                     ngapType.Cause `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasTimeToWait             bool
	TimeToWait                ngapType.TimeToWait `ngap:"id:107,criticality:ignore,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGSetupFailure) Encode() ([]byte, error) { return ngapcodec.Encode("NGSetupFailure", m) }

// DecodeNGSetupFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGSetupFailure(data []byte) (*NGSetupFailure, error) {
	var m NGSetupFailure
	err := ngapcodec.Decode("NGSetupFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
