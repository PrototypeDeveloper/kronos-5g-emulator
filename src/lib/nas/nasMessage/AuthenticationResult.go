package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// AuthenticationResult carries the EAP result of a primary authentication.
type AuthenticationResult struct {
	NgKSI      uint8 // low nibble; the high nibble is spare
	EAPMessage []byte
	HasABBA    bool
	ABBA       nasType.ABBA
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var authenticationResultOptions = nascodec.Schema{
	nasType.IEABBA: {Format: nascodec.TLV, Min: 2, Max: 255},
}

func (m *AuthenticationResult) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil AuthenticationResult")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgAuthenticationResult}
	w.U8(m.NgKSI, 15)
	w.LV(m.EAPMessage, 2, 4, 1500)
	var options []nasType.IE
	if m.HasABBA {
		options = append(options, nasType.IE{ID: nasType.IEABBA, Value: m.ABBA.Value})
	}
	w.Options(options, authenticationResultOptions)
	w.UnknownOptions(m.UnknownIEs, authenticationResultOptions)
	return w.Result()
}
func DecodeAuthenticationResult(data []byte) (*AuthenticationResult, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgAuthenticationResult {
		return nil, fmt.Errorf("unexpected header for AuthenticationResult")
	}
	m := &AuthenticationResult{}
	r := nascodec.Reader{Data: data[3:]}
	m.NgKSI = r.U8(15)
	m.EAPMessage = r.LV(2, 4, 1500)
	known, unknown := r.OptionsWithUnknown(authenticationResultOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode AuthenticationResult: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEABBA:
			m.HasABBA = true
			m.ABBA = nasType.ABBA{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
