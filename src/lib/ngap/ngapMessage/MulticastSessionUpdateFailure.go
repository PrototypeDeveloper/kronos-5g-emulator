// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionUpdateFailure represents ASN.1 MulticastSessionUpdateFailure.
type MulticastSessionUpdateFailure struct {
	MBSSessionID              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID       bool
	MBSAreaSessionID          ngapType.MBSAreaSessionID `ngap:"id:295,criticality:reject,presence:optional"`
	Cause                     ngapType.Cause            `ngap:"id:15,criticality:ignore,presence:mandatory"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionUpdateFailure) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionUpdateFailure", m)
}

// DecodeMulticastSessionUpdateFailure decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionUpdateFailure(data []byte) (*MulticastSessionUpdateFailure, error) {
	var m MulticastSessionUpdateFailure
	err := ngapcodec.Decode("MulticastSessionUpdateFailure", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
