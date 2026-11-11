// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// DeactivateTrace represents ASN.1 DeactivateTrace.
type DeactivateTrace struct {
	AMFUENGAPID       ngapType.AMFUENGAPID  `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID  `ngap:"id:85,criticality:reject,presence:mandatory"`
	NGRANTraceID      ngapType.NGRANTraceID `ngap:"id:44,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *DeactivateTrace) Encode() ([]byte, error) { return ngapcodec.Encode("DeactivateTrace", m) }

// DecodeDeactivateTrace decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeDeactivateTrace(data []byte) (*DeactivateTrace, error) {
	var m DeactivateTrace
	err := ngapcodec.Decode("DeactivateTrace", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
