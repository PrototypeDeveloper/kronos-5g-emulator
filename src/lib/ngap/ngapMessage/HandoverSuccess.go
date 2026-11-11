// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// HandoverSuccess represents ASN.1 HandoverSuccess.
type HandoverSuccess struct {
	AMFUENGAPID       ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *HandoverSuccess) Encode() ([]byte, error) { return ngapcodec.Encode("HandoverSuccess", m) }

// DecodeHandoverSuccess decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeHandoverSuccess(data []byte) (*HandoverSuccess, error) {
	var m HandoverSuccess
	err := ngapcodec.Decode("HandoverSuccess", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
