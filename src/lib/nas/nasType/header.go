package nasType

import (
	"fmt"
	"io"
)

const (
	EPD5GMM        uint8 = 0x7e
	EPD5GSM        uint8 = 0x2e
	MaxMessageSize       = 1024 * 1024
)

type Header struct {
	ExtendedProtocolDiscriminator uint8
	SecurityHeaderType            SecurityHeaderType
	MessageType                   uint8
	PDUSessionID                  uint8
	ProcedureTransactionID        uint8
}

// DecodeHeader reads a plain message header. Security-protected PDUs must first
// be verified and unwrapped by libnas.Decode.
func DecodeHeader(data []byte) (Header, error) {
	var h Header
	if len(data) < 1 {
		return h, io.ErrUnexpectedEOF
	}
	if len(data) > MaxMessageSize {
		return h, fmt.Errorf("NAS message exceeds %d bytes", MaxMessageSize)
	}
	h.ExtendedProtocolDiscriminator = data[0]
	switch data[0] {
	case EPD5GMM:
		if len(data) < 3 {
			return h, io.ErrUnexpectedEOF
		}
		if data[1] != 0 {
			return h, fmt.Errorf("expected a plain 5GMM header")
		}
		h.MessageType = data[2]
	case EPD5GSM:
		if len(data) < 4 {
			return h, io.ErrUnexpectedEOF
		}
		h.PDUSessionID = data[1]
		h.ProcedureTransactionID = data[2]
		h.MessageType = data[3]
		if h.PDUSessionID > 15 {
			return h, fmt.Errorf("invalid PDU session ID %d", h.PDUSessionID)
		}
	default:
		return h, fmt.Errorf("unsupported NAS extended protocol discriminator 0x%02x", data[0])
	}
	return h, nil
}
