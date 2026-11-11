package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// SecurityModeComplete contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type SecurityModeComplete struct {
	HasIMEISV              bool
	IMEISV                 nasType.IMEISV
	HasNASMessageContainer bool
	NASMessageContainer    nasType.NASMessageContainer
	HasNonIMEISVPEI        bool
	NonIMEISVPEI           nasType.NonIMEISVPEI
}

var securityModeCompleteOptions = nascodec.Schema{
	nasType.IEIMEISV:              {Format: nascodec.TLVE, Min: 9, Max: 9},
	nasType.IENASMessageContainer: {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IENonIMEISVPEI:        {Format: nascodec.TLVE, Min: 1, Max: 65535},
}

func (m *SecurityModeComplete) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil SecurityModeComplete")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgSecurityModeComplete}
	var options []nasType.IE
	if m.HasIMEISV {
		options = append(options, nasType.IE{ID: nasType.IEIMEISV, Value: m.IMEISV.Value})
	}
	if m.HasNASMessageContainer {
		options = append(options, nasType.IE{ID: nasType.IENASMessageContainer, Value: m.NASMessageContainer.Value})
	}
	if m.HasNonIMEISVPEI {
		options = append(options, nasType.IE{ID: nasType.IENonIMEISVPEI, Value: m.NonIMEISVPEI.Value})
	}
	w.Options(options, securityModeCompleteOptions)
	return w.Result()
}
func DecodeSecurityModeComplete(data []byte) (*SecurityModeComplete, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgSecurityModeComplete {
		return nil, fmt.Errorf("unexpected header for SecurityModeComplete")
	}
	m := &SecurityModeComplete{}
	r := nascodec.Reader{Data: data[3:]}
	for _, ie := range r.Options(securityModeCompleteOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEIMEISV:
			m.HasIMEISV = true
			m.IMEISV = nasType.IMEISV{Value: ie.Value}
		case nasType.IENASMessageContainer:
			m.HasNASMessageContainer = true
			m.NASMessageContainer = nasType.NASMessageContainer{Value: ie.Value}
		case nasType.IENonIMEISVPEI:
			m.HasNonIMEISVPEI = true
			m.NonIMEISVPEI = nasType.NonIMEISVPEI{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode SecurityModeComplete: %w", r.Err)
	}
	return m, nil
}
