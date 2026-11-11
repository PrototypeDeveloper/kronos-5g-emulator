// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// UETNLABindingReleaseRequest represents ASN.1 UETNLABindingReleaseRequest.
type UETNLABindingReleaseRequest struct {
	AMFUENGAPID       ngapType.AMFUENGAPID `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID       ngapType.RANUENGAPID `ngap:"id:85,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *UETNLABindingReleaseRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("UETNLABindingReleaseRequest", m)
}

// DecodeUETNLABindingReleaseRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeUETNLABindingReleaseRequest(data []byte) (*UETNLABindingReleaseRequest, error) {
	var m UETNLABindingReleaseRequest
	err := ngapcodec.Decode("UETNLABindingReleaseRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
