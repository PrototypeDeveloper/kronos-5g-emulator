// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionDeactivationRequest represents ASN.1 MulticastSessionDeactivationRequest.
type MulticastSessionDeactivationRequest struct {
	MBSSessionID                                ngapType.MBSSessionID                                           `ngap:"id:299,criticality:reject,presence:mandatory"`
	MulticastSessionDeactivationRequestTransfer ngapType.MulticastSessionDeactivationRequestTransferOCTETSTRING `ngap:"id:305,criticality:reject,presence:mandatory"`
	AdditionalIEs                               []ngapType.UnknownIE
	UnknownExtensions                           []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionDeactivationRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionDeactivationRequest", m)
}

// DecodeMulticastSessionDeactivationRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionDeactivationRequest(data []byte) (*MulticastSessionDeactivationRequest, error) {
	var m MulticastSessionDeactivationRequest
	err := ngapcodec.Decode("MulticastSessionDeactivationRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
