package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// NotificationResponse answers a notification.
type NotificationResponse struct {
	HasPDUSessionStatus bool
	PDUSessionStatus    nasType.PDUSessionStatus
	UnknownIEs          []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var notificationResponseOptions = nascodec.Schema{
	nasType.IEPDUSessionStatus: {Format: nascodec.TLV, Min: 2, Max: 32},
}

func (m *NotificationResponse) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil NotificationResponse")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgNotificationResponse}
	var options []nasType.IE
	if m.HasPDUSessionStatus {
		options = append(options, nasType.IE{ID: nasType.IEPDUSessionStatus, Value: m.PDUSessionStatus.Value})
	}
	w.Options(options, notificationResponseOptions)
	w.UnknownOptions(m.UnknownIEs, notificationResponseOptions)
	return w.Result()
}
func DecodeNotificationResponse(data []byte) (*NotificationResponse, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgNotificationResponse {
		return nil, fmt.Errorf("unexpected header for NotificationResponse")
	}
	m := &NotificationResponse{}
	r := nascodec.Reader{Data: data[3:]}
	known, unknown := r.OptionsWithUnknown(notificationResponseOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode NotificationResponse: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEPDUSessionStatus:
			m.HasPDUSessionStatus = true
			m.PDUSessionStatus = nasType.PDUSessionStatus{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
