// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// NGReset represents ASN.1 NGReset.
type NGReset struct {
	Cause             ngapType.Cause     `ngap:"id:15,criticality:ignore,presence:mandatory"`
	ResetType         ngapType.ResetType `ngap:"id:88,criticality:reject,presence:mandatory"`
	AdditionalIEs     []ngapType.UnknownIE
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *NGReset) Encode() ([]byte, error) { return ngapcodec.Encode("NGReset", m) }

// DecodeNGReset decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodeNGReset(data []byte) (*NGReset, error) {
	var m NGReset
	err := ngapcodec.Decode("NGReset", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
