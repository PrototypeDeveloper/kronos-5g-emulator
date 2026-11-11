package gtpType

import (
	"encoding/binary"
	"fmt"
)

const (
	Version1 uint8 = 1

	MsgEchoRequest                           uint8 = 1
	MsgEchoResponse                          uint8 = 2
	MsgErrorIndication                       uint8 = 26
	MsgSupportedExtensionHeadersNotification uint8 = 31
	MsgEndMarker                             uint8 = 254
	MsgGPDU                                  uint8 = 255
)

// Header represents the GTPv1-U header.
// Length is the number of octets following the mandatory 8-octet header.
// If any of E/S/PN is set, SequenceNumber, NPDUNumber and
// NextExtensionHeaderType are all present on the wire as a 4-octet block.
type Header struct {
	Version      uint8
	ProtocolType bool
	Spare        bool
	E            bool
	S            bool
	PN           bool

	MessageType uint8
	Length      uint16
	TEID        uint32

	SequenceNumber          uint16
	NPDUNumber              uint8
	NextExtensionHeaderType uint8
}

func NewHeader(messageType uint8, teid uint32) Header {
	return Header{
		Version:      Version1,
		ProtocolType: true,
		MessageType:  messageType,
		TEID:         teid,
	}
}

func (h Header) HasOptionalFields() bool {
	return h.E || h.S || h.PN
}

func (h Header) EncodeMandatory() ([]byte, error) {
	if h.Version != Version1 {
		return nil, fmt.Errorf("unsupported GTP version: %d", h.Version)
	}
	if !h.ProtocolType {
		return nil, fmt.Errorf("protocol type must be GTP for GTPv1-U")
	}

	b := make([]byte, 8)
	b[0] = (h.Version & 0x07) << 5
	if h.ProtocolType {
		b[0] |= 0x10
	}
	if h.Spare {
		b[0] |= 0x08
	}
	if h.E {
		b[0] |= 0x04
	}
	if h.S {
		b[0] |= 0x02
	}
	if h.PN {
		b[0] |= 0x01
	}

	b[1] = h.MessageType
	binary.BigEndian.PutUint16(b[2:4], h.Length)
	binary.BigEndian.PutUint32(b[4:8], h.TEID)
	return b, nil
}

func (h Header) EncodeOptional() []byte {
	if !h.HasOptionalFields() {
		return nil
	}
	b := make([]byte, 4)
	binary.BigEndian.PutUint16(b[0:2], h.SequenceNumber)
	b[2] = h.NPDUNumber
	b[3] = h.NextExtensionHeaderType
	return b
}

func DecodeHeader(data []byte) (Header, error) {
	var h Header
	if len(data) < 8 {
		return h, fmt.Errorf("GTP packet too short: %d", len(data))
	}

	flags := data[0]
	h.Version = (flags >> 5) & 0x07
	h.ProtocolType = flags&0x10 != 0
	h.Spare = flags&0x08 != 0
	h.E = flags&0x04 != 0
	h.S = flags&0x02 != 0
	h.PN = flags&0x01 != 0
	h.MessageType = data[1]
	h.Length = binary.BigEndian.Uint16(data[2:4])
	h.TEID = binary.BigEndian.Uint32(data[4:8])

	if h.Version != Version1 {
		return h, fmt.Errorf("unsupported GTP version: %d", h.Version)
	}
	if !h.ProtocolType {
		return h, fmt.Errorf("not a GTP packet: PT=0")
	}

	totalLength := 8 + int(h.Length)
	if len(data) != totalLength {
		return h, fmt.Errorf("GTP length mismatch: header=%d actual=%d", totalLength, len(data))
	}

	if h.HasOptionalFields() {
		if len(data) < 12 {
			return h, fmt.Errorf("GTP optional fields are truncated")
		}
		h.SequenceNumber = binary.BigEndian.Uint16(data[8:10])
		h.NPDUNumber = data[10]
		h.NextExtensionHeaderType = data[11]
	}

	return h, nil
}
