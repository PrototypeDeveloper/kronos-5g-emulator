package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionModificationCommandReject rejects a PDU session modification command.
type PDUSessionModificationCommandReject struct {
	PDUSessionID, ProcedureTransactionID    uint8
	Cause                                   uint8 // 5GSM cause
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	UnknownIEs                              []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionModificationCommandRejectOptions = nascodec.Schema{
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionModificationCommandReject) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionModificationCommandReject")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionModificationCommandReject, 255)
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionModificationCommandRejectOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionModificationCommandRejectOptions)
	return w.Result()
}
func DecodePDUSessionModificationCommandReject(data []byte) (*PDUSessionModificationCommandReject, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionModificationCommandReject {
		return nil, fmt.Errorf("unexpected header for PDUSessionModificationCommandReject")
	}
	m := &PDUSessionModificationCommandReject{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	m.Cause = r.U8(255)
	known, unknown := r.OptionsWithUnknown(pDUSessionModificationCommandRejectOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionModificationCommandReject: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
