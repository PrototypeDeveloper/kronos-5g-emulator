// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MTCommunicationHandlingRequest represents ASN.1 MTCommunicationHandlingRequest.
type MTCommunicationHandlingRequest struct {
	AMFUENGAPID              ngapType.AMFUENGAPID              `ngap:"id:10,criticality:reject,presence:mandatory"`
	RANUENGAPID              ngapType.RANUENGAPID              `ngap:"id:85,criticality:reject,presence:mandatory"`
	FiveGCAction             ngapType.FiveGCAction             `ngap:"id:406,criticality:reject,presence:mandatory"`
	PDUSessionListMTCommHReq ngapType.PDUSessionListMTCommHReq `ngap:"id:431,criticality:reject,presence:mandatory"`
	AdditionalIEs            []ngapType.UnknownIE
	UnknownExtensions        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MTCommunicationHandlingRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("MTCommunicationHandlingRequest", m)
}

// DecodeMTCommunicationHandlingRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMTCommunicationHandlingRequest(data []byte) (*MTCommunicationHandlingRequest, error) {
	var m MTCommunicationHandlingRequest
	err := ngapcodec.Decode("MTCommunicationHandlingRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
