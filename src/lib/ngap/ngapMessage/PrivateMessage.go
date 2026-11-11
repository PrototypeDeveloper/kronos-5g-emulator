// ASN.1 structure derived from TS 38.413; codec methods are maintained in libngap.

package ngapMessage

import (
	"libngap/ngapType"
	"libngap/ngapcodec"
)

// PrivateMessage represents ASN.1 PrivateMessage.
type PrivateMessage struct {
	PrivateIEs        []ngapType.PrivateMessageIEsEntry
	UnknownExtensions []ngapType.UnknownExtension
}

// Encode encodes the complete NGAP PDU.
func (m *PrivateMessage) Encode() ([]byte, error) { return ngapcodec.Encode("PrivateMessage", m) }

// DecodePrivateMessage decodes the complete NGAP PDU. NAS contents remain opaque.
func DecodePrivateMessage(data []byte) (*PrivateMessage, error) {
	var m PrivateMessage
	err := ngapcodec.Decode("PrivateMessage", data, &m)
	if ngapcodec.IsFatal(err) {
		return nil, err
	}
	return &m, err
}
