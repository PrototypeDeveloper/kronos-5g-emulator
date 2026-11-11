// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PWSCancelRequest represents ASN.1 PWSCancelRequest.
type PWSCancelRequest struct {
	MessageIdentifier           ngapType.MessageIdentifier `ngap:"id:35,criticality:reject,presence:mandatory"`
	SerialNumber                ngapType.SerialNumber      `ngap:"id:95,criticality:reject,presence:mandatory"`
	HasWarningAreaList          bool
	WarningAreaList             ngapType.WarningAreaList `ngap:"id:122,criticality:ignore,presence:optional"`
	HasCancelAllWarningMessages bool
	CancelAllWarningMessages    ngapType.CancelAllWarningMessages `ngap:"id:14,criticality:reject,presence:optional"`
	AdditionalIEs               []ngapType.UnknownIE
	UnknownExtensions           []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PWSCancelRequest) Encode() ([]byte, error) { return ngapcodec.Encode("PWSCancelRequest", m) }

// DecodePWSCancelRequest decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePWSCancelRequest(data []byte) (*PWSCancelRequest, error) {
	var m PWSCancelRequest
	err := ngapcodec.Decode("PWSCancelRequest", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
