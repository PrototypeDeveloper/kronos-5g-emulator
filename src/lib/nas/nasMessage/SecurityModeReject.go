package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// SecurityModeReject contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type SecurityModeReject struct {
	Cause uint8
}

var securityModeRejectOptions = nascodec.Schema{}

func (m *SecurityModeReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil SecurityModeReject")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgSecurityModeReject}
	w.U8(m.Cause, 255)
	w.Options(nil, securityModeRejectOptions)
	return w.Result()
}
func DecodeSecurityModeReject(data []byte) (*SecurityModeReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgSecurityModeReject {
		return nil, fmt.Errorf("unexpected header for SecurityModeReject")
	}
	m := &SecurityModeReject{}
	r := nascodec.Reader{Data: data[3:]}
	m.Cause = r.U8(255)
	r.Options(securityModeRejectOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode SecurityModeReject: %w", r.Err)
	}
	return m, nil
}
