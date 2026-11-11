package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// RegistrationComplete contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type RegistrationComplete struct {
	HasSORTransparentContainer bool
	SORTransparentContainer    nasType.SORTransparentContainer
}

var registrationCompleteOptions = nascodec.Schema{
	nasType.IESORTransparentContainer: {Format: nascodec.TLVE, Min: 0, Max: 65535},
}

func (m *RegistrationComplete) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil RegistrationComplete")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgRegistrationComplete}
	var options []nasType.IE
	if m.HasSORTransparentContainer {
		options = append(options, nasType.IE{ID: nasType.IESORTransparentContainer, Value: m.SORTransparentContainer.Value})
	}
	w.Options(options, registrationCompleteOptions)
	return w.Result()
}
func DecodeRegistrationComplete(data []byte) (*RegistrationComplete, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgRegistrationComplete {
		return nil, fmt.Errorf("unexpected header for RegistrationComplete")
	}
	m := &RegistrationComplete{}
	r := nascodec.Reader{Data: data[3:]}
	for _, ie := range r.Options(registrationCompleteOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IESORTransparentContainer:
			m.HasSORTransparentContainer = true
			m.SORTransparentContainer = nasType.SORTransparentContainer{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode RegistrationComplete: %w", r.Err)
	}
	return m, nil
}
