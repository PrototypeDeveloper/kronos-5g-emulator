// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// TraceFailureIndication represents ASN.1 TraceFailureIndication.
type TraceFailureIndication struct {
	AMFUENGAPID       ngapType.AMFUENGAPID  `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID  `ngap:"id:85,criticality:reject,presence:mandatory"`
	NGRANTraceID      ngapType.NGRANTraceID `ngap:"id:44,criticality:ignore,presence:mandatory"`
	Cause             ngapType.Cause        `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *TraceFailureIndication) Encode() ([]byte, error) {
	return ngapcodec.Encode("TraceFailureIndication", m)
}

// DecodeTraceFailureIndication decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeTraceFailureIndication(data []byte) (*TraceFailureIndication, error) {
	var m TraceFailureIndication
	err := ngapcodec.Decode("TraceFailureIndication", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
