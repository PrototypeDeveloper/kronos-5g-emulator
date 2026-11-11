package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// AuthenticationRequest contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type AuthenticationRequest struct {
	NgKSI                          uint8
	ABBA                           []byte
	HasAuthenticationParameterRAND bool
	AuthenticationParameterRAND    nasType.AuthenticationParameterRAND
	HasAuthenticationParameterAUTN bool
	AuthenticationParameterAUTN    nasType.AuthenticationParameterAUTN
	HasEAPMessage                  bool
	EAPMessage                     nasType.EAPMessage
}

var authenticationRequestOptions = nascodec.Schema{
	nasType.IEAuthenticationParameterRAND: {Format: nascodec.TV, Min: 16, Max: 16},
	nasType.IEAuthenticationParameterAUTN: {Format: nascodec.TLV, Min: 16, Max: 16},
	nasType.IEEAPMessage:                  {Format: nascodec.TLVE, Min: 4, Max: 1500},
}

func (m *AuthenticationRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil AuthenticationRequest")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgAuthenticationRequest}
	w.U8(m.NgKSI, 15)
	w.LV(m.ABBA, 1, 2, 255)
	var options []nasType.IE
	if m.HasAuthenticationParameterRAND {
		options = append(options, nasType.IE{ID: nasType.IEAuthenticationParameterRAND, Value: m.AuthenticationParameterRAND.Value})
	}
	if m.HasAuthenticationParameterAUTN {
		options = append(options, nasType.IE{ID: nasType.IEAuthenticationParameterAUTN, Value: m.AuthenticationParameterAUTN.Value})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	w.Options(options, authenticationRequestOptions)
	return w.Result()
}
func DecodeAuthenticationRequest(data []byte) (*AuthenticationRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgAuthenticationRequest {
		return nil, fmt.Errorf("unexpected header for AuthenticationRequest")
	}
	m := &AuthenticationRequest{}
	r := nascodec.Reader{Data: data[3:]}
	m.NgKSI = r.U8(15)
	m.ABBA = r.LV(1, 2, 255)
	for _, ie := range r.Options(authenticationRequestOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEAuthenticationParameterRAND:
			m.HasAuthenticationParameterRAND = true
			m.AuthenticationParameterRAND = nasType.AuthenticationParameterRAND{Value: ie.Value}
		case nasType.IEAuthenticationParameterAUTN:
			m.HasAuthenticationParameterAUTN = true
			m.AuthenticationParameterAUTN = nasType.AuthenticationParameterAUTN{Value: ie.Value}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode AuthenticationRequest: %w", r.Err)
	}
	return m, nil
}
