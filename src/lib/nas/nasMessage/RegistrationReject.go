package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// RegistrationReject contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type RegistrationReject struct {
	Cause         uint8
	HasT3346Value bool
	T3346Value    nasType.T3346Value
	HasT3502Value bool
	T3502Value    nasType.T3502Value
	HasEAPMessage bool
	EAPMessage    nasType.EAPMessage
}

var registrationRejectOptions = nascodec.Schema{
	nasType.IET3346Value: {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IET3502Value: {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEEAPMessage: {Format: nascodec.TLVE, Min: 4, Max: 1500},
}

func (m *RegistrationReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil RegistrationReject")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgRegistrationReject}
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasT3346Value {
		options = append(options, nasType.IE{ID: nasType.IET3346Value, Value: []byte{m.T3346Value.Value}})
	}
	if m.HasT3502Value {
		options = append(options, nasType.IE{ID: nasType.IET3502Value, Value: []byte{m.T3502Value.Value}})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	w.Options(options, registrationRejectOptions)
	return w.Result()
}
func DecodeRegistrationReject(data []byte) (*RegistrationReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgRegistrationReject {
		return nil, fmt.Errorf("unexpected header for RegistrationReject")
	}
	m := &RegistrationReject{}
	r := nascodec.Reader{Data: data[3:]}
	m.Cause = r.U8(255)
	for _, ie := range r.Options(registrationRejectOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IET3346Value:
			m.HasT3346Value = true
			m.T3346Value = nasType.T3346Value{Value: ie.Value[0]}
		case nasType.IET3502Value:
			m.HasT3502Value = true
			m.T3502Value = nasType.T3502Value{Value: ie.Value[0]}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode RegistrationReject: %w", r.Err)
	}
	return m, nil
}
