package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionReleaseCommand carries a plain 5GSM message (TS 24.501 8.3.12–15).
type PDUSessionReleaseCommand struct {
	PDUSessionID, ProcedureTransactionID    uint8
	Cause                                   uint8
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
}

var pDUSessionReleaseCommandOptions = nascodec.Schema{
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *PDUSessionReleaseCommand) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionReleaseCommand")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionReleaseCommand, 255)
	w.U8(m.Cause, 255)
	var options []nasType.IE
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	w.Options(options, pDUSessionReleaseCommandOptions)
	return w.Result()
}
func DecodePDUSessionReleaseCommand(data []byte) (*PDUSessionReleaseCommand, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionReleaseCommand {
		return nil, fmt.Errorf("unexpected PDUSessionReleaseCommand header")
	}
	m := &PDUSessionReleaseCommand{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	m.Cause = r.U8(255)
	for _, ie := range r.Options(pDUSessionReleaseCommandOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionReleaseCommand: %w", r.Err)
	}
	return m, nil
}
