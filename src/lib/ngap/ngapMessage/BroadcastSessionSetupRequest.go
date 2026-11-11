// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionSetupRequest represents ASN.1 BroadcastSessionSetupRequest.
type BroadcastSessionSetupRequest struct {
	MBSSessionID                   ngapType.MBSSessionID                                   `ngap:"id:299,criticality:reject,presence:mandatory"`
	SNSSAI                         ngapType.SNSSAI                                         `ngap:"id:148,criticality:reject,presence:mandatory"`
	MBSServiceArea                 ngapType.MBSServiceArea                                 `ngap:"id:298,criticality:reject,presence:mandatory"`
	MBSSessionSetupRequestTransfer ngapType.MBSSessionSetupOrModRequestTransferOCTETSTRING `ngap:"id:315,criticality:reject,presence:mandatory"`
	HasAssociatedSessionID         bool
	AssociatedSessionID            ngapType.AssociatedSessionID `ngap:"id:415,criticality:ignore,presence:optional"`
	AdditionalIEs                  []ngapType.UnknownIE
	UnknownExtensions              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionSetupRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionSetupRequest", m)
}

// DecodeBroadcastSessionSetupRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionSetupRequest(data []byte) (*BroadcastSessionSetupRequest, error) {
	var m BroadcastSessionSetupRequest
	err := ngapcodec.Decode("BroadcastSessionSetupRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
