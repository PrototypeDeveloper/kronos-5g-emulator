package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// AuthenticationFailure contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type AuthenticationFailure struct {
	Cause                             uint8
	HasAuthenticationFailureParameter bool
	AuthenticationFailureParameter    nasType.AuthenticationFailureParameter
}

var authenticationFailureOptions = nascodec.Schema{
	nasType.IEAuthenticationFailureParameter: {Format: nascodec.TLV, Min: 14, Max: 14},
}

func (m *AuthenticationFailure) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil AuthenticationFailure")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgAuthenticationFailure}
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasAuthenticationFailureParameter {
		options = append(options, nasType.IE{ID: nasType.IEAuthenticationFailureParameter, Value: m.AuthenticationFailureParameter.Value})
	}
	w.Options(options, authenticationFailureOptions)
	return w.Result()
}
func DecodeAuthenticationFailure(data []byte) (*AuthenticationFailure, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgAuthenticationFailure {
		return nil, fmt.Errorf("unexpected header for AuthenticationFailure")
	}
	m := &AuthenticationFailure{}
	r := nascodec.Reader{Data: data[3:]}
	m.Cause = r.U8(255)
	for _, ie := range r.Options(authenticationFailureOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEAuthenticationFailureParameter:
			m.HasAuthenticationFailureParameter = true
			m.AuthenticationFailureParameter = nasType.AuthenticationFailureParameter{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode AuthenticationFailure: %w", r.Err)
	}
	return m, nil
}
