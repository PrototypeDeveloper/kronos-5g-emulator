// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionTransportRequest represents ASN.1 BroadcastSessionTransportRequest.
type BroadcastSessionTransportRequest struct {
	MBSSessionID                      ngapType.MBSSessionID                                 `ngap:"id:299,criticality:reject,presence:mandatory"`
	BroadcastTransportRequestTransfer ngapType.BroadcastTransportRequestTransferOCTETSTRING `ngap:"id:418,criticality:reject,presence:mandatory"`
	AdditionalIEs                     []ngapType.UnknownIE
	UnknownExtensions                 []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionTransportRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionTransportRequest", m)
}

// DecodeBroadcastSessionTransportRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionTransportRequest(data []byte) (*BroadcastSessionTransportRequest, error) {
	var m BroadcastSessionTransportRequest
	err := ngapcodec.Decode("BroadcastSessionTransportRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
