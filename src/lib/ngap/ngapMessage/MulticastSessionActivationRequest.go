// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionActivationRequest represents ASN.1 MulticastSessionActivationRequest.
type MulticastSessionActivationRequest struct {
	MBSSessionID                              ngapType.MBSSessionID                                         `ngap:"id:299,criticality:reject,presence:mandatory"`
	MulticastSessionActivationRequestTransfer ngapType.MulticastSessionActivationRequestTransferOCTETSTRING `ngap:"id:304,criticality:reject,presence:mandatory"`
	AdditionalIEs                             []ngapType.UnknownIE
	UnknownExtensions                         []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionActivationRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionActivationRequest", m)
}

// DecodeMulticastSessionActivationRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionActivationRequest(data []byte) (*MulticastSessionActivationRequest, error) {
	var m MulticastSessionActivationRequest
	err := ngapcodec.Decode("MulticastSessionActivationRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
