package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionAuthenticationComplete carries the UE's EAP message for secondary authentication.
type PDUSessionAuthenticationComplete struct {
	PDUSessionID, ProcedureTransactionID    uint8
	EAPMessage                              []byte
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	UnknownIEs                              []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionAuthenticationCompleteOptions = nascodec.Schema{
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionAuthenticationComplete) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionAuthenticationComplete")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionAuthenticationComplete, 255)
	w.LV(m.EAPMessage, 2, 4, 1500)
	var options []nasType.IE
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionAuthenticationCompleteOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionAuthenticationCompleteOptions)
	return w.Result()
}
func DecodePDUSessionAuthenticationComplete(data []byte) (*PDUSessionAuthenticationComplete, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionAuthenticationComplete {
		return nil, fmt.Errorf("unexpected header for PDUSessionAuthenticationComplete")
	}
	m := &PDUSessionAuthenticationComplete{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	m.EAPMessage = r.LV(2, 4, 1500)
	known, unknown := r.OptionsWithUnknown(pDUSessionAuthenticationCompleteOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionAuthenticationComplete: %w", r.Err)
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
