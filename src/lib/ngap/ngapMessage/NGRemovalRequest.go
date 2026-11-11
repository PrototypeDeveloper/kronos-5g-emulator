// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGRemovalRequest represents ASN.1 NGRemovalRequest.
type NGRemovalRequest struct {
	GlobalRANNodeID   ngapType.GlobalRANNodeID `ngap:"id:27,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGRemovalRequest) Encode() ([]byte, error) { return ngapcodec.Encode("NGRemovalRequest", m) }

// DecodeNGRemovalRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGRemovalRequest(data []byte) (*NGRemovalRequest, error) {
	var m NGRemovalRequest
	err := ngapcodec.Decode("NGRemovalRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
