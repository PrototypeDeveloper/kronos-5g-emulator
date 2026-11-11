// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionReleaseRequest represents ASN.1 BroadcastSessionReleaseRequest.
type BroadcastSessionReleaseRequest struct {
	MBSSessionID      ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	Cause             ngapType.Cause        `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionReleaseRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionReleaseRequest", m)
}

// DecodeBroadcastSessionReleaseRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionReleaseRequest(data []byte) (*BroadcastSessionReleaseRequest, error) {
	var m BroadcastSessionReleaseRequest
	err := ngapcodec.Decode("BroadcastSessionReleaseRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
