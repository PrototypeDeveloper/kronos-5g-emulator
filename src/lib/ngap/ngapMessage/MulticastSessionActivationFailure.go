// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionActivationFailure represents ASN.1 MulticastSessionActivationFailure.
type MulticastSessionActivationFailure struct {
	MBSSessionID              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	Cause                     ngapType.Cause        `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionActivationFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionActivationFailure", m)
}

// DecodeMulticastSessionActivationFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionActivationFailure(data []byte) (*MulticastSessionActivationFailure, error) {
	var m MulticastSessionActivationFailure
	err := ngapcodec.Decode("MulticastSessionActivationFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
