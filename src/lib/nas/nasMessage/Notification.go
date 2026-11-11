package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// Notification pages the UE over non-3GPP access, or the reverse.
type Notification struct {
	AccessType uint8        // low nibble; the other bits are spare
	UnknownIEs []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var notificationOptions = nascodec.Schema{}

func (m *Notification) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil Notification")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgNotification}
	w.U8(m.AccessType, 3)
	w.UnknownOptions(m.UnknownIEs, notificationOptions)
	return w.Result()
}
func DecodeNotification(data []byte) (*Notification, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgNotification {
		return nil, fmt.Errorf("unexpected header for Notification")
	}
	m := &Notification{}
	r := nascodec.Reader{Data: data[3:]}
	m.AccessType = r.U8(3)
	_, unknown := r.OptionsWithUnknown(notificationOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode Notification: %w", r.Err)
	}
	m.UnknownIEs = unknown
	return m, nil
}
