package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// AuthenticationResponse contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type AuthenticationResponse struct {
	HasAuthenticationResponseParameter bool
	AuthenticationResponseParameter    nasType.AuthenticationResponseParameter
	HasEAPMessage                      bool
	EAPMessage                         nasType.EAPMessage
}

var authenticationResponseOptions = nascodec.Schema{
	nasType.IEAuthenticationResponseParameter: {Format: nascodec.TLV, Min: 4, Max: 16},
	nasType.IEEAPMessage:                      {Format: nascodec.TLVE, Min: 4, Max: 1500},
}

func (m *AuthenticationResponse) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil AuthenticationResponse")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgAuthenticationResponse}
	var options []nasType.IE
	if m.HasAuthenticationResponseParameter {
		options = append(options, nasType.IE{ID: nasType.IEAuthenticationResponseParameter, Value: m.AuthenticationResponseParameter.Value})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	w.Options(options, authenticationResponseOptions)
	return w.Result()
}
func DecodeAuthenticationResponse(data []byte) (*AuthenticationResponse, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgAuthenticationResponse {
		return nil, fmt.Errorf("unexpected header for AuthenticationResponse")
	}
	m := &AuthenticationResponse{}
	r := nascodec.Reader{Data: data[3:]}
	for _, ie := range r.Options(authenticationResponseOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEAuthenticationResponseParameter:
			m.HasAuthenticationResponseParameter = true
			m.AuthenticationResponseParameter = nasType.AuthenticationResponseParameter{Value: ie.Value}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode AuthenticationResponse: %w", r.Err)
	}
	return m, nil
}
