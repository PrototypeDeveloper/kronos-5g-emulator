package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// DeregistrationAcceptNW acknowledges a UE-terminated de-registration.
type DeregistrationAcceptNW struct {
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var deregistrationAcceptNWOptions = nascodec.Schema{}

func (m *DeregistrationAcceptNW) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil DeregistrationAcceptNW")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgDeregistrationAcceptNW}
	w.UnknownOptions(m.UnknownIEs, deregistrationAcceptNWOptions)
	return w.Result()
}
func DecodeDeregistrationAcceptNW(data []byte) (*DeregistrationAcceptNW, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgDeregistrationAcceptNW {
		return nil, fmt.Errorf("unexpected header for DeregistrationAcceptNW")
	}
	m := &DeregistrationAcceptNW{}
	r := nascodec.Reader{Data: data[3:]}
	_, unknown := r.OptionsWithUnknown(deregistrationAcceptNWOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode DeregistrationAcceptNW: %w", r.Err)
	}
	m.UnknownIEs = unknown
	return m, nil
}
