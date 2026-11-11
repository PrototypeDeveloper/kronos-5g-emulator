// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionUpdateResponse represents ASN.1 MulticastSessionUpdateResponse.
type MulticastSessionUpdateResponse struct {
	MBSSessionID              ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID       bool
	MBSAreaSessionID          ngapType.MBSAreaSessionID `ngap:"id:295,criticality:reject,presence:optional"`
	HasCriticalityDiagnostics bool
	CriticalityDiagnostics    ngapType.CriticalityDiagnostics `ngap:"id:19,criticality:ignore,presence:optional"`
	AdditionalIEs             []ngapType.UnknownIE
	UnknownExtensions         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionUpdateResponse) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionUpdateResponse", m)
}

// DecodeMulticastSessionUpdateResponse decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionUpdateResponse(data []byte) (*MulticastSessionUpdateResponse, error) {
	var m MulticastSessionUpdateResponse
	err := ngapcodec.Decode("MulticastSessionUpdateResponse", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
