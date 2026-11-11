package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// Status5GMM contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type Status5GMM struct {
	Cause uint8
}

var status5GMMOptions = nascodec.Schema{}

func (m *Status5GMM) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil Status5GMM")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgStatus5GMM}
	w.U8(m.Cause, 255)
	w.Options(nil, status5GMMOptions)
	return w.Result()
}
func DecodeStatus5GMM(data []byte) (*Status5GMM, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgStatus5GMM {
		return nil, fmt.Errorf("unexpected header for Status5GMM")
	}
	m := &Status5GMM{}
	r := nascodec.Reader{Data: data[3:]}
	m.Cause = r.U8(255)
	r.Options(status5GMMOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode Status5GMM: %w", r.Err)
	}
	return m, nil
}
