package nasMessage

import (
	"fmt"
	"libnas/nasType"
)

// DeregistrationAcceptUE acknowledges a UE-originating de-registration.
type DeregistrationAcceptUE struct{}

func (m *DeregistrationAcceptUE) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil DeregistrationAcceptUE")
	}
	return []byte{nasType.EPD5GMM, 0, nasType.MsgDeregistrationAcceptUE}, nil
}
func DecodeDeregistrationAcceptUE(data []byte) (*DeregistrationAcceptUE, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgDeregistrationAcceptUE || len(data) != 3 {
		return nil, fmt.Errorf("invalid DeregistrationAcceptUE")
	}
	return &DeregistrationAcceptUE{}, nil
}
