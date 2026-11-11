package nasMessage

import (
	"fmt"
	"libnas/nasType"
	"libnas/nascodec"
)

// PDUSessionModificationComplete acknowledges a PDU session modification command.
type PDUSessionModificationComplete struct {
	PDUSessionID, ProcedureTransactionID    uint8
	HasExtendedProtocolConfigurationOptions bool
	ExtendedProtocolConfigurationOptions    nasType.ExtendedProtocolConfigurationOptions
	HasPortManagementInformationContainer   bool
	PortManagementInformationContainer      nasType.PortManagementInformationContainer
	UnknownIEs                              []nasType.IE // optional IEs not modelled above, kept for re-encoding
}

var pDUSessionModificationCompleteOptions = nascodec.Schema{
	nasType.IEExtendedProtocolConfigurationOptions: {Format: nascodec.TLVE, Min: 1, Max: 65535},
	nasType.IEPortManagementInformationContainer:   {Format: nascodec.TLVE, Min: 5, Max: 65535},
}

func (m *PDUSessionModificationComplete) Encode() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil PDUSessionModificationComplete")
	}
	w := nascodec.Writer{}
	w.U8(nasType.EPD5GSM, 255)
	w.U8(m.PDUSessionID, 15)
	w.U8(m.ProcedureTransactionID, 255)
	w.U8(nasType.MsgPDUSessionModificationComplete, 255)
	var options []nasType.IE
	if m.HasExtendedProtocolConfigurationOptions {
		options = append(options, nasType.IE{ID: nasType.IEExtendedProtocolConfigurationOptions, Value: m.ExtendedProtocolConfigurationOptions.Value})
	}
	if m.HasPortManagementInformationContainer {
		options = append(options, nasType.IE{ID: nasType.IEPortManagementInformationContainer, Value: m.PortManagementInformationContainer.Value})
	}
	w.Options(options, pDUSessionModificationCompleteOptions)
	w.UnknownOptions(m.UnknownIEs, pDUSessionModificationCompleteOptions)
	return w.Result()
}
func DecodePDUSessionModificationComplete(data []byte) (*PDUSessionModificationComplete, error) {
	h, err := nasType.DecodeHeader(data)
	if err != nil {
		return nil, err
	}
	if h.ExtendedProtocolDiscriminator != nasType.EPD5GSM || h.MessageType != nasType.MsgPDUSessionModificationComplete {
		return nil, fmt.Errorf("unexpected header for PDUSessionModificationComplete")
	}
	m := &PDUSessionModificationComplete{PDUSessionID: h.PDUSessionID, ProcedureTransactionID: h.ProcedureTransactionID}
	r := nascodec.Reader{Data: data[4:]}
	known, unknown := r.OptionsWithUnknown(pDUSessionModificationCompleteOptions)
	if r.Err != nil {
		return nil, fmt.Errorf("decode PDUSessionModificationComplete: %w", r.Err)
	}
	for _, ie := range known {
		switch ie.ID {
		case nasType.IEExtendedProtocolConfigurationOptions:
			m.HasExtendedProtocolConfigurationOptions = true
			m.ExtendedProtocolConfigurationOptions = nasType.ExtendedProtocolConfigurationOptions{Value: ie.Value}
		case nasType.IEPortManagementInformationContainer:
			m.HasPortManagementInformationContainer = true
			m.PortManagementInformationContainer = nasType.PortManagementInformationContainer{Value: ie.Value}
		}
	}
	m.UnknownIEs = unknown
	return m, nil
}
