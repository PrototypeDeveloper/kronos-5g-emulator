// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// MulticastSessionUpdateRequest represents ASN.1 MulticastSessionUpdateRequest.
type MulticastSessionUpdateRequest struct {
	MBSSessionID                          ngapType.MBSSessionID `ngap:"id:299,criticality:reject,presence:mandatory"`
	HasMBSAreaSessionID                   bool
	MBSAreaSessionID                      ngapType.MBSAreaSessionID                                 `ngap:"id:295,criticality:reject,presence:optional"`
	MulticastSessionUpdateRequestTransfer ngapType.MulticastSessionUpdateRequestTransferOCTETSTRING `ngap:"id:306,criticality:reject,presence:mandatory"`
	AdditionalIEs                         []ngapType.UnknownIE
	UnknownExtensions                     []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *MulticastSessionUpdateRequest) Encode() ([]byte, error) {
	return ngapcodec.Encode("MulticastSessionUpdateRequest", m)
}

// DecodeMulticastSessionUpdateRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeMulticastSessionUpdateRequest(data []byte) (*MulticastSessionUpdateRequest, error) {
	var m MulticastSessionUpdateRequest
	err := ngapcodec.Decode("MulticastSessionUpdateRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
