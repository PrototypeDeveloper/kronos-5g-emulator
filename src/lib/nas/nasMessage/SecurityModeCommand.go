package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// SecurityModeCommand contains a plain NAS message. Optional fields are encoded only when their Has flag is true.
type SecurityModeCommand struct {
	SelectedNASAlgorithms               uint8
	NgKSI                               uint8
	ReplayedUESecurityCapabilities      []byte
	HasIMEISVRequest                    bool
	IMEISVRequest                       nasType.IMEISVRequest
	HasSelectedEPSNASAlgorithms         bool
	SelectedEPSNASAlgorithms            nasType.SelectedEPSNASAlgorithms
	HasAdditional5GSecurityInformation  bool
	Additional5GSecurityInformation     nasType.Additional5GSecurityInformation
	HasEAPMessage                       bool
	EAPMessage                          nasType.EAPMessage
	HasABBA                             bool
	ABBA                                nasType.ABBA
	HasReplayedS1UESecurityCapabilities bool
	ReplayedS1UESecurityCapabilities    nasType.ReplayedS1UESecurityCapabilities
}

var securityModeCommandOptions = nascodec.Schema{
	nasType.IEIMEISVRequest:                    {Format: nascodec.HalfTV, Min: 1, Max: 1},
	nasType.IESelectedEPSNASAlgorithms:         {Format: nascodec.TV, Min: 1, Max: 1},
	nasType.IEAdditional5GSecurityInformation:  {Format: nascodec.TLV, Min: 1, Max: 1},
	nasType.IEEAPMessage:                       {Format: nascodec.TLVE, Min: 4, Max: 1500},
	nasType.IEABBA:                             {Format: nascodec.TLV, Min: 2, Max: 255},
	nasType.IEReplayedS1UESecurityCapabilities: {Format: nascodec.TLV, Min: 2, Max: 5},
}

func (m *SecurityModeCommand) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil SecurityModeCommand")
	}
	w := nascodec.Writer{}
	w.Data = []byte{nasType.EPD5GMM, 0, nasType.MsgSecurityModeCommand}
	w.U8(m.SelectedNASAlgorithms, 255)
	w.U8(m.NgKSI, 15)
	w.LV(m.ReplayedUESecurityCapabilities, 1, 2, 8)
	var options []nasType.IE
	if m.HasIMEISVRequest {
		options = append(options, nasType.IE{ID: nasType.IEIMEISVRequest, Value: []byte{m.IMEISVRequest.Value}})
	}
	if m.HasSelectedEPSNASAlgorithms {
		options = append(options, nasType.IE{ID: nasType.IESelectedEPSNASAlgorithms, Value: []byte{m.SelectedEPSNASAlgorithms.Value}})
	}
	if m.HasAdditional5GSecurityInformation {
		options = append(options, nasType.IE{ID: nasType.IEAdditional5GSecurityInformation, Value: []byte{m.Additional5GSecurityInformation.Value}})
	}
	if m.HasEAPMessage {
		options = append(options, nasType.IE{ID: nasType.IEEAPMessage, Value: m.EAPMessage.Value})
	}
	if m.HasABBA {
		options = append(options, nasType.IE{ID: nasType.IEABBA, Value: m.ABBA.Value})
	}
	if m.HasReplayedS1UESecurityCapabilities {
		options = append(options, nasType.IE{ID: nasType.IEReplayedS1UESecurityCapabilities, Value: m.ReplayedS1UESecurityCapabilities.Value})
	}
	w.Options(options, securityModeCommandOptions)
	return w.Result()
}
func DecodeSecurityModeCommand(data []byte) (*SecurityModeCommand, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GMM || h.MessageType != nasType.MsgSecurityModeCommand {
		return nil, fmt.Errorf("unexpected header for SecurityModeCommand")
	}
	m := &SecurityModeCommand{}
	r := nascodec.Reader{Data: data[3:]}
	m.SelectedNASAlgorithms = r.U8(255)
	m.NgKSI = r.U8(15)
	m.ReplayedUESecurityCapabilities = r.LV(1, 2, 8)
	for _, ie := range r.Options(securityModeCommandOptions) {
		if r.Err != nil {
			break
		}
		switch ie.ID {
		case nasType.IEIMEISVRequest:
			m.HasIMEISVRequest = true
			m.IMEISVRequest = nasType.IMEISVRequest{Value: ie.Value[0]}
		case nasType.IESelectedEPSNASAlgorithms:
			m.HasSelectedEPSNASAlgorithms = true
			m.SelectedEPSNASAlgorithms = nasType.SelectedEPSNASAlgorithms{Value: ie.Value[0]}
		case nasType.IEAdditional5GSecurityInformation:
			m.HasAdditional5GSecurityInformation = true
			m.Additional5GSecurityInformation = nasType.Additional5GSecurityInformation{Value: ie.Value[0]}
		case nasType.IEEAPMessage:
			m.HasEAPMessage = true
			m.EAPMessage = nasType.EAPMessage{Value: ie.Value}
		case nasType.IEABBA:
			m.HasABBA = true
			m.ABBA = nasType.ABBA{Value: ie.Value}
		case nasType.IEReplayedS1UESecurityCapabilities:
			m.HasReplayedS1UESecurityCapabilities = true
			m.ReplayedS1UESecurityCapabilities = nasType.ReplayedS1UESecurityCapabilities{Value: ie.Value}
		}
	}
	if r.Err != nil {
		return nil, fmt.Errorf("decode SecurityModeCommand: %w", r.Err)
	}
	return m, nil
}
