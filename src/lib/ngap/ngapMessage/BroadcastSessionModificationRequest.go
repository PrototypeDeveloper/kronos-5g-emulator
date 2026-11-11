// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// BroadcastSessionModificationRequest represents ASN.1 BroadcastSessionModificationRequest.
type BroadcastSessionModificationRequest struct {
	MBSSessionID                             ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSServiceArea                        bool
	MBSServiceArea                           ngapType.MBSServiceArea `ngap:"id:298,criticality:reject,presence:optional"`
	HasMBSSessionModificationRequestTransfer bool
	MBSSessionModificationRequestTransfer    ngapType.MBSSessionSetupOrModRequestTransferOCTETSTRING `ngap:"id:349,criticality:reject,presence:optional"`
	AdditionalIEs                            []ngapType.UnknownIE
	UnknownExtensions                        []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *BroadcastSessionModificationRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("BroadcastSessionModificationRequest", m)
}

// DecodeBroadcastSessionModificationRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeBroadcastSessionModificationRequest(data []byte) (*BroadcastSessionModificationRequest, error) {
	var m BroadcastSessionModificationRequest
	err := ngapcodec.Decode("BroadcastSessionModificationRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
