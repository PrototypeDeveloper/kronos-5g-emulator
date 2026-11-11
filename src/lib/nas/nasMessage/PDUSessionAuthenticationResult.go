package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionAuthenticationResult carries the EAP result of secondary authentication.
type PDUSessionAuthenticationResult struct {
	PDUSessionID, ProcedureTransactionID    uint8
	HasEAPMessage                           bool
	EAPMessage                              nasType.EAPMessage
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	UnknownIEs                              []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionAuthenticationResultOptions = nascodec.Schema{
	nasType.IEEAPMessage:                           {Format: nascodec.TLVE, Min: 4, Max: 1500},
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionAuthenticationResult) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionAuthenticationResult")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionAuthenticationResult, 255)
	var options []nasType.IE
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionAuthenticationResultOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionAuthenticationResultOptions)
	return w.Result()
}
func DecodePDUSessionAuthenticationResult(data []byte) (*PDUSessionAuthenticationResult, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionAuthenticationResult {
		return nil, fmt.Errorf("unexpected header for PDUSessionAuthenticationResult")
	}
	m := &PDUSessionAuthenticationResult{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	known, unknown := r.OptionsWithUnknown(pDUSessionAuthenticationResultOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionAuthenticationResult: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
