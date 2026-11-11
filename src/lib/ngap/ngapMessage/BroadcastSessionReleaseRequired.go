// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionReleaseRequired represents ASN.1 BroadcastSessionReleaseRequired.
type BroadcastSessionReleaseRequired struct {
	MBSSessionID      ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	Cause             ngapType.Cause        `ngap:"id:15,criticality:ignore,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionReleaseRequired) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionReleaseRequired", m)
}

// DecodeBroadcastSessionReleaseRequired decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionReleaseRequired(data []byte) (*BroadcastSessionReleaseRequired, error) {
	var m BroadcastSessionReleaseRequired
	err := ngapcodec.Decode("BroadcastSessionReleaseRequired", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
