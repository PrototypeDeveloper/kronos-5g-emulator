package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// IdentityResponse contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type IdentityResponse struct {
	MobileIdentity []byte
}

var identityResponseOptions = nascodec.Schema{}

func (m *IdentityResponse) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil IdentityResponse")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgIdentityResponse}
	w.LV(m.MobileIdentity, 2, 4, 65535)
	w.Options(nil, identityResponseOptions)
	return w.Result()
}
func DecodeIdentityResponse(data []byte) (*IdentityResponse, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgIdentityResponse {
		return nil, fmt.Errorf("unexpected header for IdentityResponse")
	}
	m := &IdentityResponse{}
	r := nascodec.Reader{Data: data[3:]}
	m.MobileIdentity = r.LV(2, 4, 65535)
	r.Options(identityResponseOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode IdentityResponse: %w", r.Err)
	}
	return m, nil
}
