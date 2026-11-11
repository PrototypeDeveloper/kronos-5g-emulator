// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// TraceStart represents ASN.1 TraceStart.
type TraceStart struct {
	AMFUENGAPID       ngapType.AMFUENGAPID     `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID     `ngap:"id:85,criticality:reject,presence:mandatory"`
	TraceActivation   ngapType.TraceActivation `ngap:"id:108,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *TraceStart) Encode() ([]byte, error) { return ngapcodec.Encode("TraceStart", m) }

// DecodeTraceStart decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeTraceStart(data []byte) (*TraceStart, error) {
	var m TraceStart
	err := ngapcodec.Decode("TraceStart", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
