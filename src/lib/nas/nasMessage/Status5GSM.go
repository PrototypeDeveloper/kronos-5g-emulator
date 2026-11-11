package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// Status5GSM contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type Status5GSM struct {
	PDUSessionID           uint8
	ProcedureTransactionID uint8
	Cause                  uint8
}

var status5GSMOptions = nascodec.Schema{}

func (m *Status5GSM) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil Status5GSM")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgStatus5GSM, 255)
	w.U8(m.Cause, 255)
	w.Options(nil, status5GSMOptions)
	return w.Result()
}
func DecodeStatus5GSM(data []byte) (*Status5GSM, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgStatus5GSM {
		return nil, fmt.Errorf("unexpected header for Status5GSM")
	}
	m := &Status5GSM{}
	r := nascodec.Reader{Data: data[4:]}
	m.PDUSessionID = h.PDUSessionID
	m.ProcedureTransactionID = h.ProcedureTransactionID
	m.Cause = r.U8(255)
	r.Options(status5GSMOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode Status5GSM: %w", r.Err)
	}
	return m, nil
}
