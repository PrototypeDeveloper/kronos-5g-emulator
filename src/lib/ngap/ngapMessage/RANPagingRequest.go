// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// RANPagingRequest represents ASN.1 RANPagingRequest.
type RANPagingRequest struct {
	AMFUENGAPID                    ngapType.AMFUENGAPID `ngap:"id:10,criticality:ignore,presence:mandatory"`
	RANUENGAPID                    ngapType.RANUENGAPID `ngap:"id:85,criticality:ignore,presence:mandatory"`
	HasPagingPolicyDifferentiation bool
	PagingPolicyDifferentiation    ngapType.PagingPolicyDifferentiation `ngap:"id:407,criticality:ignore,presence:optional"`
	HasDLSignalling                bool
	DLSignalling                   ngapType.DLSignalling `ngap:"id:408,criticality:ignore,presence:optional"`
	AdditionalIEs                  []ngapType.UnknownIE
	UnknownExtensions              []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *RANPagingRequest) Encode() ([]byte, error) { return ngapcodec.Encode("RANPagingRequest", m) }

// DecodeRANPagingRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeRANPagingRequest(data []byte) (*RANPagingRequest, error) {
	var m RANPagingRequest
	err := ngapcodec.Decode("RANPagingRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
