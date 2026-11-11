// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverCancel represents ASN.1 HandoverCancel.
type HandoverCancel struct {
	AMFUENGAPID       ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	Cause             ngapType.Cause       `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverCancel) Encode() ([]byte, error) { return ngapcodec.Encode("HandoverCancel", m) }

// DecodeHandoverCancel decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverCancel(data []byte) (*HandoverCancel, error) {
	var m HandoverCancel
	err := ngapcodec.Decode("HandoverCancel", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
