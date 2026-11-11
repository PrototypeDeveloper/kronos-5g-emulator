package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// IdentityRequest contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type IdentityRequest struct {
	IdentityType uint8
}

var identityRequestOptions = nascodec.Schema{}

func (m *IdentityRequest) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil IdentityRequest")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgIdentityRequest}
	w.U8(m.IdentityType, 7)
	w.Options(nil, identityRequestOptions)
	return w.Result()
}
func DecodeIdentityRequest(data []byte) (*IdentityRequest, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgIdentityRequest {
		return nil, fmt.Errorf("unexpected header for IdentityRequest")
	}
	m := &IdentityRequest{}
	r := nascodec.Reader{Data: data[3:]}
	m.IdentityType = r.U8(7)
	r.Options(identityRequestOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode IdentityRequest: %w", r.Err)
	}
	return m, nil
}
