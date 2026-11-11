package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionEstablishmentReject contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type PDUSessionEstablishmentReject struct {
	PDUSessionID                            uint8
	ProcedureTransactionID                  uint8
	Cause                                   uint8
	HasBackOffTimerValue                    bool
	BackOffTimerValue                       nasType.BackOffTimerValue
	HasCongestionReAttemptIndicator5GSM     bool
	CongestionReAttemptIndicator5GSM        nasType.CongestionReAttemptIndicator5GSM
	HasEAPMessage                           bool
	EAPMessage                              nasType.EAPMessage
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
}

var pDUSessionEstablishmentRejectOptions = nascodec.Schema{
	nasType.IEBackOffTimerValue:                    {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IECongestionReAttemptIndicator5GSM:     {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEEAPMessage:                           {Format: nascodec.TLVE, Min: 4, Max: 1500},
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionEstablishmentReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionEstablishmentReject")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionEstablishmentReject, 255)
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasBackOffTimerValue {
		options = append(options, nasType.IE{ID: nasType.IEBackOffTimerValue, Value: []byte{m.BackOffTimerValue.Value}})
	}
	if m.HasCongestionReAttemptIndicator5GSM {
		options = append(options, nasType.IE{ID: nasType.IECongestionReAttemptIndicator5GSM, Value: []byte{m.CongestionReAttemptIndicator5GSM.Value}})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionEstablishmentRejectOptions)
	return w.Result()
}
func DecodePDUSessionEstablishmentReject(data []byte) (*PDUSessionEstablishmentReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionEstablishmentReject {
		return nil, fmt.Errorf("unexpected header for PDUSessionEstablishmentReject")
	}
	m := &PDUSessionEstablishmentReject{}
	r := nascodec.Reader{Data: data[4:]}
	m.PDUSessionID = h.PDUSessionID
	m.ProcedureTransactionID = h.ProcedureTransactionID
	m.Cause = r.U8(255)
	for _, ie := range r.Options(pDUSessionEstablishmentRejectOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEBackOffTimerValue:
			m.HasBackOffTimerValue = true
			m.BackOffTimerValue = nasType.BackOffTimerValue{Value: ie.Value[0]}
		case nasType.IECongestionReAttemptIndicator5GSM:
			m.HasCongestionReAttemptIndicator5GSM = true
			m.CongestionReAttemptIndicator5GSM = nasType.CongestionReAttemptIndicator5GSM{Value: ie.Value[0]}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionEstablishmentReject: %w", r.Err)
	}
	return m, nil
}
