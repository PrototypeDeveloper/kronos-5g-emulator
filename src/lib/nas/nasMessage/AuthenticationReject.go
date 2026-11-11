package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// AuthenticationReject contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type AuthenticationReject struct {
	HasEAPMessage bool
	EAPMessage    nasType.EAPMessage
}

var authenticationRejectOptions = nascodec.Schema{
	nasType.IEEAPMessage: {Format: nascodec.TLVE, Min: 4, Max: 1500},
}

func (m *AuthenticationReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil AuthenticationReject")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgAuthenticationReject}
	var options []nasType.IE
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	w.Options(options, authenticationRejectOptions)
	return w.Result()
}
func DecodeAuthenticationReject(data []byte) (*AuthenticationReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgAuthenticationReject {
		return nil, fmt.Errorf("unexpected header for AuthenticationReject")
	}
	m := &AuthenticationReject{}
	r := nascodec.Reader{Data: data[3:]}
	for _, ie := range r.Options(authenticationRejectOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode AuthenticationReject: %w", r.Err)
	}
	return m, nil
}
