package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionAuthenticationCommand carries an EAP message for secondary authentication.
type PDUSessionAuthenticationCommand struct {
	PDUSessionID, ProcedureTransactionID    uint8
	EAPMessage                              []byte
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	UnknownIEs                              []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionAuthenticationCommandOptions = nascodec.Schema{
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionAuthenticationCommand) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionAuthenticationCommand")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionAuthenticationCommand, 255)
	w.LV(m.EAPMessage, 2, 4, 1500)
	var options []nasType.IE
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionAuthenticationCommandOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionAuthenticationCommandOptions)
	return w.Result()
}
func DecodePDUSessionAuthenticationCommand(data []byte) (*PDUSessionAuthenticationCommand, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionAuthenticationCommand {
		return nil, fmt.Errorf("unexpected header for PDUSessionAuthenticationCommand")
	}
	m := &PDUSessionAuthenticationCommand{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	m.EAPMessage = r.LV(2, 4, 1500)
	known, unknown := r.OptionsWithUnknown(pDUSessionAuthenticationCommandOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionAuthenticationCommand: %w", r.Err)
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
