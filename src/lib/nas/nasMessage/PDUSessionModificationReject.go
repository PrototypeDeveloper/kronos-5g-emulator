package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionModificationReject rejects a PDU session modification request.
type PDUSessionModificationReject struct {
	PDUSessionID, ProcedureTransactionID    uint8
	Cause                                   uint8 // 5GSM cause
	HasBackOffTimerValue                    bool
	BackOffTimerValue                       nasType.BackOffTimerValue
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	HasReAttemptIndicator                   bool
	ReAttemptIndicator                      nasType.ReAttemptIndicator
	HasCongestionReAttemptIndicator5GSM     bool
	CongestionReAttemptIndicator5GSM        nasType.CongestionReAttemptIndicator5GSM
	UnknownIEs                              []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionModificationRejectOptions = nascodec.Schema{
	nasType.IEBackOffTimerValue:                    {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEReAttemptIndicator:                   {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IECongestionReAttemptIndicator5GSM:     {Format: nascodec.TLV, Min: 1, Max: 1},
}

func (m *PDUSessionModificationReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionModificationReject")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionModificationReject, 255)
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasBackOffTimerValue {
		options = append(options, nasType.IE{ID: nasType.IEBackOffTimerValue, Value: []byte{m.BackOffTimerValue.Value}})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	if m.HasReAttemptIndicator {
		options = append(options, nasType.IE{ID: nasType.IEReAttemptIndicator, Value: []byte{m.ReAttemptIndicator.Value}})
	}
	if m.HasCongestionReAttemptIndicator5GSM {
		options = append(options, nasType.IE{ID: nasType.IECongestionReAttemptIndicator5GSM, Value: []byte{m.CongestionReAttemptIndicator5GSM.Value}})
	}
	w.Options(options, pDUSessionModificationRejectOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionModificationRejectOptions)
	return w.Result()
}
func DecodePDUSessionModificationReject(data []byte) (*PDUSessionModificationReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionModificationReject {
		return nil, fmt.Errorf("unexpected header for PDUSessionModificationReject")
	}
	m := &PDUSessionModificationReject{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	m.Cause = r.U8(255)
	known, unknown := r.OptionsWithUnknown(pDUSessionModificationRejectOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionModificationReject: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEBackOffTimerValue:
			m.HasBackOffTimerValue = true
			m.BackOffTimerValue = nasType.BackOffTimerValue{Value: ie.Value[0]}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		case nasType.IEReAttemptIndicator:
			m.HasReAttemptIndicator = true
			m.ReAttemptIndicator = nasType.ReAttemptIndicator{Value: ie.Value[0]}
		case nasType.IECongestionReAttemptIndicator5GSM:
			m.HasCongestionReAttemptIndicator5GSM = true
			m.CongestionReAttemptIndicator5GSM = nasType.CongestionReAttemptIndicator5GSM{Value: ie.Value[0]}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
